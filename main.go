package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"net/http"
	"strconv"
	"github.com/Dorbmon/MyHouse/config"
)
var sensorType = dht.DHT11
func main(){
	_,_,_, err := dht.ReadDHTxxWithRetry(sensorType, config.DHT11, false, 10)
	// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
	// You may enable "boost GPIO performance" parameter, if your device is old
	// as Raspberry PI 1 (this will require root privileges). You can switch off
	// "boost GPIO performance" parameter for old devices, but it may increase
	// retry attempts. Play with this parameter.
	if err != nil {
		fmt.Println(err);
	}
	http.HandleFunc("/dht",Dht)
	http.ListenAndServe("0.0.0.0",nil)
}
func Dht(w http.ResponseWriter, r *http.Request){
	//读取数据
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(sensorType, config.DHT11, false, 10)
	var Data []byte
	if err != nil{
		Data = []byte(`
	<div class="container">
   <div class="jumbotron">
        <h1>来自DHT11传感器的讯息</h1>
        传感器错误。请检查引脚是否插在:a
   </div>
</div>
`)
	}
	Data = []byte(`
	<div class="container">
   <div class="jumbotron">
        <h1>来自DHT11传感器的讯息</h1>
        <p>温度:<h1>` + FloatToString(temperature) + `</h1></n>
		湿度:<h1>` + FloatToString(humidity) +`%</h1></p></n>
		尝试次数:<1>`+ strconv.Itoa(retried) +`
   </div>
</div>
`)
	w.Write(Data)
	return
}
func FloatToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}