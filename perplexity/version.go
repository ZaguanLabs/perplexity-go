package perplexity

// Version is the current version of the SDK.
const Version = "1.0.0"

// UserAgent returns the default User-Agent string for the SDK.
func UserAgent() string {
	return "perplexity-go/" + Version
}
