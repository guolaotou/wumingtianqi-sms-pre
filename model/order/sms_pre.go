package order

type Model struct {
	UserId int    `json:"user_id"`
	Value  string `json:"value"`
}

// Q1，用来存订单
var Queue1 = map[string][]Model{}

// Q2，用来存要发短信的订单
var queue2 = map[string][]Model{}

// Q3，所有要发的短信内容放进这个队列
var queue3 = map[string][]Model{}
