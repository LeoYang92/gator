package config

import (
	"fmt"
	"os"
)

func Read() (Config, error) {

	user_home_dir, err := os.UserHomeDir()

	if err != nil {
		return Config{}, err
	}

	config_filename := user_home_dir + "/" + configFileName

	config := Config{
		DBURL:           "postgres://example",
		CurrentUserName: "",
		filename:        config_filename,
	}

	// 如果配置文件不存在则创建
	if _, err := os.Stat(config_filename); os.IsNotExist(err) {
		file, err := os.Create(config_filename)
		if err != nil {
			return Config{}, fmt.Errorf("Create config file fail: %w", err)
		}
		write(config)
		defer file.Close()
	}

	write(config)

	return config, nil
}
