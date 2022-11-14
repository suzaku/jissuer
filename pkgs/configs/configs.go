package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/pkg/errors"
)

var (
	configFolder = configdir.LocalConfig("jissuer")
	configFile   = filepath.Join(configFolder, "configs.json")

	cfgNotExist = errors.New("Config file does not exist")
)

type Config struct {
	Email   string
	Token   string
	BaseURL string
}

func Load() (*Config, error) {
	return loadFile(configFile)
}

func loadFile(path string) (*Config, error) {
	fh, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, cfgNotExist
	}
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to load config: %s", path)
	}
	defer fh.Close()
	cfg := new(Config)
	decoder := json.NewDecoder(fh)
	decoder.Decode(cfg)
	return cfg, nil
}

func (cfg *Config) saveFile(path string) error {
	log.Printf("Saving configs to %s", path)
	if err := ensureFolderExist(path); err != nil {
		return err
	}
	fh, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fh.Close()
	encoder := json.NewEncoder(fh)
	return encoder.Encode(cfg)
}

func (cfg *Config) Save() error {
	return cfg.saveFile(configFile)
}

func ensureFolderExist(path string) error {
	base := filepath.Dir(path)
	return os.MkdirAll(base, 0766)
}
