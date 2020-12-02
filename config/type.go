package config

type Config struct {
	B2 B2Config
}

type B2Config struct {
	KeyID          string
	KeyName        string
	ApplicationKey string
	BucketName     string
	Endpoint       string
	Region         string
}
