package config

import (
	"encoding/json"
	"os"
)

func write(cfg Config) error {

	write_json, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err = os.WriteFile(cfg.filename, write_json, 0666); err != nil {
		return err
	}

	return nil
}
