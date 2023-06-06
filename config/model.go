package config

type LocalConfig struct {
	Name          string `mapstructure:"APP_NAME"`
	Port          string `mapstructure:"PORT"`
	ChannelSecret string `mapstructure:"LINE_CHANNEL_SECRET"`
	ChannelToken  string `mapstructure:"LINE_CHANNEL_TOKEN"`
	ChannelID     string `mapstructure:"LINE_CHANNEL_ID"`
	DBHost        string `mapstructure:"POSTGRES_HOST"`
	DBPort        string `mapstructure:"POSTGRES_PORT"`
	DBName        string `mapstructure:"POSTGRES_NAME"`
	DBUser        string `mapstructure:"POSTGRES_USER"`
	DBPassword    string `mapstructure:"POSTGRES_PASSWORD"`
}
