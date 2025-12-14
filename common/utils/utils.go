package utils

import (
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJson(dest any, fileName, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(fileName)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(dest)
	if err != nil {
		logrus.Error("failed to unmarshal %v", err)
		return err
	}

	return nil
}

func SetEnvFromKonsulKV(v *viper.Viper) error {
	env := make(map[string]any)

	err := v.Unmarshal(env)
	if err != nil {
		logrus.Error("failed to unmarshal %v", err)
		return err
	}

	for k, v := range env {
		var (
			valOf = reflect.ValueOf(v)
			val string
		)
		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int:
			val = strconv.Itoa(int(valOf.Int()))
		case reflect.Uint :
			val = strconv.Itoa(int(valOf.Uint()))
		case reflect.Float32:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Float64:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Bool:
			val = valOf.String()
		default:
			panic("unsupported types")
		}
		
		err = os.Setenv(k, val)
		if err != nil {
			logrus.Error("failed to set env %v", err)
			return err
		}
	}

	return nil
}
 

func BindFromKonsulKV(dest any, endPoint, path string) error {
	v := viper.New()

	v.SetConfigType("json")

	err := v.AddRemoteProvider("consul", endPoint, path)
	if err != nil {
		logrus.Error("failed to add remote provider %v", err)
		return err
	}


	err = v.ReadRemoteConfig()
	if err != nil {
		logrus.Error("failed to read remote config %v", err)
		return err
	}

	err = v.Unmarshal(dest)
	if err != nil {
		logrus.Error("failed to unmarshal %v", err)
		return err
	}

	err = SetEnvFromKonsulKV(v)
	if err != nil {
		logrus.Error("failed to set env from konsul kv %v", err)
		return err
	}

	return nil 
}