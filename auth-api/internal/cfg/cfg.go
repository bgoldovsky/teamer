package cfg

import "os"

func GetHTTPPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
		//panic("http port not specified")
	}
	return ":" + p
}

func GetSecret() string {
	s := os.Getenv("SECRET")
	if s == "" {
		s = "mySecretSigningString"
		//panic("secret not specified")
	}
	return s
}
