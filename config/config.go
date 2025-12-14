package config

import (
	"os"
	"user-service/common/utils"

	"github.com/sirupsen/logrus"
)

var config *AppConfig

type AppConfig struct {
	Port         int    `json:"port"`
	AppName      string `json:"appName"`
	AppEnv       string `json:"appEnv"`
	SignatureKey string `json:"signatureKey"`
	Database     Database `json:"database"`
	RateLimiterMaxRequest int    `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int    `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string `json:"jwtSecretKey"`
	JwtExpirationTime     int    `json:"jwtExpirationTime"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func init () {
	err := utils.BindFromJson(&config, "config.json", ".")

	if err != nil {
		logrus.Infof("Failed to load config %v", err)
		err = utils.BindFromKonsulKV(&config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY") )
		if err != nil {
			panic(err)
		}
	}
}