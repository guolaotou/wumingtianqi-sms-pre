package weather

type WeatherFlagContent struct {

}

// 天气信息类别对照表
type WeatherFlag struct {
	Id int `json:"id" xorm:"pk autoincr INT(11)"`
	Flag string `json:"flag" xorm:"VARCHAR(30)"`
	Content WeatherFlagContent `json:"content" xorm:"json"`  // todo alter table wumingtianqi.weather_flag change content content  json;
}