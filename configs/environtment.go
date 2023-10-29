package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var (
	environtmentApp string
	configFileApp   string
)

func InitializeConfigAndEnvirontment(fileName string) error {
	splits := strings.Split(filepath.Base(fileName), ".")
	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(fileName))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	configFileApp = fileName
	environtmentApp = GetConfigString("server.mode")
	return nil
}

func checkConfigKey(key string) {
	if !viper.IsSet(key) {
		fmt.Printf("Configuration key %s not found; aborting \n", key)
		os.Exit(1)
	}
}

func GetConfigString(key string) string {
	checkConfigKey(key)
	return viper.GetString(key)
}

func GetConfigInt(key string) int {
	checkConfigKey(key)
	return viper.GetInt(key)
}

func GetConfigBoolean(key string) bool {
	checkConfigKey(key)
	return viper.GetBool(key)
}

func GetConfigInAppModeString(key string) string {
	keyInAppMode := fmt.Sprintf("%s.%s", environtmentApp, key)
	checkConfigKey(keyInAppMode)
	return viper.GetString(keyInAppMode)
}

func GetConfigInAppModeInt(key string) int {
	keyInAppMode := fmt.Sprintf("%s.%s", environtmentApp, key)
	checkConfigKey(keyInAppMode)
	return viper.GetInt(keyInAppMode)
}

func GetConfigInAppModeBoolean(key string) bool {
	keyInAppMode := fmt.Sprintf("%s.%s", environtmentApp, key)
	checkConfigKey(keyInAppMode)
	return viper.GetBool(keyInAppMode)
}

func GetEnvirontmentApp() string {
	return environtmentApp
}
