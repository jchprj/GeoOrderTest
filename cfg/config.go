package cfg

import "time"

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
