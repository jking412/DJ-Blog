package conf

import "DJ-Blog/pkg/config"

func Init() {
	GithubClientId = config.LoadString("github.clientId")
	GithubSecret = config.LoadString("github.clientSecret")
}
