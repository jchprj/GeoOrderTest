package mgr

import (
	"errors"
	"net/http"
	"time"

	"github.com/jchprj/GeoOrderTest/models"
)

//For short, there is only one sync.Map to store orders,
// if in very huge orders condition, lock and unlock would be a bottleneck, may separate orders to different areas
var orders []models.Order

//NewOrder create new order
func NewOrder(msg models.PlaceRequest) (resp *models.PlaceResponse, statusCode int, err error) {
	lock.Lock()
	defer lock.Unlock()

	if len(msg.Origin) != 2 || len(msg.Destination) != 2 {
		return nil, http.StatusBadRequest, errors.New(models.ErrorDescription)
	}
	startLatitude, startLongitude := msg.Origin[0], msg.Origin[1]
	endLatitude, endLongitude := msg.Destination[0], msg.Destination[1]
	if validateLatLong(startLatitude, startLongitude) == false || validateLatLong(endLatitude, endLongitude) == false {
		return nil, http.StatusBadRequest, errors.New(models.ErrorDescription)
	}
	distance, err := calculateDistance(msg.Origin, msg.Destination)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	order := models.Order{
		ID:             generateOrderID(),
		Distance:       distance,
		Status:         models.OrderStatusUnassigned,
		StartLatitude:  startLatitude,
		StartLongitude: startLongitude,
		EndLatitude:    endLatitude,
		EndLongitude:   endLongitude,
		CreateTime:     time.Now(),
	}
	orders = append(orders, order)
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
	for i := start; i < end && i < len(orders); i++ {
		order := orders[i]
		if order == (models.Order{}) {
			break
		}
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

	var order models.Order
	for _, v := range orders {
		if v.ID == orderID {
			order = v
		}
	}
	if order == (models.Order{}) {
		return http.StatusNotFound, errors.New(models.ErrorOrderNotFound)
	}
	if order.Status == models.OrderStatusTaken {
		return http.StatusConflict, errors.New(models.ErrorOrderAlreadyBeenTaken)
	}
	if isTest == true {
		time.Sleep(time.Second)
	}
	order.Status = models.OrderStatusTaken
	return http.StatusOK, nil
}
