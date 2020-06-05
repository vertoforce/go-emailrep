package emailrep

// Client is an emailrep client that stores an api key
type Client struct {
	// It is not required to use an API key but emailrep.io will ask for one if you make too many requests
	APIKey string
	// Optional user agent.  Emailrepo.io asks for each app to set their user agent to the name of your app
	// If you leave it blank, the default go user agent will be used
	UserAgent string
}
