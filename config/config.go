package config

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"path/filepath"
	"smartparking/pkg/tools"
	"time"
)

var (
	GlobalConfig Config
	DBConn       *gorm.DB
	CacheConn    *memcache.Client
)

type Config struct {
	DB         DB         `mapstructure:"db"`
	Cache      Cache      `mapstructure:"cache"`
	Web        Web        `mapstructure:"web"`
	JWT        JWT        `mapstructure:"jwt"`
	Email      Email      `mapstructure:"email"`
	Recognizer Recognizer `mapstructure:"recognizer"`
	ConfigPath string     `mapstructure:"config_path"`
	HashSalt   string     `mapstructure:"hash_salt"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Cache struct {
	Host string        `mapstructure:"host"`
	Port string        `mapstructure:"port"`
	TTL  time.Duration `mapstructure:"ttl"`
}

type Web struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	TLSEnable   bool   `mapstructure:"tlsenable"`
	CertFile    string `mapstructure:"certfile"`
	KeyFile     string `mapstructure:"keyfile"`
	FileStorage string `mapstructure:"file_storage"`
}

type JWT struct {
	SecretKey       string        `mapstructure:"secret"`
	AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
}

type Email struct {
	Sender   string `mapstructure:"sender"`
	Password string `mapstructure:"password"`
	SMTPHost string `mapstructure:"smtp_host"`
	SMTPPort string `mapstructure:"smtp_port"`
}

type Recognizer struct {
	URL   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

func (w Web) String() string {
	return fmt.Sprintf("%s:%s", w.Host, w.Port)
}

// Init - reading config from file
func Init(configPath string) error {
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return err
	}

	return nil
}

// SaveConfig - saves config from struct to file
func SaveConfig(path string) error {
	if err := tools.CreateDir(filepath.Dir(path)); err != nil {
		return err
	}

	data, err := yaml.Marshal(GlobalConfig)
	if err != nil {
		return err
	}

	if err = tools.CreateFileAndWrite(path, data); err != nil {
		return err
	}

	return nil
}
