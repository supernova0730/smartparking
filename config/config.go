package config

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
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
	DB         DB         `yaml:"db"`
	Cache      Cache      `yaml:"cache"`
	Web        Web        `yaml:"web"`
	JWT        JWT        `yaml:"jwt"`
	Email      Email      `yaml:"email"`
	Recognizer Recognizer `yaml:"recognizer"`
	ConfigPath string     `yaml:"config_path"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

type Cache struct {
	Host string        `yaml:"host"`
	Port string        `yaml:"port"`
	TTL  time.Duration `yaml:"ttl"`
}

type Web struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	TLSEnable   bool   `yaml:"tlsenable"`
	CertFile    string `yaml:"certfile"`
	KeyFile     string `yaml:"keyfile"`
	FileStorage string `yaml:"file_storage"`
}

type JWT struct {
	SecretKey       string        `yaml:"secret"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
}

type Email struct {
	Sender   string `yaml:"sender"`
	Password string `yaml:"password"`
	SMTPHost string `yaml:"smtp_host"`
	SMTPPort string `yaml:"smtp_port"`
}

type Recognizer struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

func (w Web) String() string {
	return fmt.Sprintf("%s:%s", w.Host, w.Port)
}

// Init - reading config from file
func Init(configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &GlobalConfig)
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
