package cfg

import "os"

func GetKafkaAddress() string {
	cs := os.Getenv("KAFKA")
	if cs == "" {
		panic("kafka connection not specified")
	}
	return cs
}

func GetDutyerHost() string {
	h := os.Getenv("SERVICE_DUTYER")
	if h == "" {
		panic("client teams service address not specified")
	}
	return h
}

func GetSlackToken() string {
	cs := os.Getenv("SLACK")
	if cs == "" {
		panic("slack token not specified")
	}
	return cs
}
