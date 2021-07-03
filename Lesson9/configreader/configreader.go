package configreader

import (
	"encoding/json"
	"os"
)

type Consigment struct {
	Port         string `json:"port"`
	DB_url       string `json:"db_url"`
	JaegerUrl    string `json:"jaeger_url"`
	SentryUrl    string `json:"sentry_url"`
	Kafka_broker struct{
		Kafka string `json:"kafka"`
	}
	Some_app_id  string `json:"some_app_id"`
	Some_app_key string `json:"some_app_key"`
}

func ParsConfig(filename string) (Consigment, error) {
	var config Consigment
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}


