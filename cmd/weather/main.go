package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Coord      CoordData   `json:"coord"`
	Weather    WeatherData `json:"weather"`
	Base       string      `json:"base"`
	Main       MainData    `json:"main"`
	Visibility string      `json:"visibility"`
	Wind       WindData    `json:"wind"`
	Clouds     CloudsData  `json:"clouds"`
	Dt         float64     `json:"dt"`
	Sys        SysData     `json:"sys"`
	Timezone   int32       `json:"timezon"`
	Id         float64     `json:"id"`
	Name       string      `json:"name"`
	Cod        int32       `json:"cod"`
}

type WeatherData struct {
	Id          int32
	Main        string
	Description string
	Icon        string
}

type MainData struct {
	Temp       float32
	Feels_like float32
	Temp_min   float32
	Temp_max   float32
	Pressure   float32
	Humidity   float32
}

type WindData struct {
	Speed float32
	Deg   float32
}

type CloudsData struct {
	All int32
}

type SysData struct {
	Type_   int32
	Id      int64
	Country string
	Sunrise float64
	Sunset  float64
}
type CoordData struct {
	Lon float32
	Lat float32
}

type ConfigData struct {
	AppId    string
	Location string
	Units    string
}

// shows weather from CLI, no browser/phone distraction
func main() {

	apiKeyFile := ".weather_appid.json"
	homedir := os.Getenv("HOME")
	appId, err := os.ReadFile(homedir + "/" + apiKeyFile)
	if err != nil {
		panic(err)
	}
	var conf ConfigData
	json.Unmarshal(appId, &conf)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s",
		conf.Location, conf.Units, conf.AppId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var weather WeatherResponse

	json.Unmarshal([]byte(body), &weather)
	fmt.Printf("Location=%s Temp:%.2f Clouds:%d%% Wind:%.2f",
		conf.Location, weather.Main.Temp, weather.Clouds.All, weather.Wind.Speed)
}
