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
	DB         DB         `mapstructure:"db" yaml:"db"`
	Cache      Cache      `mapstructure:"cache" yaml:"cache"`
	Web        Web        `mapstructure:"web" yaml:"web"`
	JWT        JWT        `mapstructure:"jwt" yaml:"jwt"`
	Email      Email      `mapstructure:"email" yaml:"email"`
	Recognizer Recognizer `mapstructure:"recognizer" yaml:"recognizer"`
	ConfigPath string     `mapstructure:"config_path" yaml:"config_path"`
	HashSalt   string     `mapstructure:"hash_salt" yaml:"hash_salt"`
}

type DB struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	Name     string `mapstructure:"name" yaml:"name"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	SSLMode  string `mapstructure:"sslmode" yaml:"sslmode"`
}

type Cache struct {
	Host string        `mapstructure:"host" yaml:"host"`
	Port string        `mapstructure:"port" yaml:"port"`
	TTL  time.Duration `mapstructure:"ttl" yaml:"ttl"`
}

type Web struct {
	Host        string `mapstructure:"host" yaml:"host"`
	Port        string `mapstructure:"port" yaml:"port"`
	TLSEnable   bool   `mapstructure:"tlsenable" yaml:"tlsenable"`
	CertFile    string `mapstructure:"certfile" yaml:"certfile"`
	KeyFile     string `mapstructure:"keyfile" yaml:"keyfile"`
	FileStorage string `mapstructure:"file_storage" yaml:"file_storage"`
}

type JWT struct {
	SecretKey       string        `mapstructure:"secret" yaml:"secret"`
	AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl" yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl" yaml:"refresh_token_ttl"`
}

type Email struct {
	Sender   string `mapstructure:"sender" yaml:"sender"`
	Password string `mapstructure:"password" yaml:"password"`
	SMTPHost string `mapstructure:"smtp_host" yaml:"smtp_host"`
	SMTPPort string `mapstructure:"smtp_port" yaml:"smtp_port"`
}

type Recognizer struct {
	URL   string `mapstructure:"url" yaml:"url"`
	Token string `mapstructure:"token" yaml:"token"`
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
