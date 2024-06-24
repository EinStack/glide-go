package glide

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ClientVersion is the current version of this client.
var clientVersion = "0.1.0"

// GoVersion is the required version of the Go runtime.
var goVersion = getVersion()

func getVersion() string {
	version := runtime.Version()
	after, _ := strings.CutPrefix(version, "go")
	return after
}

// userAgent is a default User-Agent header value.
var userAgent = fmt.Sprintf("Glide/%s (Go; Ver. %s)", clientVersion, goVersion)

var envApiKey = getEnv("GLIDE_API_KEY", "")
var envUserAgent = getEnv("GLIDE_USER_AGENT", userAgent)
var envBaseUrl = getEnv("GLIDE_BASE_URL", "http://127.0.0.1:9099/")

func getEnv(key, df string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return df
}
