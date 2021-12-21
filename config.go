package main

import "App/routers"

type Config struct {
	routers.Config
}

var config = &Config{
	Config: routers.Config{
		HttpPort:  ":8080",
		HttpsPort: ":8443",
	},
}
