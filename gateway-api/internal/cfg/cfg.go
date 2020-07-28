package cfg

import "os"

func GetPeopleHost() string {
	h := os.Getenv("PEOPLE_CLIENT")
	if h == "" {
		h = ":50051"
		//panic("client teams service address not specified")
	}

	return h
}

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
