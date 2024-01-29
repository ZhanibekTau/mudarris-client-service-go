package structures

type RestConfig struct {
}

type TokenConfig struct {
	ApiSecret string `mapstructure:"API_SECRET" json:"api_secret"`
}

type RabbitConfig struct {
	RabbitHost     string `mapstructure:"RABBIT_HOST" json:"rabbit_host"`
	RabbitLogin    string `mapstructure:"RABBIT_LOGIN" json:"rabbit_login"`
	RabbitPassword string `mapstructure:"RABBIT_PASSWORD" json:"rabbit_password"`
	RabbitPort     string `mapstructure:"RABBIT_PORT" json:"rabbit_port"`
}
