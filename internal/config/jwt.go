package config

type JWT struct {
	AccessTokenExpiryHour  int    `toml:"access_token_expiry_hour"`
	RefreshTokenExpiryHour int    `toml:"refresh_token_expiry_hour"`
	AccessTokenSecret      string `toml:"access_token_secret"`
	RefreshTokenSecret     string `toml:"refresh_token_secret"`
}