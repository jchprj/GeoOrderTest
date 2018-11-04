package mgr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"sync/atomic"

	"github.com/Sirupsen/logrus"

	"github.com/jchprj/GeoOrderTest/cfg"
)

var autoID int64
var lock sync.RWMutex
var isTest bool

//Test use test mode
func Test() {
	isTest = true
}

//The expression was searched from Google
var re = regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)

//GetCurrentAutoID for test only
func GetCurrentAutoID() int64 {
	return autoID
}
func generateOrderID() int64 {
	atomic.AddInt64(&autoID, 1)
	return autoID
}

func validateLatLong(latitude, longitude string) bool {
	return re.MatchString(latitude + "," + longitude)
}

//In order for easily test, if Google returns ZERO_RESULTS, return 0 as legal distance
func calculateDistance(start, end []string) (result int, err error) {
	url := fmt.Sprintf(cfg.ThirdParty.GoogleMapsAPIUrl, start[0], start[1], end[0], end[1], cfg.ThirdParty.GoogleMapsAPIKey)
	defer logrus.Info(url, result)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(contents, &obj)
	if err != nil {
		return 0, err
	}
	rows := obj["rows"].([]interface{})
	row := rows[0].(map[string]interface{})
	elements := row["elements"].([]interface{})
	element := elements[0].(map[string]interface{})
	status := element["status"].(string)
	if status == "ZERO_RESULTS" {
		return 0, nil
	}
	if status != "OK" {
		return 0, errors.New(status)
	}
	distance := element["distance"].(map[string]interface{})
	value := distance["value"].(float64)

	return (int)(value), nil
}
