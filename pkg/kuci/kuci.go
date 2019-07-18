package kuci

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Controller struct
type Controller struct {
}

// NewController function
func NewController() *Controller {
	c := &Controller{}
	return c
}

func (c *Controller) Start() {
	for {
		start := time.Now()
	}
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kuci")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
