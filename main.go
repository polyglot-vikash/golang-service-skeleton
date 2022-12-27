package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sample-app/router"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func main() {
	// Read YAML file
	// Todo: Move to viper based config
	var c Config
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatal(err)
	}

	// Use environment variables
	fmt.Println("Database host:", c.Database.Host)
	fmt.Println("Database port:", c.Database.Port)
	fmt.Println("Database user:", c.Database.User)
	fmt.Println("Database password:", c.Database.Password)
	fmt.Println("Database name:", c.Database.Name)
	fmt.Println("Redis host:", c.Redis.Host)
	fmt.Println("Redis port:", c.Redis.Port)
	fmt.Println("Redis password:", c.Redis.Password)
	fmt.Println("Server port:", c.Server.Port)

	router.SetupRouter()

}
