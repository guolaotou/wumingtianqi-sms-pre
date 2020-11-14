package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/robfig/cron"
	"log"
	"time"
	"wumingtianqi/model/user"
	"sort"
	"strconv"
	"strings"
	"wumingtianqi/config"
	"wumingtianqi/model/common"
	orderModel "wumingtianqi/model/order"
	"wumingtianqi/model/remind"
	weatherModel "wumingtianqi/model/weather"
	"wumingtianqi/utils"
	"wumingtianqi/utils/errnum"
)

// todo 用户建立订单
// req 有一些参数
// 中间处理过程比较麻烦？
// 然后存到订单表

// todo what I really want to do
// 1. 把符合提醒要求的订单放到队列（生产者）
// 1.1 拉取昨日，今日的天气表，存到map（以后做到缓存里）,涉及map？ 日后redis
// 1.2 订单表拉取9:00的订单
// 2. 拼接n个订单的信息，每个订单有多个提醒模式
// 3. 规范第2步：取队列的代码
// 20200530

// func1 假数据，天气信息, todo 放到weather 天气里
type WeatherItem map[string]interface{}
type Weather map[string]WeatherItem

func FakeWeather() (Weather, Weather) {  // todo 之后把天气信息弄成真的，服务器装tmux跑数据
	yesterdayWeather := Weather{
		"Beijing": WeatherItem{
			"city":      "Beijing",
			"code_text": "晴",
			"code_id":   1,
			"high":      20,
		},
	}
	todayWeather := Weather{
		"Beijing": WeatherItem{
			"city":      "Beijing",
			"code_text": "阵雨",
			"code_id":   10,
			"high":      25,
		},
	}
	return yesterdayWeather, todayWeather
}

func Weather2Map(city string) (Weather, Weather) {
	// 计算今日时间
	yesterdayDate8Int := utils.GetSpecificDate8Int(0)  // todo 这个之后需要协调today和yestoday
	yesterdayWeatherAll, _, _ := weatherModel.QueryByCityDate(city, yesterdayDate8Int)
	yesterdayWeather := Weather{
		city: WeatherItem{
			"city": city,
			"code_text": yesterdayWeatherAll.TextDay,  // todo 以后考虑把白天天气和晚上天气统一
			"code_id": yesterdayWeatherAll.CodeDay,
			"high": yesterdayWeatherAll.High,
		},
	}

	todayDate8Int := utils.GetSpecificDate8Int(1)  // todo 这个之后需要协调today和yestoday
	todayWeatherAll, _, _ := weatherModel.QueryByCityDate(city, todayDate8Int)
	todayWeather := Weather{
		city: WeatherItem{
			"city": city,
			"code_text": todayWeatherAll.TextDay,
			"code_id": todayWeatherAll.CodeDay,
			"high": todayWeatherAll.High,
		},
	}
	return yesterdayWeather, todayWeather
}

type SplicePatternModel struct {
	RemindSplicedText string `json:"remind_spliced_text"` // 拼接好的语句
	Priority          int    `json:"priority"`
}

//
//func splicePattern1(city string, remindPattern *remind.RemindPattern) SplicePatternModel{
//	// 1. 突然降雨
//	// 枚举"天气现象"表，整理突然降雨的触发条件 ![1,2,3] -> [1,2,3]，考虑remind_pattern里新增一个extension字段（json格式），这个字段不同业务不一样，需要的东西也不一样。
//	// 降雨对应的id todo 以后再弄个天气代码映射表？或者在某个地方弄个静态变量存
//	RainPatternIds := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
//
//	var yesterdayWeather, todayWeather Weather
//	if config.GlobalConfig.Weather.FakeData {
//		yesterdayWeather, todayWeather = FakeWeather()
//	} else {
//		yesterdayWeather, todayWeather = Weather2Map(city)
//	}
//
//	codeText := todayWeather[city]["code_text"].(string)
//	codeYesterday := yesterdayWeather[city]["code_id"].(int)
//	codeToday := todayWeather[city]["code_id"].(int)
//	isYesRain, _ := utils.IsContain(codeYesterday, RainPatternIds)
//	isTodayRain, _ := utils.IsContain(codeToday, RainPatternIds)
//
//	var pattern = new(SplicePatternModel)
//	if !isYesRain && isTodayRain {
//		pattern.RemindSplicedText = "有" + codeText + "记得带伞"  // 有阵雨 todo以后用format_text做通配，封装写法
//		pattern.Priority = remindPattern.PriorityRemind
//	}
//	return *pattern
//}

func splicePattern1(city string, remindPattern *remind.RemindPattern) SplicePatternModel{
	// 1. 降水天气
	// 枚举"天气现象"表，整理降水的所有情况 [1,2,3]，考虑remind_pattern里新增一个extension字段（json格式），这个字段不同业务不一样，需要的东西也不一样。
	// 降雨对应的id todo 以后再弄个天气代码映射表？或者在某个地方弄个静态变量存
	RainPatternIds := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var  todayWeather Weather
	if config.GlobalConfig.Weather.FakeData {
		_, todayWeather = FakeWeather()
	} else {
		_, todayWeather = Weather2Map(city)
	}

	codeText := todayWeather[city]["code_text"].(string)
	codeToday := todayWeather[city]["code_id"].(int)
	isTodayRain, _ := utils.IsContain(codeToday, RainPatternIds)

	var pattern = new(SplicePatternModel)
	if isTodayRain {
		pattern.RemindSplicedText = "有" + codeText + "记得带伞"  // 有阵雨 todo以后用format_text做通配，封装写法
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern
}

func splicePattern2(city string, remindPattern *remind.RemindPattern, value int) SplicePatternModel {
	// 2. 突然升温
	var yesterdayWeather, todayWeather Weather
	if config.GlobalConfig.Weather.FakeData {
		yesterdayWeather, todayWeather = FakeWeather()
	} else {
		yesterdayWeather, todayWeather = Weather2Map(city)
	}
	highYesterday := yesterdayWeather[city]["high"].(int)
	highToday := todayWeather[city]["high"].(int)
	highTodayStr := strconv.Itoa(todayWeather[city]["high"].(int))
	valueStr := strconv.Itoa(highToday - highYesterday)

	var pattern = new(SplicePatternModel)
	log.Println("highToday - highYesterday", highToday - highYesterday)
	log.Println("valuevalue", value)
	if highToday - highYesterday >= value {  // 最高气温较前一日增加5度，升至25度，注意防范
		remindObject := remindPattern.RemindObject
		pattern.RemindSplicedText = remindObject + "较前一日增加" + valueStr + "度，升至" + highTodayStr + "度，注意防范"
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern
}

func splicePattern3(city string, remindPattern *remind.RemindPattern, value int) SplicePatternModel {
	// 3. 突然降温
	var yesterdayWeather, todayWeather Weather
	if config.GlobalConfig.Weather.FakeData {
		yesterdayWeather, todayWeather = FakeWeather()
	} else {
		yesterdayWeather, todayWeather = Weather2Map(city)
	}
	highYesterday := yesterdayWeather[city]["high"].(int)
	highToday := todayWeather[city]["high"].(int)
	highTodayStr := strconv.Itoa(todayWeather[city]["high"].(int))
	valueStr := strconv.Itoa(highToday - highYesterday)

	var pattern = new(SplicePatternModel)
	log.Println("highToday - highYesterday", highToday -highYesterday)
	log.Println("valuevalue", value)
	if highYesterday- highToday >= value {  // 最高气温较前一日降低5度，降到18度，注意防范
		remindObject := remindPattern.RemindObject
		pattern.RemindSplicedText = remindObject + "较前一日降低" + valueStr + "度，降至" + highTodayStr + "度，注意防范"
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern
}

func splicePattern8(city string, remindPattern *remind.RemindPattern) SplicePatternModel{
	// 8. 雨过天晴
	// 枚举"天气现象"表，整理突然降雨的触发条件 ![1,2,3] -> [1,2,3]，考虑remind_pattern里新增一个extension字段（json格式），这个字段不同业务不一样，需要的东西也不一样。
	// 降雨对应的id todo 以后再弄个天气代码映射表？或者在某个地方弄个静态变量存
	RainPatternIds := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var yesterdayWeather, todayWeather Weather
	if config.GlobalConfig.Weather.FakeData {
		yesterdayWeather, todayWeather = FakeWeather()
	} else {
		yesterdayWeather, todayWeather = Weather2Map(city)
	}

	codeText := todayWeather[city]["code_text"].(string)
	codeYesterday := yesterdayWeather[city]["code_id"].(int)
	codeToday := todayWeather[city]["code_id"].(int)
	isYesRain, _ := utils.IsContain(codeYesterday, RainPatternIds)
	isTodayRain, _ := utils.IsContain(codeToday, RainPatternIds)

	var pattern = new(SplicePatternModel)
	if isYesRain && !isTodayRain {
		pattern.RemindSplicedText = "没有降水啦！明天天气：" + codeText // 雨过天晴
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern
}

// func2 拼接某时刻的所有订单的信息，每个订单多个提醒模式
func SpliceOrders(time string) {
	// 查询所有order表中时间等于0900的model，for这些model，判断model下是否至少有2个提醒条件满足；
	// 若有，则拼接提醒用语，放到队列
	// 以上操作可以考虑分批开goroutine
	order := orderModel.Order{}
	orderModelList, err := order.QueryListByTime(time)  // todo 测试：如果该时间下有脏数据，后面代码的健壮性
	if err != nil {
		panic(err)
	}
	for _, oneOrderModel := range orderModelList {
		orderId := oneOrderModel.OrderId
		city := oneOrderModel.RemindCity

		// 根据order_id找到order_detail
		orderDetail := orderModel.OrderDetail{}
		orderDetailList, err := orderDetail.QueryListByOrderId(orderId)
		if err != nil {
			panic(err)
		}

		// 定义一个model，用来接单个提醒
		patterns := make([]SplicePatternModel, 0)
		for _, oneOrderDetailModel := range orderDetailList {
			value := oneOrderDetailModel.Value

			remindPattern := new(remind.RemindPattern)
			remindPattern, _, _ = remindPattern.QueryOneById(oneOrderDetailModel.RemindPatternId)

			switch oneOrderDetailModel.RemindPatternId {
			case 1: // 突然降雨
				pattern1 := splicePattern1(city, remindPattern)
				if pattern1.Priority >= 1 {  // 以后可以用标准一点的用法
					patterns = append(patterns, pattern1)
				}
			case 2: // 突然升温
				pattern2 := splicePattern2(city, remindPattern, value)
				if pattern2.Priority >= 1 {  // 以后可以用标准一点的用法
					patterns = append(patterns, pattern2)
				}
			case 3: // 突然降温
				pattern3 := splicePattern3(city, remindPattern, value)
				if pattern3.Priority >= 1 {  // 以后可以用标准一点的用法
					patterns = append(patterns, pattern3)
				}
			case 4: // 空气质量变差
				println(2)
			case 5: // 9点突然升温
				println(2)
			case 6: // 高温预警
				println(2)
			case 7: // 低温预警
				println(2)
			case 8: // 雨过天晴
				pattern8 := splicePattern8(city, remindPattern)
				if pattern8.Priority >= 1 {  // 以后可以用标准一点的用法
					patterns = append(patterns, pattern8)
				}
			}
		}
		sort.Slice(patterns, func(i, j int) bool {  // list内包含字典排序，参考https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
			return patterns[i].Priority < patterns[j].Priority
		})
		// 拼接提醒用语，优先级，for循环后按照优先级排序，然后最终拼接用语，加到队列
		// [[有阵雨, 1], [最高气温较前一日增加5度，升至25度，注意防范, 2]]
		if len(patterns) >= 1{  // 需要提醒
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
				City:           city,
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
			fmt.Println(needToRemindOrder)
			msg := message.NewMessage(watermill.NewUUID(), needToRemindOrderJson)  // 封装用户信息和带拼接的短信
			if err := common.PubSub.Publish("Topic.needToRemindOrder", msg); err != nil {
				panic(err)
			}
		}
		fmt.Println("patterns", patterns)
		// 提醒： 明日 有阵雨，注意带伞；（优先级1）最高气温较前一日增加5度，升至25度，注意防范（优先级2）
	}

	// todo 后面考虑用remind_pattern的met_classification字段做switch case，需要改表：remind_object -> remind_object_id
}

// 每1分钟查看一次订单，将符合条件的放到队列里 - 子func
func cronOrderFunc() {
	// 得到当前的时间：精确到分钟，调用SpliceOrders
	localDateStr := utils.GetLocalHourMin4Str()
	log.Println("localDateStr", localDateStr)
	go SpliceOrders(localDateStr)
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
func AddUserOrderTel(userId int, telephone string, city string, remindTime string,
	orderDetail []orderModel.OrderDetailItem) (map[string]interface{}, error){
	// step1 params校验：手机号、城市校验、提醒时间校验，提醒权限校验
	// step2 获取用户当前可配置的手机号订单数量，若不够，直接返回报错
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err := userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		println("get user_info_flexible model error: ", err.Error())
		return nil , err
	} else if !has {
		println("user_info_flexible model not exist")
		return nil, errors.New("user_info_flexible model not exist")
	} else {
		println("model: ", userInfoFlexibleModel)
	}
	telOrderRemaining := userInfoFlexibleModel.TelOrderRemaining
	if telOrderRemaining <= 0 {
		err = errnum.New(errnum.ErrTelOrderChanceInsufficient, nil)
		return nil, err
	}
	// step3 将用户配置的order表中（后面的其他步骤如果操作失败，手动删除刚添加的数据）
	currentTime := time.Now()
	orderModelToAdd := orderModel.Order{}
	orderModelToAdd.RemindCity = city
	orderModelToAdd.RemindTime = remindTime
	orderModelToAdd.TelephoneNum = telephone
	orderModelToAdd.Creator = -1  // todo 更新这个
	orderModelToAdd.CreateTime = currentTime
	orderModelToAdd.UpdateTime = currentTime
	err = orderModelToAdd.Create()
	if err != nil {
		err = errnum.New(errnum.DbError, nil)
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
			err = errnum.New(errnum.DbError, nil)
			return nil, err
		}
	}

	// 4.2 更新userInfoFlexibleModel
	userInfoFlexibleModel.TelOrderRemaining -= 1
	userInfoFlexibleModel.TodayEditChanceRemaining -= 1
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
	resultData := map[string]interface{}{
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
		return nil, err
	}

	resOrderAndDetailList := make([]orderModel.ResOrderAndDetail, 0)
	for _, oneOrderModel := range orderModelList {
		// 这里和libs/order/order.go SpliceOrders方法差不多；
		var resOrderAndDetail = new(orderModel.ResOrderAndDetail)
		// PreTele, CityName未来做，其中CityName做map
		resOrderAndDetail.OrderId = oneOrderModel.OrderId
		resOrderAndDetail.Telephone = oneOrderModel.TelephoneNum  // todo 现在数据库中存的还是+86
		resOrderAndDetail.City = oneOrderModel.RemindCity
		resOrderAndDetail.RemindTime = oneOrderModel.RemindTime

		// 先查数据库
		orderDetailInstance := orderModel.OrderDetail{}
		orderDetailList, err := orderDetailInstance.QueryListByOrderId(oneOrderModel.OrderId)
		if err != nil {
			err = errnum.New(errnum.DbError, err)
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

// todo 改
func UpdateUserOrderTel(resOrderAndDetail orderModel.ResOrderAndDetail, userId int) (map[string]interface{}, error) {
	// todo 添加注释？
	/*
	step1 根据order_id 查该order信息，判断该order是否属于该用户；若属于才进行下一步
	step2 参数校验
	step3 修改该order信息，提交
	step4 返回
	 */
	return nil, nil
}

// todo 删除
func DeleteUserOrderTel(orderId int) (error) {
	/* todo写注释
	step1: 根据order_id查创建者，若属于该用户，则删除
	step2: 返回
	 */
	orderModelInstance := orderModel.Order{}
	theOrderModel, has, err := orderModelInstance.QueryOneByOrderId(orderId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		// todo log
		return err
	} else if !has {
		err = errnum.New(errnum.ErrOrderNotFound, nil)
	}

	// delete
	err = theOrderModel.Delete()
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		// todo log
		return err
	}
	return nil
}














