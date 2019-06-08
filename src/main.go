package main

import (
	"BSFServer"
	"InfluxDBDriver"
	"net/http"
)

func main()  {
	println("begin")
	server:=&BSFServer.BSFServer{
		InfluxDriver: &InfluxDBDriver.InfluxDBDriver{
					DBUrl:  "http://192.168.0.110:9999",
					ORG: "dsf",
					Bucket: "test2",
					Client: &http.Client{},
					Token: "Token IoqiJeKK831LcyFUVz_ul0zYrLhHDcW7rBCMxu0pwrkQg773CqydapZKFtaAdF7pLnbYNxXxATg8kD38-B7yQA==",
				},

	}
	server.Start()
}

