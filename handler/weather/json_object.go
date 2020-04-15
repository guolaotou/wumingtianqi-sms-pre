package weather

// 存放json结构体
type XinZhiWeatherDailyResults struct {
	Results []XinZhiWeatherDailyResultsItem `json:"results"`
}

type XinZhiWeatherDailyResultsItem struct {
	Location   Location                 `json:"location"`
	Daily      []XinZhiWeatherDailyItem `json:"daily"`
	LastUpdate string                   `json:"last_update"`
}

type Location struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	Path           string `json:"path"`
	Timezone       string `json:"timezone"`
	TimezoneOffset string `json:"timezone_offset"`
}

type XinZhiWeatherDailyItem struct {
	Date                string `json:"date"`
	TextDay             string `json:"text_day"`
	CodeDay             string `json:"code_day"`
	TextNight           string `json:"text_night"`
	CodeNight           string `json:"code_night"`
	High                string `json:"high"`
	Low                 string `json:"low"`
	Rainfall            string `json:"rainfall"`
	Precip              string `json:"precip"`
	WindDirection       string `json:"wind_direction"`
	WindDirectionDegree string `json:"wind_direction_degree"`
	WindSpeed           string `json:"wind_speed"`
	WindScale           string `json:"wind_scale"`
	Humidity            string `json:"humidity"`
}
