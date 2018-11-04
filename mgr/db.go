package mgr

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jchprj/GeoOrderTest/models"
)

func selectAll() (int, error) {
	engine, _ := GetEngine()
	sql := fmt.Sprintf("select * from `order` where Status='%v' or Status='%v' order by order_i_d", models.OrderStatusUnassigned, models.OrderStatusTaken)
	result, err := engine.QueryString(sql)
	if err != nil {
		return 0, err
	}

	for _, obj := range result {
		order := models.Order{}
		for k, v := range obj {
			switch k {
			case "order_i_d":
				order.OrderID, _ = strconv.ParseInt(v, 10, 64)
			case "distance":
				order.Distance, _ = strconv.Atoi(v)
			case "status":
				order.Status = v
			case "start_latitude":
				order.StartLatitude = v
			case "start_longitude":
				order.StartLongitude = v
			case "end_latitude":
				order.EndLatitude = v
			case "end_longitude":
				order.EndLongitude = v
			case "createTime":
				order.CreateTime, _ = time.Parse(models.SQLTimeFormat, v)
			case "takenTime":
				order.TakenTime, _ = time.Parse(models.SQLTimeFormat, v)
			}
		}
		orders = append(orders, &order)
	}
	return len(result), nil
}

func selectOne(orderID int64) (int, error) {
	engine, _ := GetEngine()
	sql := fmt.Sprintf("select * from `order` where order_i_d=%v", orderID)
	result, err := engine.QueryString(sql)
	if err != nil {
		return 0, err
	}
	return len(result), nil
}

func insert(order *models.Order) error {
	engine, _ := GetEngine()
	num, err := engine.Insert(order)
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("0 inserted")
	}
	return nil
}

func update(order models.Order) error {
	engine, _ := GetEngine()
	num, err := engine.Update(order, &models.Order{OrderID: order.OrderID})
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("0 updated")
	}
	return nil
}
