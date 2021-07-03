package main

import (
	"fmt"
	"github.com/Deny7676yar/DZ-GO1/Lesson9/configreader"
)

func main() {
	config, _ := configreader.ParsConfig("conf.json")
	fmt.Println(config)
	configYaml := configreader.ParsConfigYaml("conf.yaml")
	fmt.Println(configYaml)

	fmt.Println("Port valid URL:", configreader.IsUrl(config.Port))
	fmt.Println("DB_url valid URL:", configreader.IsUrl(config.DB_url))
	fmt.Println("JaegerUrl valid URL:", configreader.IsUrl(config.JaegerUrl))
	fmt.Println("SentryUrl valid URL:", configreader.IsUrl(config.SentryUrl))
	fmt.Println("Some_app_id valid URL:", configreader.IsUrl(config.Some_app_id))
	fmt.Println("Some_app_key valid URL:", configreader.IsUrl(config.Some_app_key))
}




