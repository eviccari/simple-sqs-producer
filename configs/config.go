package configs

import (
	"fmt"
	"github.com/eviccari/simple-sqs-producer/adapters"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

const (
	DefaultAppName      = "simple-sqs-producer"
	DefaultHTTPPort     = 8080
	DefaultLoggingLevel = "INFO"
	ConfigPath          = "./cmd/.env"
	ContainerConfigPath = ".env"
	ConfigType          = "json"
	ConfigFile          = "configs"
)

var (
	appName      string
	httpPort     int
	loggingLevel string
	awsEndpoint  string
)

// GetAppName - Get application string name, to enable log customization
func GetAppName() string {
	return appName
}

// GetHTTPPort - Get HTTP application port
func GetHTTPPort() int {
	return httpPort
}

// GetLoggingLevel - Get application logging level
func GetLoggingLevel() string {
	return loggingLevel
}

// GetAWSEndpoint - Get AWS Endpoint
func GetAWSEndpoint() string {
	return awsEndpoint
}

func init() {
	viper.AddConfigPath(ConfigPath)
	viper.AddConfigPath(ContainerConfigPath)
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("ERROR ON LOAD CONFIGURATION FILE: %s/%s.%s ---> %s", ConfigPath, ConfigFile, ConfigType, err)
		panic(msg)
	}

	awsEndpoint = viper.GetString("awsEndpoint")

	appName = viper.GetString("appName")
	if appName == "" {
		appName = DefaultAppName
	}

	httpPort = viper.GetInt("httpPort")
	if httpPort == 0 {
		httpPort = DefaultHTTPPort
	}

	loggingLevel = strings.ToUpper(viper.GetString("loggingLevel"))
	if loggingLevel == "" {
		loggingLevel = DefaultLoggingLevel
	}
}

// PrintConfig - Print application configuration file upon console
func PrintConfig(logger adapters.LoggerAdapter) {
	logger.Info(fmt.Sprintf("Starting with application name.....: %s", appName))
	logger.Info(fmt.Sprintf("Starting with app aws endpoint.....: %s", awsEndpoint))
	logger.Info(fmt.Sprintf("Starting with HTTP port............: %s", strconv.Itoa(httpPort)))
}
