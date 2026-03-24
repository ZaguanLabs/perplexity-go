package perplexity

import "runtime"

// Version is the current version of the SDK.
const Version = "1.1.0"

// UserAgent returns the default User-Agent string for the SDK.
func UserAgent() string {
	return "perplexity-go/" + Version
}

func defaultPlatformHeaders() map[string]string {
	return map[string]string{
		"X-Stainless-Lang":            "go",
		"X-Stainless-Package-Version": Version,
		"X-Stainless-OS":              runtime.GOOS,
		"X-Stainless-Arch":            runtime.GOARCH,
		"X-Stainless-Runtime":         "go",
		"X-Stainless-Runtime-Version": runtime.Version(),
	}
}
