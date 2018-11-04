package mgr

import (
	"errors"
	"net/http"
	"regexp"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jchprj/GeoOrderTest/models"
)

var orders sync.Map
var autoID int64

//The expression was searched from Google
var re = regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)

func generateOrderID() int64 {
	atomic.AddInt64(&autoID, 1)
	return autoID
}

func validateLatLong(latitude, longitude string) bool {
	return re.MatchString(latitude + "," + longitude)
}

func calculateDistance(start, end []string) int {
	return 1
}

//NewOrder create new order
func NewOrder(msg models.PlaceRequest) (statusCode int, err error) {
	if len(msg.Origin) != 2 || len(msg.Destination) != 2 {
		return http.StatusBadRequest, errors.New("string length illegal")
	}
	startLatitude, startLongitude := msg.Origin[0], msg.Origin[1]
	endLatitude, endLongitude := msg.Destination[0], msg.Destination[1]
	if validateLatLong(startLatitude, startLongitude) == false || validateLatLong(endLatitude, endLongitude) == false {
		return http.StatusBadRequest, errors.New("latitude or longitude illegal")
	}
	order := models.Order{
		ID:             generateOrderID(),
		Distance:       calculateDistance(msg.Origin, msg.Destination),
		Status:         models.OrderStatusUnassigned,
		StartLatitude:  startLatitude,
		StartLongitude: startLongitude,
		EndLatitude:    endLatitude,
		EndLongitude:   endLongitude,
		CreateTime:     time.Now(),
	}
	orders.Store(order.ID, order)
	return http.StatusOK, nil
}
