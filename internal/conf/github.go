package conf

import "DJ-Blog/pkg/config"

var (
	GithubClientId    = config.LoadString("github.clientId")
	GithubSecret      = config.LoadString("github.clientSecret")
	GithubCallbackUrl = "https://github.com/login/oauth/authorize?client_id=%s"
)
