package emailrep

// Client is an emailrep client that stores an api key
// It is not required to use an API key but emailrep.io will ask for one if you make too many requests
type Client struct {
	APIKey    string
	UserAgent string
}
