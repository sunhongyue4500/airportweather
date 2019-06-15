package main

import (
	"airportweather/config"
	"airportweather/db"
	"airportweather/defs"
	"airportweather/mq"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// 从民航气象中心爬取气象数据,爬取成功后写入MongoDB数据库并推送到消息队列RabbitMQ
func MakeRequest() {
	form := url.Values{}
	form.Add("cmd", "GetIndexAirportInformation")
	resp, err := http.PostForm(config.FetchURL, form)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string][]defs.AirportWeather
	json.NewDecoder(resp.Body).Decode(&result)

	temp := result["Data"]
	tempSlice := make([]interface{}, len(temp))
	for idx, val := range temp {
		tempSlice[idx] = val
	}
	log.Printf("获取%d条气象数据\n", len(tempSlice))

	// 写入数据库
	// 先删除
	if err := db.DropCollection(); err != nil {
		log.Fatalln(err)
	}
	// 再插入
	if _, err := db.InsertManyDocs(tempSlice); err != nil {
		log.Fatalln(err)
	}
	slcB, _ := json.Marshal(temp)
	// 推送到消息队列
	mq.Send2MQ([]byte(string(slcB)))
}
