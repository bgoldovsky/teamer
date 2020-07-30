package cfg

import "os"

func GetKafkaAddress() string {
	cs := os.Getenv("KAFKA")
	if cs == "" {
		panic("connection string not specified")
	}

	return cs
}
