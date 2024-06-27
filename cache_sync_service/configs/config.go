package configs

import "github.com/spf13/viper"

type conf struct {
	DBDriver  string `mapstructure:"DB_DRIVER"`
	DBUri     string `mapstructure:"DB_URI"`
	CacheUri  string `mapstructure:"CACHE_URI"`
	CachePass string `mapstructure:"CACHE_PASS"`
}

func LoadConfig(paths []string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
