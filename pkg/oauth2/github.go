package oauth2

type Token struct {
	AccessToken string
}

type GithubUserInfo struct {
	Username  string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}
