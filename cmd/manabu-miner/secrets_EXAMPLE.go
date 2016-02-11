package main

//APIKeysEXAMPLE contain all the API keys from the Twitter Dev account
type APIKeysEXAMPLE struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// SecretsEXAMPLE contain API keys, passwords, hash whatevers
var SecretsEXAMPLE = APIKeysEXAMPLE{
	ConsumerKey:       "EXAMPLE",
	ConsumerSecret:    "EXAMPLE",
	AccessToken:       "EXAMPLE",
	AccessTokenSecret: "EXAMPLE",
}
