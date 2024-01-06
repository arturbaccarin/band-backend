package auth

import "github.com/spf13/viper"

type AuthConfig struct {
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
}

func LoadConfig() *AuthConfig {
	var cfg AuthConfig

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
