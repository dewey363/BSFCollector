package BSFServer

import (
	"InfluxDBDriver"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BSFServer struct {
	InfluxDriver *InfluxDBDriver.InfluxDBDriver
}

func (bs *BSFServer) Start()  {
	http.HandleFunc("/bsf/v1/human_body",bs.handleHumanBody)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func (bs *BSFServer) handleHumanBody(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	println(request.Method)
	if request.Method == "POST" {
		println("ok")
		result, _ := ioutil.ReadAll(request.Body)
		request.Body.Close()
		fmt.Printf("%s\n", result)
		var f map[string]string
		json.Unmarshal([]byte(result), &f)
		println(f["weight"])
		bs.InfluxDriver.Write("bodyhealth,name=test weight="+f["weight"])
		writer.WriteHeader(200)
	} else if request.Method == "OPTIONS" {
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(400)
		writer.Write([]byte("not suport method"))
	}
}
