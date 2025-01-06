package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// TODO: create a function to validate env variables, add a required flag, and then a check to see if its required but
// not given print a warning to the console: USING DEFAULT VALUE FOR ENV VAR: <var>. Why? To continue to enable easy
// local development but also to make deployment error handling a better ux.

// NOTE: this is the localhost url used within the docker compose file when spinning up the minio instance locally
const LocalBucketURL = "http://localhost:9000"

var Envs = loadConfig()

type Config struct {
	DBType              string
	AppPort             string
	AwsAccessKeyId      string
	AwsAccountId        string
	StorageBackendType  string
	DBConnString        string
	AwsSecretAccessKey  string
	DBName              string
	AwsS3BucketName     string
	AwsS3Region         string
	CloudflareAccountId string
	GithubClientSecret  string
	GithubClientId      string
	GoogleClientId      string
	GoogleClientSecret  string
	AuthSecret          string
	JWTExpiration       time.Duration
	IsProd              bool
}

// Loads all environment variables from the .env file
func loadConfig() Config {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
	return Config{
		AppPort:             getEnv("PORT", "3000"),
		DBConnString:        getEnv("CONN_STRING", "mongodb://root:example@localhost:27017"),
		DBName:              getEnv("DB_NAME", "test"),
		DBType:              strings.ToLower(getEnv("DB_TYPE", "mongo")),
		StorageBackendType:  strings.ToLower(getEnv("STORAGE_BACKEND_TYPE", "s3")),
		AwsS3BucketName:     getEnv("AWS_S3_BUCKET_NAME", "testBucket"),
		AwsS3Region:         getEnv("AWS_S3_REGION", "ap-southeast-2"),
		AwsAccessKeyId:      getEnv("AWS_ACCESS_KEY_ID", ""),
		AwsSecretAccessKey:  getEnv("AWS_SECRET_ACCESS_KEY", ""),
		AwsAccountId:        getEnv("AWS_ACCOUNT_ID", ""),
		JWTExpiration:       getEnvAsDuration("JWT_EXPIRATION_IN_SECONDS", 3600),
		CloudflareAccountId: getEnv("CLOUDFLARE_ACCOUNT_ID", ""),
		IsProd:              getEnvAsBoolean("IS_PROD", false),
		AuthSecret:          getEnv("AUTH_SECRET", "superSecretNeedsToBeChanged"),
		GithubClientId:      getEnv("GITHUB_CLIENT_ID", ""),
		GithubClientSecret:  getEnv("GITHUB_CLIENT_SECRET", ""),
		GoogleClientId:      getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret:  getEnv("GOOGLE_CLIENT_SECRET", ""),
	}
}

// Gets an environment variable and provides a default value if not set
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Gets an environment variable as an int64 and provides a default value if not set
func getEnvAsInt64(key string, fallback int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return fallback
}

// Modified function to return time.Duration
func getEnvAsDuration(key string, fallback int64) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return time.Duration(intValue) * time.Second
		}
	}
	return time.Duration(fallback) * time.Second
}

// Gets an environment variable as a boolean from a string, provides the fallback value if not set
func getEnvAsBoolean(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}
