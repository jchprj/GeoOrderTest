package mgr

import (
	"github.com/jchprj/GeoOrderTest/cfg"

	"github.com/Sirupsen/logrus"

	"github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func initMySQL() error {
	conf := mysql.Config{
		User:                 cfg.Database.User,
		Passwd:               cfg.Database.Passwd,
		Net:                  "tcp",
		Addr:                 cfg.Database.Addr,
		DBName:               cfg.Database.DBName,
		AllowNativePasswords: true,
		Params:               map[string]string{"charset": "utf8"},
	}
	logrus.WithField("configs", conf).Infoln("mysql config")
	return createEngine(&conf)
}

// GetEngine Get *xorm.Engine
func GetEngine() (*xorm.Engine, error) {
	return engine, nil
}

func createEngine(conf *mysql.Config) error {
	dsn := conf.FormatDSN()
	tmpEngine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return err
	}
	if err := tmpEngine.Ping(); err != nil {
		logrus.WithField("dns", dsn).Errorln("DB Connect failed")
		return err
	}
	engine = tmpEngine
	return nil
}
