package config

import (
	"fmt"
	"os"
	"strconv"
)

type JwtConfig struct {
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}

func GetJwtConfig() JwtConfig {
	p1, err := strconv.Atoi(os.Getenv("AccessTokenExpiryHour"))
	if err != nil {
		panic(fmt.Sprintf("failed to parse jwt config: %v", err))
	}
	p2, err := strconv.Atoi(os.Getenv("RefreshTokenExpiryHour"))
	if err != nil {
		panic(fmt.Sprintf("failed to parse jwt config: %v", err))
	}

	return JwtConfig{
		AccessTokenExpiryHour: p1,
		RefreshTokenExpiryHour: p2,
		AccessTokenSecret: os.Getenv("AccessTokenSecret"),
		RefreshTokenSecret: os.Getenv("RefreshTokenSecret"),
	}
}