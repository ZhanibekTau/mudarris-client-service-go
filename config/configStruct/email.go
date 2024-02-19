package structures

type EmailConfig struct {
	EmailUsername string `mapstructure:"EMAIL_USERNAME" json:"email_username"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD" json:"email_password"`
	EmailHost     string `mapstructure:"EMAIL_HOST" json:"email_host"`
	EmailPort     string `mapstructure:"EMAIL_PORT" json:"email_port"`
}
