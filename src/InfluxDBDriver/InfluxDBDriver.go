package InfluxDBDriver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type InfluxDBDriver struct {
	DBUrl  string
	ORG    string
	Bucket string
	Client *http.Client
	Token  string
}

func (driver *InfluxDBDriver) Write(data string) {
	url := fmt.Sprintf("%s/api/v2/write?org=%s&bucket=%s&precision=ns",
		driver.DBUrl, driver.ORG, driver.Bucket)
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {

	}
	req.Header.Set("Authorization", driver.Token)
	req.Header.Set("Content-Type", "text/plain")
	resp, err := driver.Client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
