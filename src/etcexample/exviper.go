package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println(viper.Get("server.url"))
	fmt.Println(viper.Get("server.port"))
	fmt.Println(viper.Get("db.url"))
	fmt.Println(viper.Get("db.port"))

}
