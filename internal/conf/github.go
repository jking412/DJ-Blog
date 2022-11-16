package conf

var (
	GithubClientId    string
	GithubSecret      string
	GithubCallbackUrl = "https://github.com/login/oauth/authorize?client_id=%s"
)
