#!/bin/bash

cd /Users/hongyi/Code/Practice/Go/src/airportweather
go install

cd ~/Code/Practice/Go/bin
#如果文件夹不存在，创建文件夹
if [ ! -d "~/Code/Practice/Go/bin/airportweather" ]; then
  mkdir -p ~/Code/Practice/Go/bin/airport_weather
fi

cp ~/Code/Practice/Go/bin/airportweather ~/Code/Practice/Go/bin/airport_weather
cp ./config.json ~/Code/Practice/Go/bin/airport_weather/config.json