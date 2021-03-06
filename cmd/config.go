package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"smartparking/config"
	"smartparking/pkg/logger"
	"smartparking/pkg/tools"
	"strings"
	"time"
)

var ConfigPath = tools.FilepathFromHome(".smartparking/config.yml")

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "initial config generation",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := cmd.Flags().GetString("path")
		if err != nil {
			logger.Log.Fatal("generate config", zap.Error(err))
		}

		if err = viper.Unmarshal(&config.GlobalConfig); err != nil {
			logger.Log.Fatal("generate config", zap.Error(err))
		}

		err = config.SaveConfig(configPath)
		if err != nil {
			logger.Log.Fatal("generate config", zap.Error(err))
		}

		logger.Log.Info("config file generated")
	},
}

func init() {
	viper.SetEnvPrefix("sp")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	confCmd := configCmd.PersistentFlags()

	confCmd.String("path", ConfigPath, "generate config file to (default .smartparking/config.yml)")

	confCmd.StringVar(&config.GlobalConfig.DB.Host, "db_host", "localhost", "database host")
	confCmd.StringVar(&config.GlobalConfig.DB.Port, "db_port", "5432", "database port")
	confCmd.StringVar(&config.GlobalConfig.DB.Name, "db_name", "smartparking", "database name")
	confCmd.StringVar(&config.GlobalConfig.DB.Username, "db_username", "yeldar", "database username")
	confCmd.StringVar(&config.GlobalConfig.DB.Password, "db_password", "000730", "database password")
	confCmd.StringVar(&config.GlobalConfig.DB.SSLMode, "db_sslmode", "disable", "database sslmode")
	_ = viper.BindPFlag("DB.Host", confCmd.Lookup("db_host"))
	_ = viper.BindPFlag("DB.Port", confCmd.Lookup("db_port"))
	_ = viper.BindPFlag("DB.Name", confCmd.Lookup("db_name"))
	_ = viper.BindPFlag("DB.Username", confCmd.Lookup("db_username"))
	_ = viper.BindPFlag("DB.Password", confCmd.Lookup("db_password"))
	_ = viper.BindPFlag("DB.SSLMode", confCmd.Lookup("db_sslmode"))

	confCmd.StringVar(&config.GlobalConfig.Cache.Host, "cache_host", "localhost", "cache host")
	confCmd.StringVar(&config.GlobalConfig.Cache.Port, "cache_port", "11211", "cache port")
	confCmd.DurationVar(&config.GlobalConfig.Cache.TTL, "cache_ttl", 10*time.Minute, "cache host")
	_ = viper.BindPFlag("CACHE.Host", confCmd.Lookup("cache_host"))
	_ = viper.BindPFlag("CACHE.Port", confCmd.Lookup("cache_port"))
	_ = viper.BindPFlag("CACHE.TTL", confCmd.Lookup("cache_ttl"))

	confCmd.StringVar(&config.GlobalConfig.Web.Host, "web_host", "localhost", "web host")
	confCmd.StringVar(&config.GlobalConfig.Web.Port, "web_port", "8080", "web port")
	confCmd.BoolVar(&config.GlobalConfig.Web.TLSEnable, "web_tlsenable", false, "web tls enable")
	confCmd.StringVar(&config.GlobalConfig.Web.CertFile, "web_certfile", "", "web cert file")
	confCmd.StringVar(&config.GlobalConfig.Web.KeyFile, "web_keyfile", "", "web key file")
	confCmd.StringVar(&config.GlobalConfig.Web.FileStorage, "web_file_storage", filepath.Join(tools.HomePath(), ".smartparking/images"), "web file storage")
	_ = viper.BindPFlag("WEB.Host", confCmd.Lookup("web_host"))
	_ = viper.BindPFlag("WEB.Port", confCmd.Lookup("web_port"))
	_ = viper.BindPFlag("WEB.TLSEnable", confCmd.Lookup("web_tlsenable"))
	_ = viper.BindPFlag("WEB.CertFile", confCmd.Lookup("web_certfile"))
	_ = viper.BindPFlag("WEB.KeyFile", confCmd.Lookup("web_keyfile"))
	_ = viper.BindPFlag("WEB.FileStorage", confCmd.Lookup("web_file_storage"))

	confCmd.StringVar(&config.GlobalConfig.JWT.SecretKey, "jwt_secret", "my_secret_key", "jwt secret")
	confCmd.DurationVar(&config.GlobalConfig.JWT.AccessTokenTTL, "jwt_access_token_ttl", 15*time.Minute, "jwt access token ttl")
	confCmd.DurationVar(&config.GlobalConfig.JWT.RefreshTokenTTL, "jwt_refresh_token_ttl", 24*time.Hour, "jwt refresh token ttl")
	_ = viper.BindPFlag("JWT.Secret", confCmd.Lookup("jwt_secret"))
	_ = viper.BindPFlag("JWT.AccessTokenTTL", confCmd.Lookup("jwt_access_token_ttl"))
	_ = viper.BindPFlag("JWT.RefreshTokenTTL", confCmd.Lookup("jwt_refresh_token_ttl"))

	confCmd.StringVar(&config.GlobalConfig.Email.Sender, "email_sender", "kuanyshev_eldar@mail.ru", "email sender")
	confCmd.StringVar(&config.GlobalConfig.Email.Password, "email_password", "GrKrsLTy5rzGqkFC8ZhX", "email password")
	confCmd.StringVar(&config.GlobalConfig.Email.SMTPHost, "email_smtp_host", "smtp.mail.ru", "email smtp host")
	confCmd.StringVar(&config.GlobalConfig.Email.SMTPPort, "email_smtp_port", "587", "email smtp port")
	_ = viper.BindPFlag("EMAIL.Sender", confCmd.Lookup("email_sender"))
	_ = viper.BindPFlag("EMAIL.Password", confCmd.Lookup("email_password"))
	_ = viper.BindPFlag("EMAIL.SMTPHost", confCmd.Lookup("email_smtp_host"))
	_ = viper.BindPFlag("EMAIL.SMTPPort", confCmd.Lookup("email_smtp_port"))

	confCmd.StringVar(&config.GlobalConfig.Recognizer.URL, "recognizer_url", "https://api.platerecognizer.com", "recognizer url")
	confCmd.StringVar(&config.GlobalConfig.Recognizer.Token, "recognizer_token", "f548497ecd62d281f5ee97e497dd67426bb8a586", "recognizer token")
	_ = viper.BindPFlag("RECOGNIZER.URL", confCmd.Lookup("recognizer_url"))
	_ = viper.BindPFlag("RECOGNIZER.Token", confCmd.Lookup("recognizer_token"))
}
