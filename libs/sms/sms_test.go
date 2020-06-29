package sms

import (
	"testing"
	"wumingtianqi-sms-pre/config"
	smsLib "wumingtianqi-sms-pre/libs/sms"
	test "wumingtianqi-sms-pre/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
//  go test libs/sms/sms_test.go
//  go test -v libs/sms/sms_test.go

func TestSms(t *testing.T) {
	test.Setup()
	smsConfig := config.GlobalConfig.Sms

	// 秘钥填写
	smsSdk := smsLib.SmsSdk{
		SecretId:    smsConfig.SecretId,
		SecretKey:   smsConfig.SecretKey,
		SmsSdkAppId: smsConfig.SmsSdkAppId,
		Sign:        smsConfig.Sign,
	}
	//TemplateParamSet := []string{"海淀区", "这是2；", "这是3；", "这是4；", "这是5；", "这是6。"}
	TemplateParamSet := []string{"海淀区", "22222", "33333", "44444", "5", "6"}
	TemplateID := "634989"
	PhoneNumberSet := smsConfig.TestPhone

	smsSdk.SmsSdk(TemplateParamSet, TemplateID, PhoneNumberSet)
}
