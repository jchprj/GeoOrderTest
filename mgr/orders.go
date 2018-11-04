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

//For short, there is only one sync.Map to store orders,
// if in very huge orders condition, lock and unlock would be a bottleneck, may separate orders to different areas
var orders sync.Map
var autoID int64
var lock sync.RWMutex
var isTest bool

//Test use test mode
func Test() {
	isTest = true
}

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
func NewOrder(msg models.PlaceRequest) (resp *models.PlaceResponse, statusCode int, err error) {
	if len(msg.Origin) != 2 || len(msg.Destination) != 2 {
		return nil, http.StatusBadRequest, errors.New(models.ErrorDescription)
	}
	startLatitude, startLongitude := msg.Origin[0], msg.Origin[1]
	endLatitude, endLongitude := msg.Destination[0], msg.Destination[1]
	if validateLatLong(startLatitude, startLongitude) == false || validateLatLong(endLatitude, endLongitude) == false {
		return nil, http.StatusBadRequest, errors.New(models.ErrorDescription)
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
	resp = &models.PlaceResponse{
		ID:       order.ID,
		Distance: order.Distance,
		Status:   order.Status,
	}
	return resp, http.StatusOK, nil
}

//GetOrderList get orders between start and end if exist
func GetOrderList(page, limit int) (list []models.PlaceResponse) {
	start := (page-1)*limit + 1
	end := page * limit
	for i := start; i < end; i++ {
		tmpOrder, ok := orders.Load(int64(i))
		if tmpOrder == nil || ok == false {
			break
		}
		order := tmpOrder.(models.Order)
		orderResponse := models.PlaceResponse{
			ID:       order.ID,
			Distance: order.Distance,
			Status:   order.Status,
		}
		list = append(list, orderResponse)
	}
	return list
}

//TakeOrder change order status to taken
func TakeOrder(orderID int64) (int, error) {
	lock.Lock()
	defer lock.Unlock()
	tmpOrder, ok := orders.Load(orderID)
	if tmpOrder == nil || ok == false {
		return http.StatusNotFound, errors.New(models.ErrorOrderNotFound)
	}
	order := tmpOrder.(models.Order)
	if order.Status == models.OrderStatusTaken {
		return http.StatusConflict, errors.New(models.ErrorOrderAlreadyBeenTaken)
	}
	order.Status = models.OrderStatusTaken
	if isTest == true {
		time.Sleep(time.Second)
	}
	orders.Store(orderID, order)
	return http.StatusOK, nil
}
