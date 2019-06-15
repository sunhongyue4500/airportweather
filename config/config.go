package config

import (
	"airportweather/db"
	"airportweather/defs"
	"airportweather/mq"
	"encoding/json"
	"io/ioutil"
	"log"
)

var conf *defs.QiXiangConfig
var FetchURL string

func init() {
	// 针对不同的
	result, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(result))
	conf = &defs.QiXiangConfig{}
	err = json.Unmarshal([]byte(result), conf)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v", conf)
	FetchURL = conf.FetchURL

	// 连接数据库
	err = db.ConnMongoDB(conf.QiXiangMongoConfig.MongoIP, conf.QiXiangMongoConfig.MongoPort)
	if err != nil {
		log.Fatalf("db connect error:%v\n", err)
	}
	// 连接mq
	err = mq.ConnMQ(conf.QiXiangRabbitMQConfig.MQLoginUser, conf.QiXiangRabbitMQConfig.MQLoginPwd, conf.QiXiangRabbitMQConfig.MQHost, conf.QiXiangRabbitMQConfig.MQPort)
	if err != nil {
		log.Fatalf("mq connect error:%v\n", err)
	}
}
