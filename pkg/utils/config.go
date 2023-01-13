package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

// Single Point in Program to Fetch all the Environment Variables.

func GetBaseUrl() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return fmt.Sprint(viper.Get("BASEURL"))
}

func GetMaxGoRoutines() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return fmt.Sprint(viper.Get("MAX_GOROUTINES"))
}

func GetDatabaseUser() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("POSTGRES_USER")
}

func GetGitHubGQLApi() string {
	viper.SetConfigFile("env")
	viper.ReadInConfig()

	return fmt.Sprint(viper.Get("GITHUB_GRAPHQL_API"))
}

func GetSourceGraphGQLApi() string {
	viper.SetConfigFile("env")
	viper.ReadInConfig()

	return fmt.Sprint(viper.Get("SOURCEGRAPH_GRAPHQL_API"))
}

func GetProcessDirPath() string {
	viper.SetConfigFile("env")
	viper.ReadInConfig()

	return fmt.Sprint(viper.Get("PROCESS_DIR"))
}

func GetDefaultGoPath() string {
	viper.SetConfigFile("env")
	viper.ReadInConfig()

	return fmt.Sprint(viper.Get("GOPATH"))
}

func GetDatabasePassword() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("POSTGRES_PASSWORD")
}

func GetDatabaseName() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("POSTGRES_DB")
}

func GetDatabaseHost() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("POSTGRES_HOST")
}

func GetDatabasePort() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("POSTGRES_PORT")
}

func GetGithubToken() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("GITHUB_TOKEN")
}

func GetGithubUsername() interface{} {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return viper.Get("GITHUB_USERNAME")
}
