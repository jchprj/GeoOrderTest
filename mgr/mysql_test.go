package mgr

import (
	"fmt"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/jchprj/GeoOrderTest/cfg"

	"github.com/go-sql-driver/mysql"
)

func TestCreateEngine(t *testing.T) {
	cfg.InitConfig("../docker/config.yml")
	err := initMySQL()
	if err != nil {
		t.Errorf("init err %v", err)
	}
	engine, err := GetEngine()
	if err != nil {
		t.Errorf("err %v", err)
	}
	if engine == nil {
		t.Errorf("engine err all nil")
	}
	if err := engine.Ping(); err != nil {
		t.Errorf("engine Ping %v", err)
	}
}

func Test__(t *testing.T) {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "geoordertest",
		AllowNativePasswords: true,
		Params:               map[string]string{"charset": "utf8"},
	}
	dsn := conf.FormatDSN()
	fmt.Println(dsn)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		t.Errorf("err %v", err)
	}
	if engine == nil {
		t.Errorf("engine err all nil")
	}
	if err := engine.Ping(); err != nil {
		t.Errorf("engine Ping %v", err)
	}
}
