package config

import (
	"encoding/json"
	"log"
)

var cfg *Config

func Init() {
	err := json.Unmarshal([]byte(CONFIG), &cfg)
	if err != nil {
		log.Fatalln("[fatal][config] - failed to load config", err)
	}
	return
}

func GetConfig() *Config {
	return cfg
}

//replace values in quotes/angle brackets with your actual bucket and appkey information
const CONFIG = `
{
	"B2": {
		"KeyID": "<replace with your appkey ID like '002d936312349320001234034'>",
		"KeyName": "<replace with your appkey name like 'golang-codeexample-key'>",
		"ApplicationKey": "<replace with your appkey like 'K0025D1234erWKCDjgEy712345kEfNk'>",
		"BucketName": "<replace with your bucket name like 'golang-codeexample'>",
		"Endpoint": "<replace with your server endpoint like 'https://s3.us-west-002.backblazeb2.com'>",
		"Region": "<replace with your server region like 'us-west-002'>"
	}
}`
