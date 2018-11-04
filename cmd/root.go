package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jchprj/GeoOrderTest/api"
	"github.com/jchprj/GeoOrderTest/cfg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(finishInit)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	rootCmd.PersistentFlags().StringVar(&logDir, "logDir", "logs", "log file directory")
	rootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "error", "log level")
	rootCmd.PersistentFlags().StringVar(&logPrefix, "logPrefix", "", "log file prefix")
	rootCmd.PersistentFlags().BoolVar(&logStd, "logStd", false, "log to console as well")
}
func finishInit() {
	initLog()
	initConfig()
	logrus.Info("init complete")
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("No config file found, exit")
		os.Exit(1)
	}
	logrus.Info("ReadInConfig")
	cfg.HTTPServer = cfg.HTTPServerConfig{
		Addr:            viper.GetString("HTTPServer.Addr"),
		ShutdownTimeout: time.Second * viper.GetDuration("HTTPServer.ShutdownTimeout"),
		ReadTimeout:     time.Second * viper.GetDuration("HTTPServer.ReadTimeout"),
		WriteTimeout:    time.Second * viper.GetDuration("HTTPServer.WriteTimeout"),
		IdleTimeout:     time.Second * viper.GetDuration("HTTPServer.IdleTimeout"),
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "GeoOrderTest config.yml",
	Short: "Based on Google Maps API",
	Long: `Use MySQL store data, also cache data in memory.
	Support 3 API, list/place/take.`,
	Args: cobra.MinimumNArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		api.Init()
	},
}
