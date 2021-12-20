package routers

type Config struct {
	HttpPort     string
	HttpsPort    string
	CertDirCache string
}

func DefaultConfig() *Config {
	return &Config{
		HttpPort:     ":8080",
		HttpsPort:    ":8081",
		CertDirCache: ".cache",
	}
}
