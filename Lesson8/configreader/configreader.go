package configreader

type Config struct {
	Port         string
	DB_url       string
	JaegerUrl    string
	SentryUrl    string
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func NewConfig() *Config {
	return &Config{}
}

