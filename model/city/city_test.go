package city

import (
	"testing"
	"wumingtianqi/model/city"
	"wumingtianqi/model/common"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test model/city/city_test.go
// go test -v model/city/city_test.go

func TestCity(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	cityCode := "WX4FBXXFKE4F"
	c := &city.City{
		Id:       2,
		Province: "福建省",
		City:     "厦门市",
		District: "思明区",
		PinYin:   "Siming",
		Abbr:     "福建/厦门/思明",
		Code:     cityCode,
	}
	t.Log("*** begin create session****** ")

	if err := c.Create(); err != nil {
		panic(err)
	}

	// 2. 查询
	t.Log("*** begin query session****** ")

	c2, has, err := city.QueryByCityCode(cityCode)
	if err != nil || !has {
		t.Error("city not found")
	} else {
		t.Log("city: ", c2)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	t.Log("*** begin update session****** ")

	c2.Abbr = "福建/厦门/思明思明"
	c2.Update()
	c3, _, _ := city.QueryByCityCode(cityCode)
	t.Log("city: ", c3)
	t.Log("*** end update session****** ")

	// 4. 删除
	t.Log("*** begin delete session****** ")
	c3.Delete()
	t.Log("*** end delete session****** ")
}
