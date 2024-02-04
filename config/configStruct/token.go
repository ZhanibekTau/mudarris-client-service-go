package structures

type TokenConfig struct {
	ApiSecretClient string `mapstructure:"API_SECRET_CLIENT" json:"api_secret_client"`
	ApiSecretUstaz  string `mapstructure:"API_SECRET_USTAZ" json:"api_secret_ustaz"`
}
