package model

type OAuthTokens struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessKey      string
	AccessSecret   string
}
type Tweet struct {
	Message string
	Token   OAuthTokens
}
