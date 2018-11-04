package mgr

import (
	"testing"
	"time"

	"github.com/jchprj/GeoOrderTest/models"
)

func TestSelect(t *testing.T) {
	TestCreateEngine(t)
	result, err := selectAll()
	if err != nil {
		t.Errorf("err %v", err)
	}
	orderCount := len(orders)
	if orderCount != result {
		t.Errorf("err got %v want %v", orderCount, result)
	}
	t.Logf("len %d", result)
}

func TestInsert(t *testing.T) {
	TestCreateEngine(t)
	order := models.Order{
		Status:     models.OrderStatusUnassigned,
		CreateTime: time.Now(),
	}
	err := insert(&order)
	if order.OrderID == 0 {
		t.Errorf("OrderID == 0")
	}
	if err != nil {
		t.Errorf("err %v", err)
	}
}

func TestUpdate(t *testing.T) {
	TestCreateEngine(t)
	num, err := selectOne(11)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if num == 0 {
		order := models.Order{
			OrderID:    11,
			Status:     models.OrderStatusUnassigned,
			CreateTime: time.Now(),
		}
		err := insert(&order)
		if err != nil {
			t.Errorf("err %v", err)
		}
	}
	order := models.Order{
		OrderID:   11,
		Status:    models.OrderStatusTaken,
		TakenTime: time.Now(),
	}
	err = update(order)
	if err != nil {
		t.Errorf("err %v", err)
	}
}

func TestStartAutoID(t *testing.T) {
	TestSelect(t)
	autoID := GetCurrentAutoID()
	if autoID == 0 {
		t.Error("autoID == 0")
	}
}
