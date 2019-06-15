package defs

// RabbitMQ配置
type QiXiangRabbitMQConfig struct {
	MQHost      string `json:"host"`
	MQPort      int    `json:"port"`
	MQLoginUser string `json:"loginuser"`
	MQLoginPwd  string `json:"loginpwd"`
}

// MongoDB配置
type QiXiangMongoConfig struct {
	MongoIP   string `json:"ip"`
	MongoPort int    `json:"port"`
}

// 程序配置 `json:"fetchURL"` :号后面不能有空格
type QiXiangConfig struct {
	FetchURL              string                `json:"fetchURL"`
	QiXiangMongoConfig    QiXiangMongoConfig    `json:"mongodb"`
	QiXiangRabbitMQConfig QiXiangRabbitMQConfig `json:"rabbitmq"`
}

// airprot weather
type AirportWeather struct {
	AirportName string `json:"airport_name"`
	Area        string `json:"area"`
	ICAO        string `json:"cccc"`
	NameCN      string `json:"cname"`
	NameEN      string `json:"eng_name"`
	Country     string `json:"county"`
	DetailCN    string `json:"detail"`
	DetailEN    string `json:"detail_en"`
	Lat         string `json:"lat"`
	Lon         string `json:"lng"`
	Temp        string `json:"temp"`
	Weather     string `json:"weather"`
}
