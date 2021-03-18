package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/robfig/cron"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
	"wumingtianqi/config"
	"wumingtianqi/model/common"
	orderModel "wumingtianqi/model/order"
	"wumingtianqi/model/remind"
	"wumingtianqi/model/user"
	userLib "wumingtianqi/libs/user"
	"wumingtianqi/model/vip"
	weatherModel "wumingtianqi/model/weather"
	"wumingtianqi/utils"
	"wumingtianqi/utils/errnum"
)

// todo 用户建立订单
// req 有一些参数
// 中间处理过程比较麻烦？
// 然后存到订单表

// 1. 把符合提醒要求的订单放到队列（生产者）
// 1.1 拉取昨日，今日的天气表，存到map（以后做到缓存里）,涉及map？ 日后redis
// 1.2 订单表拉取9:00的订单
// 2. 拼接n个订单的信息，每个订单有多个提醒模式
// 3. 规范第2步：取队列的代码
// 20200530

// func1 假数据，天气信息, todo 放到weather 天气里
type WeatherItem map[string]interface{}
type Weather map[string]WeatherItem

func FakeWeather() (Weather, Weather) {  // 如果数据库中没有数据，可以用这里的假数据暂时测试
	yesterdayWeather := Weather{
		"WX4FBXXFKE4F": WeatherItem{
			"city_code":      "Beijing",
			"code_text": "晴",
			"code_id":   1,
			"high":      20,
		},
	}
	todayWeather := Weather{
		"WX4FBXXFKE4F": WeatherItem{
			"city_code":      "Beijing",
			"code_text": "阵雨",
			"code_id":   10,
			"high":      25,
		},
	}
	return yesterdayWeather, todayWeather
}

/**
 * @Author Evan
 * @Description 获取指定城市的昨天、今天的天气信息
	（这个注释是后来加的）
 * @Date 21:52 2021-02-23
 * @Param
 * @return
 **/
func Weather2Map(cityCode string) (Weather, Weather, error) {
	// 计算今日时间
	yesterdayDate8Int := utils.GetSpecificDate8Int(0)  // todo 这个之后需要协调today和yestoday
	yesterdayWeatherAll, has, err := weatherModel.QueryByCityDate(cityCode, yesterdayDate8Int)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, nil, err
	} else if !has {
		return nil, nil, errors.New("no Weather data of yesterday")
	}
	yesterdayWeather := Weather{
		cityCode: WeatherItem{
			"city_code": cityCode,
			"code_text": yesterdayWeatherAll.TextDay,  // todo 以后考虑把白天天气和晚上天气统一
			"code_id": yesterdayWeatherAll.CodeDay,
			"high": yesterdayWeatherAll.High,
		},
	}

	todayDate8Int := utils.GetSpecificDate8Int(1)  // todo 这个之后需要协调today和yestoday
	todayWeatherAll, has, err := weatherModel.QueryByCityDate(cityCode, todayDate8Int)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, nil, err
	} else if !has {
		return nil, nil, errors.New("no Weather data of today")
	}
	todayWeather := Weather{
		cityCode: WeatherItem{
			"city_code": cityCode,
			"code_text": todayWeatherAll.TextDay,
			"code_id": todayWeatherAll.CodeDay,
			"high": todayWeatherAll.High,
		},
	}
	return yesterdayWeather, todayWeather, nil
}

type SplicePatternModel struct {
	RemindSplicedText string `json:"remind_spliced_text"` // 拼接好的语句
	Priority          int    `json:"priority"`
}

func splicePattern1(cityCode string, remindPattern *remind.RemindPattern) (SplicePatternModel, error){
	// 1. 降水天气
	// 枚举"天气现象"表，整理降水的所有情况 [1,2,3]，考虑remind_pattern里新增一个extension字段（json格式），这个字段不同业务不一样，需要的东西也不一样。
	// 降雨对应的id todo 以后再弄个天气代码映射表？或者在某个地方弄个静态变量存
	RainPatternIds := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var todayWeather Weather
	var err error
	if config.GlobalConfig.Weather.FakeData {
		_, todayWeather = FakeWeather()
	} else {
		_, todayWeather, err = Weather2Map(cityCode)
		if err != nil {
			return SplicePatternModel{}, err
		}
	}

	codeText := todayWeather[cityCode]["code_text"].(string)
	codeToday := todayWeather[cityCode]["code_id"].(int)
	isTodayRain, _ := utils.IsContain(codeToday, RainPatternIds)

	var pattern = new(SplicePatternModel)
	if isTodayRain {
		pattern.RemindSplicedText = "有" + codeText + "记得带伞"  // 有阵雨 todo以后用format_text做通配，封装写法
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern, nil
}

func splicePattern2(cityCode string, remindPattern *remind.RemindPattern, value int) (SplicePatternModel, error) {
	// 2. 突然升温
	var yesterdayWeather, todayWeather Weather
	var err error
	if config.GlobalConfig.Weather.FakeData {
		yesterdayWeather, todayWeather = FakeWeather()
	} else {
		yesterdayWeather, todayWeather, err = Weather2Map(cityCode)
		if err != nil {
			return SplicePatternModel{}, err
		}
	}

	highYesterday := yesterdayWeather[cityCode]["high"].(int)
	highToday := todayWeather[cityCode]["high"].(int)
	highTodayStr := strconv.Itoa(todayWeather[cityCode]["high"].(int))
	valueStr := strconv.Itoa(highToday - highYesterday)

	var pattern = new(SplicePatternModel)
	log.Println("highToday - highYesterday", highToday - highYesterday)
	if highToday - highYesterday >= value {  // 最高气温较前一日增加5度，升至25度，注意防范
		remindObject := remindPattern.RemindObject
		pattern.RemindSplicedText = remindObject + "较前一日增加" + valueStr + "度，升至" + highTodayStr + "度，注意防范"
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern, nil
}

func splicePattern3(cityCode string, remindPattern *remind.RemindPattern, value int) (SplicePatternModel, error) {
	// 3. 突然降温
	var yesterdayWeather, todayWeather Weather
	var err error
	if config.GlobalConfig.Weather.FakeData {
		yesterdayWeather, todayWeather = FakeWeather()
	} else {
		yesterdayWeather, todayWeather, err = Weather2Map(cityCode)
		if err != nil {
			return SplicePatternModel{}, err
		}
	}
	highYesterday := yesterdayWeather[cityCode]["high"].(int)
	highToday := todayWeather[cityCode]["high"].(int)
	highTodayStr := strconv.Itoa(todayWeather[cityCode]["high"].(int))
	valueStr := strconv.Itoa(highYesterday - highToday)

	var pattern = new(SplicePatternModel)
	log.Println("highToday - highYesterday", highToday -highYesterday)
	if highYesterday- highToday >= value {  // 最高气温较前一日降低5度，降到18度，注意防范
		remindObject := remindPattern.RemindObject
		pattern.RemindSplicedText = remindObject + "较前一日降低" + valueStr + "度，降至" + highTodayStr + "度，注意防范"
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern, nil
}

func splicePattern8(cityCode string, remindPattern *remind.RemindPattern) (SplicePatternModel, error){
	// 8. 雨过天晴
	// 枚举"天气现象"表，整理突然降雨的触发条件 ![1,2,3] -> [1,2,3]，考虑remind_pattern里新增一个extension字段（json格式），这个字段不同业务不一样，需要的东西也不一样。
	// 降雨对应的id todo 以后再弄个天气代码映射表？或者在某个地方弄个静态变量存
	RainPatternIds := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var yesterdayWeather, todayWeather Weather
	var err error
	if config.GlobalConfig.Weather.FakeData {
		yesterdayWeather, todayWeather = FakeWeather()
	} else {
		yesterdayWeather, todayWeather, err = Weather2Map(cityCode)
		if err != nil {
			return SplicePatternModel{}, err
		}
	}

	codeText := todayWeather[cityCode]["code_text"].(string)
	codeYesterday := yesterdayWeather[cityCode]["code_id"].(int)
	codeToday := todayWeather[cityCode]["code_id"].(int)
	isYesRain, _ := utils.IsContain(codeYesterday, RainPatternIds)
	isTodayRain, _ := utils.IsContain(codeToday, RainPatternIds)

	var pattern = new(SplicePatternModel)
	if isYesRain && !isTodayRain {
		pattern.RemindSplicedText = "没有降水啦！明天天气：" + codeText // 雨过天晴
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern, nil
}

/**
 * @Author Evan
 * @Description 通知前预处理用户
	1. 获取用户信息
	2. vip等级过期处理
	3. LastRemindTime更新，如果LastRemindTime是昨天之前，需要把权益更新
	4. 剩余次数判断
	todo 未来做vip提醒权限的判断：判断是否有该提醒的权限？
 * @Date 20:58 2021-02-25
 * @Param
 * @return false 用户权限不足/处理失败，跳过通知; true 处理成功可以通知
 **/
func processBeforeNotify(userId int) (bool, error) {
	// 1 获取用户信息
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err := userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("get user_info_flexible model error: ", err.Error())
		return false, err
	} else if !has {
		log.Println("user_info_flexible model not exist")
		return false, errors.New("user_info_flexible model not exist")
	}
	// 2.vip等级过期判断
	userInfoFlexibleModel, err = userLib.CheckVipExpiration(userInfoFlexibleModel)
	if err != nil {
		log.Println("err: ", err.Error())
		return false, err
	}

	// 3.LastRemindTime更新，如果LastRemindTime是昨天之前，需要把权益更新
	todayDate8Int := utils.GetSpecificDate8Int(0)
	if userInfoFlexibleModel.LastRemindTime < todayDate8Int {
		userInfoFlexibleModel.LastRemindTime = todayDate8Int

		vipRightsMap := &vip.VipRightsMap{}
		vipRightsMap, has, err = vipRightsMap.QueryByVipLevel(userInfoFlexibleModel.VipLevel)
		if err != nil {
			err = errnum.New(errnum.DbError, err)
			log.Println("QueryByVipLevel err:", err.Error())
			return false, err
		} else if !has {
			log.Println("QueryByVipLevel not found")
			return false, errors.New("QueryByVipLevel not found")
		}
		userInfoFlexibleModel.TodayTelRemindRemaining = vipRightsMap.TelOrderMax
	}
	// 4. 剩余次数判断
	if userInfoFlexibleModel.TodayTelRemindRemaining <= 0 {
		return false, nil
	}
	// 提交更新
	if err := userInfoFlexibleModel.Update(); err != nil {
		err = errnum.New(errnum.DbError, err)
		return false, err
	}
	return true, nil
}


func SpliceOrders(oneOrderModel orderModel.Order) ([]SplicePatternModel, error){
	orderId := oneOrderModel.OrderId
	cityCode := oneOrderModel.RemindCity

	// 根据order_id找到order_detail
	orderDetail := orderModel.OrderDetail{}
	orderDetailList, err := orderDetail.QueryListByOrderId(orderId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}

	// 定义一个model，用来接单个提醒
	patterns := make([]SplicePatternModel, 0)
	for _, oneOrderDetailModel := range orderDetailList {
		value := oneOrderDetailModel.Value

		remindPattern := new(remind.RemindPattern)
		remindPattern, _, _ = remindPattern.QueryOneById(oneOrderDetailModel.RemindPatternId)

		switch oneOrderDetailModel.RemindPatternId {
		case 1: // 突然降雨
			pattern1, err := splicePattern1(cityCode, remindPattern)
			if err != nil {
				log.Println("splicePattern1 err", err)
				continue
			}
			if pattern1.Priority >= 1 {  // 以后可以用标准一点的用法
				patterns = append(patterns, pattern1)
			}
		case 2: // 突然升温
			pattern2, err := splicePattern2(cityCode, remindPattern, value)
			if err != nil {
				log.Println("splicePattern2 err", err)
			}
			if pattern2.Priority >= 1 {  // 以后可以用标准一点的用法
				patterns = append(patterns, pattern2)
			}
		case 3: // 突然降温
			pattern3, err := splicePattern3(cityCode, remindPattern, value)
			if err != nil {
				log.Println("splicePattern3 err", err)
			}
			if pattern3.Priority >= 1 {  // 以后可以用标准一点的用法
				patterns = append(patterns, pattern3)
			}
		case 4: // 空气质量变差
			//println(2)
		case 5: // 9点突然升温
			//println(2)
		case 6: // 高温预警
			//println(2)
		case 7: // 低温预警
			//println(2)
		case 8: // 雨过天晴
			pattern8, err := splicePattern8(cityCode, remindPattern)
			if err != nil {
				log.Println("splicePattern8 err", err)
			}
			if pattern8.Priority >= 1 {  // 以后可以用标准一点的用法
				patterns = append(patterns, pattern8)
			}
		}
	}
	sort.Slice(patterns, func(i, j int) bool {  // list内包含字典排序，参考https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
		return patterns[i].Priority < patterns[j].Priority
	})
	return patterns, nil
}

// func2 拼接某时刻的所有订单的信息，每个订单多个提醒模式
// todo 拆分函数： ProcessOrdersOfTime(预处理 + 调用SpliceOrders) + SpliceOrders
func ProcessOrdersOfTime(time string) {
	// 查询所有order表中时间等于0900的model，for这些model，判断model下是否至少有1个提醒条件满足；
	// 若有，则拼接提醒用语，放到队列
	// 以上操作可以考虑分批开goroutine
	order := orderModel.Order{}
	orderModelList, err := order.QueryListByTime(time)  // todo 测试：如果该时间下有脏数据，后面代码的健壮性
	if err != nil {
		log.Println(err.Error())
		return
	}
	session := common.Engine.NewSession()
	defer session.Close()
	for _, oneOrderModel := range orderModelList {
		// 预处理用户，判断是否有权限发送订单
		isAuth, err := processBeforeNotify(oneOrderModel.Creator)
		log.Println("isAuth", isAuth)
		log.Println("err", err)
		if err != nil {
			log.Println(err.Error())
			continue
		} else if !isAuth {
			log.Println("User has no auth")
			continue
		}

		patterns, err := SpliceOrders(oneOrderModel)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		// 拼接提醒用语，优先级，for循环后按照优先级排序，然后最终拼接用语，加到队列
		// [[有阵雨, 1], [最高气温较前一日增加5度，升至25度，注意防范, 2]]
		if len(patterns) >= 1{  // 需要提醒
			// 先把提醒次数减一
			userInfoFlexibleModel := &user.UserInfoFlexible{}
			userInfoFlexibleModel, has, err := userInfoFlexibleModel.QueryByUserId(oneOrderModel.Creator)
			if err != nil {
				err = errnum.New(errnum.DbError, err)
				log.Println("get user_info_flexible model error: ", err.Error())
				continue
			} else if !has {
				log.Println("user_info_flexible model not exist")
				continue
			}
			userInfoFlexibleModel.TodayTelRemindRemaining -= 1
			rowsAffected, err := session.AllCols().Where(
				"user_id=?", oneOrderModel.Creator).And(
				"today_tel_remind_remaining>=1").Update(*userInfoFlexibleModel)
			if err != nil {
				log.Println(err.Error())
			}
			if rowsAffected <= 0 {
				// 并发的时候，该用户的其他订单已经触发了提醒，并将提醒次数用尽
				log.Println("No enough today_tel_remind_remaining")
				continue
			}
			var tips string
			for _, pattern := range patterns {
				if tips == "" {
					tips += pattern.RemindSplicedText
				} else {
					tips = strings.Join([]string{tips, pattern.RemindSplicedText}, ";")
				}
			}
			needToRemindOrder := common.NeedToRemindOrder{
				//SubscriberId: oneOrderModel.UserId,
				CityCode:       oneOrderModel.RemindCity,
				SubscriberName: oneOrderModel.SubscriberName,
				TelephoneNum:   oneOrderModel.TelephoneNum,
				Creator:        oneOrderModel.Creator,
				Tips:           tips,
			}
			needToRemindOrderJson, err := json.Marshal(needToRemindOrder)
			if err != nil {
				log.Println(err.Error())
				panic(err)
			}
			log.Println("duandian needToRemindOrder begin")
			fmt.Println(needToRemindOrder)
			log.Println("duandian needToRemindOrder end")
			msg := message.NewMessage(watermill.NewUUID(), needToRemindOrderJson)  // 封装用户信息和带拼接的短信
			_ = msg
			// todo 测试的话，注释掉下面语句
			if err := common.PubSub.Publish("Topic.needToRemindOrder", msg); err != nil {
				panic(err)
			}
			fmt.Println()
		}
		fmt.Println("patternsxx", patterns)
		// 提醒： 明日 有阵雨，注意带伞；（优先级1）最高气温较前一日增加5度，升至25度，注意防范（优先级2）
	}

	// todo 后面考虑用remind_pattern的met_classification字段做switch case，需要改表：remind_object -> remind_object_id
}

// 每1分钟查看一次订单，将符合条件的放到队列里 - 子func
func cronOrderFunc() {
	// 得到当前的时间：精确到分钟，调用ProcessOrdersOfTime
	localDateStr := utils.GetLocalHourMin4Str()
	log.Println("localDateStr", localDateStr)
	go ProcessOrdersOfTime(localDateStr)
}
// func 每1分钟查看一次订单，将符合条件的放到队列里
// pubsub参考: https://studygolang.com/articles/26894
func CronOrder() {
	//c := utils.NewWithSeconds()  # 以前的用法
	c := cron.New()
	err := c.AddFunc("0 */1 * * * *", cronOrderFunc)  // 1分钟一次，且是整点
	//_, err := c.AddFunc("*/2 * * * * *", cronOrderFunc)    // 为了测试，2秒钟1次; 测试完务必打开开关
	if err != nil {
		fmt.Println(err.Error())
	}
	go c.Start()
	defer c.Stop()

	select {}
}


//  todo 接口函数要放到一个interface里吗？
/**
 * @Author Evan
 * @Description 新增手机号提醒订单，lib函数
	step1 params校验：手机号、城市校验、提醒时间校验
	step2 获取用户当前可配置的手机号订单数量，若不够，直接返回报错
	step3 将用户配置的order, orderDetail写到数据库中
step4 事务：order表、order_detail表、user_info_flexible表同时更新
 * @Date 18:41 2020-10-07
 * @Param 
 * @return 
 **/
func AddUserOrderTel(userId int, preTele string, telephone string, cityCode string, remindTime string,
	orderDetail []orderModel.OrderDetailItem) (map[string]interface{}, error){
	// todo 判断今天是否还有新增/编辑的次数，最后新建成功后要再将次数减一
	// step1 params校验：手机号、城市校验、提醒时间校验，提醒权限校验
	// step2 获取用户当前可配置的手机号订单数量，若不够，直接返回报错
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err := userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		println("get user_info_flexible model error: ", err.Error())
		return nil , err
	} else if !has {
		log.Println("user_info_flexible model not exist")
		return nil, errors.New("user_info_flexible model not exist")
	}
	telOrderRemaining := userInfoFlexibleModel.TelOrderRemaining
	if telOrderRemaining <= 0 {
		err = errnum.New(errnum.ErrTelOrderChanceInsufficient, nil)
		return nil, err
	}
	// step3 将用户配置的order表中（后面的其他步骤如果操作失败，手动删除刚添加的数据）
	currentTime := time.Now()
	orderModelToAdd := orderModel.Order{}
	orderModelToAdd.RemindCity = cityCode
	orderModelToAdd.RemindTime = remindTime
	orderModelToAdd.TelephoneNum = preTele + telephone
	orderModelToAdd.Creator = userId
	orderModelToAdd.CreateTime = currentTime
	orderModelToAdd.UpdateTime = currentTime
	err = orderModelToAdd.Create()
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		println("err", err.Error())
		return nil, err
	}
	orderId := orderModelToAdd.OrderId

	// step4 事务：order_detail表、user_info_flexible表同时更新
	session := common.Engine.NewSession()
	defer session.Close()
	if session.Begin() != nil {  // 事务开启
		err = errnum.New(errnum.DbError, nil)
		return nil, err
	}
	// 4.1 新增order_detail表
	for _, orderDetailItem := range orderDetail {
		orderDetailModelToAdd := &orderModel.OrderDetail{
			OrderId:         orderId,
			RemindPatternId: orderDetailItem.RemindPatternId,
			Value:           orderDetailItem.Value,
			CreateTime:      currentTime,
			UpdateTime:      currentTime,
		}
		if _, err = session.InsertOne(orderDetailModelToAdd); err != nil {
			err = errnum.New(errnum.DbError, err)
			return nil, err
		}
	}

	// 4.2 更新userInfoFlexibleModel
	userInfoFlexibleModel.TelOrderRemaining -= 1
	if _, err = session.AllCols().Where("user_id=?", userId).Update(*userInfoFlexibleModel); err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}

	if err = session.Commit(); err != nil {
		err = errnum.New(errnum.DbError, err)
		_ = session.Rollback()
		_ = orderModelToAdd.Delete()  // todo未来可以考虑在这里加上err，并且加上log，打上断点
		return nil, err
	}
	resultData := map[string]interface{} {
		"result": "success",
	}
	return resultData, nil
}

/**
 * @Author Evan
 * @Description 查询用户手机号订单，lib函数
	step1: 查询用户order表（by creator）和orderDetail表（by order_id）
	step2: 拼接返回值，返回
 * @Date 10:13 2020-11-03
 * @Param
 * @return
 **/
func GetUserOrderTel(userId int) (map[string]interface{}, error){
	orderModelInstance := orderModel.Order{}
	orderModelList, err := orderModelInstance.QueryListByCreator(userId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("err: " + err.Error())
		return nil, err
	}

	resOrderAndDetailList := make([]orderModel.ResOrderAndDetail, 0)
	for _, oneOrderModel := range orderModelList {
		// 这里和libs/order/order.go ProcessOrdersOfTime方法差不多；
		var resOrderAndDetail = new(orderModel.ResOrderAndDetail)
		// PreTele, CityName未来做，其中CityName做map
		resOrderAndDetail.OrderId = oneOrderModel.OrderId
		resOrderAndDetail.Telephone = oneOrderModel.TelephoneNum  // todo 现在数据库中存的还是+86
		resOrderAndDetail.CityCode = oneOrderModel.RemindCity  // city code
		resOrderAndDetail.RemindTime = oneOrderModel.RemindTime

		// 先查数据库
		orderDetailInstance := orderModel.OrderDetail{}
		orderDetailList, err := orderDetailInstance.QueryListByOrderId(oneOrderModel.OrderId)
		if err != nil {
			err = errnum.New(errnum.DbError, err)
			log.Println("err: " + err.Error())
			return nil, err
		}

		for _, oneOrderDetailModel := range orderDetailList {
			var resOrderDetailItem = new(orderModel.ResOrderDetailItem)
			resOrderDetailItem.Value = oneOrderDetailModel.Value
			resOrderDetailItem.RemindPatternId = oneOrderDetailModel.RemindPatternId
			resOrderAndDetail.OrderDetail = append(resOrderAndDetail.OrderDetail, *resOrderDetailItem)
		}
		resOrderAndDetailList = append(resOrderAndDetailList, *resOrderAndDetail)
	}
	resultData := map[string]interface{}{
		"orders": resOrderAndDetailList,
	}
	// todo 测试空的情况；写测试用例?
	return resultData, nil
}

func UpdateUserOrderTel(resOrderAndDetail orderModel.ResOrderAndDetail, userId int) (map[string]interface{}, error) {
	/*
	step1 判断用户今天是否还有编辑的次数
	step2 根据order_id 查该order信息，判断该order是否属于该用户；若属于才进行下一步
	step3 参数校验
	step4 事务：修改order，删除旧order_detail，新增新order_detail，提交
	step5 返回
	 */
	// todo 之后和AddUserOrderTel尝试合并（两个表的增加部分合并；更改前做md5校验，前后的订单信息做md5，若无变化则不改）
	// step1 判断用户今天是否还有编辑的次数
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err := userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		println("get user_info_flexible model error: ", err.Error())
		return nil , err
	} else if !has {
		println("user_info_flexible model not exist")
		return nil, errors.New("user_info_flexible model not exist")
	}
	telOrderRemaining := userInfoFlexibleModel.TelOrderRemaining
	if telOrderRemaining <= 0 {
		err = errnum.New(errnum.ErrTelOrderChanceInsufficient, nil)
		return nil, err
	}
	// step2 根据order_id 查该order信息，判断该order是否属于该用户
	orderId := resOrderAndDetail.OrderId
	orderModelInstance := orderModel.Order{}
	theOrderModel, has, err := orderModelInstance.QueryOneByOrderId(orderId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("err: " + err.Error())
		return nil, err
	} else if !has {
		err = errnum.New(errnum.ErrOrderNotFound, nil)
		log.Println("err: " + err.Error())
		return nil, err
	}
	if userId != theOrderModel.Creator {
		err = errnum.New(errnum.ErrNoAuth, nil)
		log.Println("err: " + err.Error())
		return nil, err
	}

	// step3 参数校验 todo

	// step4 事务：修改order，删除旧order_detail，新增新order_detail，提交
	session := common.Engine.NewSession()
	defer session.Close()
	if session.Begin() != nil {  // 事务开启
		err = errnum.New(errnum.DbError, nil)
		log.Println("err: " + err.Error())
		return nil, err
	}
	currentTime := time.Now()
	// 4.1 order表修改
	theOrderModel.RemindCity = resOrderAndDetail.CityCode
	theOrderModel.RemindTime = resOrderAndDetail.RemindTime
	theOrderModel.TelephoneNum = resOrderAndDetail.PreTele + resOrderAndDetail.Telephone
	theOrderModel.UpdateTime = currentTime
	if _, err = session.AllCols().Where("order_id=?", orderId).Update(*theOrderModel); err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}

	// 4.2 旧order_detail删除
	_, err = session.Where("order_id=?", orderId).Delete(orderModel.OrderDetail{})
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("err: " + err.Error())
		return nil, err
	}

	// 4.3 新order_detail新增
	for _, orderDetailItem := range resOrderAndDetail.OrderDetail {
		orderDetailModelToAdd := &orderModel.OrderDetail{
			OrderId:         orderId,
			RemindPatternId: orderDetailItem.RemindPatternId,
			Value:           orderDetailItem.Value,
			CreateTime:      currentTime,
			UpdateTime:      currentTime,
		}
		if _, err = session.InsertOne(orderDetailModelToAdd); err != nil {
			err = errnum.New(errnum.DbError, err)
			return nil, err
		}
	}

	if err = session.Commit(); err != nil {
		err = errnum.New(errnum.DbError, err)
		_ = session.Rollback()
	}
	resultData := map[string]interface{} {
		"result": "success",
	}
	return resultData, nil
}


/**
 * @Author Evan
 * @Description 用户删除某订单（真删）
	step1: 根据order_id查创建者，若不属于该用户，报错
	step2: 查询该order_id关联的order_detail表
	step3: 事务：order表和order_detail一起删除，同时更新用户的可配置订单数
 * @Date 18:00 2020-11-16
 * @Param
 * @return
 **/
// todo delete的时候加上用户的次数可配置次数
// todo 后端做防连点？ （前端先做放连点）
func DeleteUserOrderTel(orderId int, userId int) (map[string]interface{}, error) {
	orderModelInstance := orderModel.Order{}
	theOrderModel, has, err := orderModelInstance.QueryOneByOrderId(orderId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("err: " + err.Error())
		return nil, err
	} else if !has {
		err = errnum.New(errnum.ErrOrderNotFound, nil)
		log.Println("err: " + err.Error())
		return nil, err
	}

	// 待删除订单不是本人创建
	if userId != theOrderModel.Creator {
		err = errnum.New(errnum.ErrNoAuth, nil)
		log.Println("err: " + err.Error())
		return nil, err
	}

	// delete
	// 事务：order和order_detail一起删除
	session := common.Engine.NewSession()
	defer session.Close()
	if session.Begin() != nil {  // 事务开启
		err = errnum.New(errnum.DbError, nil)
		return nil, err
	}
	_, err = session.Where("order_id=?", orderId).Delete(*theOrderModel)
	if err != nil {
		log.Println("err: " + err.Error())
		return nil, err
	}
	_, err = session.Where("order_id=?", orderId).Delete(orderModel.OrderDetail{})
	if err != nil {
		err = errnum.New(errnum.DbError, nil)
		log.Println("err: " + err.Error())
		return nil, err
	}
	// 用户可配置订单加一（todo这里需要做防连点！）
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err = userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		err = errnum.New(errnum.DbError, nil)
		println("get user_info_flexible model error: ", err.Error())
		return nil , err
	} else if !has {
		log.Println("user_info_flexible model not exist")
		return nil, errors.New("user_info_flexible model not exist")
	}
	userInfoFlexibleModel.TelOrderRemaining += 1
	if _, err = session.AllCols().Where("user_id=?", userId).Update(*userInfoFlexibleModel); err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}
	err = session.Commit()
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("err: " + err.Error())
		_ = session.Rollback()
		return nil, err
	}
	resultData := map[string]interface{}{
		"result": "success",
	}
	return resultData, nil
}














