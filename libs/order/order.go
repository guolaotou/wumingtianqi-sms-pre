package order

import (
	"fmt"
	"github.com/bradfitz/slice"
	"strconv"
	orderModel "wumingtianqi-sms-pre/model/order"
	"wumingtianqi-sms-pre/model/remind"
	"wumingtianqi-sms-pre/utils"
)

// todo 改package

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

type SplicePatternModel struct {
	RemindSplicedText string `json:"remind_spliced_text"` // 拼接好的语句
	Priority          int    `json:"priority"`
}


func splicePattern1(city string, remindPattern *remind.RemindPattern) SplicePatternModel{
	// 1. 突然降雨

	// 枚举"天气现象"表，整理突然降雨的触发条件 ![1,2,3] -> [1,2,3]，考虑remind_pattern里新增一个extension字段（json格式），这个字段不同业务不一样，需要的东西也不一样。
	// 降雨对应的id todo 以后再弄个天气代码映射表？或者在某个地方弄个静态变量存
	RainPatternIds := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	yesterdayWeather, todayWeather := FakeWeather()
	codeText := todayWeather[city]["code_text"].(string)
	codeYesterday := yesterdayWeather[city]["code_id"].(int)
	codeToday := todayWeather[city]["code_id"].(int)
	isYesRain, _ := utils.IsContain(codeYesterday, RainPatternIds)
	isTodayRain, _ := utils.IsContain(codeToday, RainPatternIds)

	var pattern = new(SplicePatternModel)
	if !isYesRain && isTodayRain {
		pattern.RemindSplicedText = "有" + codeText + "记得带伞"  // 有阵雨 todo以后用format_text做通配，封装写法
		pattern.Priority = remindPattern.PriorityRemind
	}
	return *pattern
}

func splicePattern2(city string, remindPattern *remind.RemindPattern, value int) SplicePatternModel {
	// 2. 突然升温

	yesterdayWeather, todayWeather := FakeWeather()
	highYesterday := yesterdayWeather[city]["high"].(int)
	highToday := todayWeather[city]["high"].(int)
	highTodayStr := strconv.Itoa(todayWeather[city]["high"].(int))
	valueStr := strconv.Itoa(value)

	var pattern = new(SplicePatternModel)
	if highToday - highYesterday >= value {  // 最高气温较前一日增加5度，升至25度，注意防范
		remindObject := remindPattern.RemindObject
		pattern.RemindSplicedText = remindObject + "较前一日增加" + valueStr + "度，升至" + highTodayStr + "度，注意防范"
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
	orderModelList, err := order.QueryListByTime(time)
	if err != nil {
		panic(err)
	}
	for _, oneOrderModel := range orderModelList {
		order_id := oneOrderModel.OrderId
		//user_id := oneOrderModel.UserId
		city := oneOrderModel.RemindCity

		// 根据order_id找到order_detail
		orderDetail := orderModel.OrderDetail{}
		orderDetailList, err := orderDetail.QueryListByOrderId(order_id)
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
				pattern1 := splicePattern2(city, remindPattern, value)
				if pattern1.Priority >= 1 {  // 以后可以用标准一点的用法
					patterns = append(patterns, pattern1)
				}
			case 3: // 突然降温
				println(2)
			case 4: // 空气质量变差
				println(2)
			case 5: // 9点突然升温
				println(2)
			case 6: // 高温预警
				println(2)
			case 7: // 低温预警
				println(2)
			}
		}
		slice.Sort(patterns[:], func(i, j int) bool {
			return patterns[i].Priority < patterns[j].Priority
		})
		// 拼接提醒用语，优先级，for循环后按照优先级排序，然后最终拼接用语，加到队列
		// [[有阵雨, 1], [最高气温较前一日增加5度，升至25度，注意防范, 2]]
		fmt.Println("patterns", patterns)
		// 提醒： 明日 有阵雨，注意带伞；（优先级1）最高气温较前一日增加5度，升至25度，注意防范（优先级2）
		// or 提醒: 明日有阵雨；（优先级1）最高气温较前一日增加5度，升至25度，注意带伞,注意防范（优先级2）
	}

	// todo 后面考虑用remind_pattern的met_classification字段做switch case，需要改表：remind_object -> remind_object_id
}