package glide

import (
	"fmt"
	"os"
)

// ClientVersion is the current version of this client.
var ClientVersion string = "0.1.0"

// GoVersion is the required version of the Go runtime.
var GoVersion string = "0.1.0"

var envApiKey string = getEnv("GLIDE_API_KEY", "development")
var userAgent = fmt.Sprintf("Glide/%s (Go; Ver %s)", ClientVersion, GoVersion)
var envUserAgent string = getEnv("GLIDE_USER_AGENT", userAgent)
var envBaseUrl string = getEnv("GLIDE_BASE_URL", "http://127.0.0.1:9099/")

func getEnv(key, df string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return df
}
