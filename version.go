package glide

import (
	"fmt"
	"os"
)

// Version is a supported API version.
var Version string = "0.1.0"

var envApiKey string = getEnv("GLIDE_API_KEY", "development")
var envUserAgent string = getEnv("GLIDE_USER_AGENT", fmt.Sprintf("glide-go/%s", Version))
var envBaseUrl string = getEnv("GLIDE_BASE_URL", "https://api.einstack.com")

func getEnv(key, df string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return df
}
