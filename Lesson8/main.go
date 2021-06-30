package main

import (
	"flag"
	"fmt"
	"github.com/Deny7676yar/DZ-GO1/Lesson8/configreader"
	"net/url"
)



/**
перебор конфигурации через flag
 */
func main() {

	var port = flag.String("port","8080", "PortNumb")
	var db_url = flag.String("db_url","postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "DB_url")
	var jaeger_url = flag.String("jaeger", "http://jaeger:16686", "url")
	var sentry_url = flag.String("sentry", "http://sentry:9000", "url")
	var kafka_broker = flag.String("kafka", "kafka:9092", "broker")
	var some_app_id = flag.String("some", "testid", "test1")
	var some_app_key = flag.String("key", "testkey", "test2")
	flag.Parse()

	config := configreader.Config{
		Port: *port,
		DB_url: *db_url,
		JaegerUrl: *jaeger_url,
		SentryUrl: *sentry_url,
		Kafka_broker: *kafka_broker,
		Some_app_id: *some_app_id,
		Some_app_key: *some_app_key,
	}


	fmt.Println(config.Port)
	fmt.Println(IsUrl(*port), "\n")
	fmt.Println(config.DB_url)
	fmt.Println(IsUrl(*db_url), "\n")
	fmt.Println(config.JaegerUrl)
	fmt.Println(IsUrl(*jaeger_url), "\n")
	fmt.Println(config.SentryUrl)
	fmt.Println(IsUrl(*sentry_url), "\n")
	fmt.Println(config.Kafka_broker)
	fmt.Println(IsUrl(*kafka_broker),"\n")
	fmt.Println(config.Some_app_id)
	fmt.Println(IsUrl(*some_app_id), "\n")
	fmt.Println(config.Some_app_key)
	fmt.Println(IsUrl(*some_app_id), "\n")

}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}