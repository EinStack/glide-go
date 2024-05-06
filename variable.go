package glide

import (
	"fmt"
	"os"
)

// ClientVersion is the current version of this client.
var ClientVersion = "0.1.0"

// GoVersion is the required version of the Go runtime.
var GoVersion = "1.22.2"

var envApiKey = getEnv("GLIDE_API_KEY", "")
var userAgent = fmt.Sprintf("Glide/%s (Go; Ver. %s)", ClientVersion, GoVersion)
var envUserAgent = getEnv("GLIDE_USER_AGENT", userAgent)
var envBaseUrl = getEnv("GLIDE_BASE_URL", "http://127.0.0.1:9099/")

func getEnv(key, df string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return df
}
