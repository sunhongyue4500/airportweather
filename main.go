package main

import (
	_ "airportweather/config"
	"fmt"
	"time"
)

func init() {

}

func main() {
	// 5min获取一次气象数据
	ticker := time.NewTicker(time.Second * 4 * 1)
	defer ticker.Stop()
	for t := range ticker.C {
		go func() {
			MakeRequest()
		}()
		fmt.Println("Current time: ", t)
	}
}
