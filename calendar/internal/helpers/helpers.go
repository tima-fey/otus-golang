package helpers

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type Config struct {
	Address  string `json:"address"`
	WebPort  int    `json:"web_port"`
	GrpcPort int    `json:"grpc_port"`
	LogFile  string `json:"log_file"`
	LogLevel string `json:"log_level"`
}

func checkLogLvl(logLvl string) {
	switch logLvl {
	case
		"info",
		"warn",
		"debug",
		"error":
		return
	}
	log.Fatal("Error: plese specify correct log lvl (info, warn, debug, error)")
	return
}
func checkPort(port int) {
	if port > 65536 {
		log.Fatal("Error port mast be less than 65536")
	}
	if port < 1 {
		log.Fatal("Error port mast be greater than 1")
	}
	return
}

func GetConfig() Config {
	var config string
	flag.StringVar(&config, "config", "", "soource file")
	flag.Parse()
	file, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}
	data := Config{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal(err)
	}
	checkLogLvl(data.LogLevel)
	checkPort(data.WebPort)
	checkPort(data.GrpcPort)

	return data
}
