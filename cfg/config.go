package cfg

import (
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

//HTTPServerConfig httpServer config
type HTTPServerConfig struct {
	Addr            string
	ShutdownTimeout time.Duration
	WriteTimeout    time.Duration
	ReadTimeout     time.Duration
	IdleTimeout     time.Duration
}

//HTTPServer httpServer config
var HTTPServer HTTPServerConfig

//DatabaseConfig database config
type DatabaseConfig struct {
	User   string
	Passwd string
	Addr   string
	DBName string
}

//Database database config
var Database DatabaseConfig

//ThirdPartyConfig current only Google Maps API
type ThirdPartyConfig struct {
	GoogleMapsAPIKey string
	GoogleMapsAPIUrl string
}

//ThirdParty ThirdParty config
var ThirdParty ThirdPartyConfig

//InitConfig InitConfig
func InitConfig(file string) {
	viper.SetConfigFile(file)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("No config file found, exit")
		os.Exit(1)
	}
	logrus.Info("ReadInConfig")
	HTTPServer = HTTPServerConfig{
		Addr:            viper.GetString("HTTPServer.Addr"),
		ShutdownTimeout: time.Second * viper.GetDuration("HTTPServer.ShutdownTimeout"),
		ReadTimeout:     time.Second * viper.GetDuration("HTTPServer.ReadTimeout"),
		WriteTimeout:    time.Second * viper.GetDuration("HTTPServer.WriteTimeout"),
		IdleTimeout:     time.Second * viper.GetDuration("HTTPServer.IdleTimeout"),
	}
	Database = DatabaseConfig{
		User:   viper.GetString("Database.User"),
		Passwd: viper.GetString("Database.Passwd"),
		Addr:   viper.GetString("Database.Addr"),
		DBName: viper.GetString("Database.DBName"),
	}
	ThirdParty = ThirdPartyConfig{
		GoogleMapsAPIKey: viper.GetString("ThirdParty.GoogleMapsAPIKey"),
		GoogleMapsAPIUrl: viper.GetString("ThirdParty.GoogleMapsAPIUrl"),
	}
}
