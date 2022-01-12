package configreader

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

)

type ConfigYaml struct {

	Port         string `yaml:"port"`
	DB_url       string `yaml:"db_url"`
	JaegerUrl    string `yaml:"jaeger_url"`
	SentryUrl    string `yaml:"sentry_url"`
	Kafka_broker string `yaml:"kafka_broker"`
	Some_app_id  string `yaml:"some_app_id"`
	Some_app_key string `yaml:"some_app_key"`

}

func ParsConfigYaml(filename string) *ConfigYaml {
	confYaml := ConfigYaml{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &confYaml)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &confYaml
}