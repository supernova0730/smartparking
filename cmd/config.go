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

	confCmd.StringVar(&config.GlobalConfig.Cache.Host, "cache_host", "localhost", "cache host")
	confCmd.StringVar(&config.GlobalConfig.Cache.Port, "cache_port", "11211", "cache port")
	confCmd.DurationVar(&config.GlobalConfig.Cache.TTL, "cache_ttl", 10*time.Minute, "cache host")

	confCmd.StringVar(&config.GlobalConfig.Web.Host, "web_host", "localhost", "web host")
	confCmd.StringVar(&config.GlobalConfig.Web.Port, "web_port", "8080", "web port")
	confCmd.BoolVar(&config.GlobalConfig.Web.TLSEnable, "web_tlsenable", false, "web tls enable")
	confCmd.StringVar(&config.GlobalConfig.Web.CertFile, "web_certfile", "", "web cert file")
	confCmd.StringVar(&config.GlobalConfig.Web.KeyFile, "web_keyfile", "", "web key file")
	confCmd.StringVar(&config.GlobalConfig.Web.FileStorage, "web_file_storage", filepath.Join(tools.HomePath(), ".smartparking/images"), "web file storage")

	confCmd.StringVar(&config.GlobalConfig.JWT.SecretKey, "jwt_secret", "my_secret_key", "jwt secret")
	confCmd.DurationVar(&config.GlobalConfig.JWT.AccessTokenTTL, "jwt_access_token_ttl", 15*time.Minute, "jwt access token ttl")
	confCmd.DurationVar(&config.GlobalConfig.JWT.RefreshTokenTTL, "jwt_refresh_token_ttl", 24*time.Hour, "jwt refresh token ttl")

	confCmd.StringVar(&config.GlobalConfig.Email.Sender, "email_sender", "kuanyshev_eldar@mail.ru", "email sender")
	confCmd.StringVar(&config.GlobalConfig.Email.Password, "email_password", "GrKrsLTy5rzGqkFC8ZhX", "email password")
	confCmd.StringVar(&config.GlobalConfig.Email.SMTPHost, "email_smtp_host", "smtp.mail.ru", "email smtp host")
	confCmd.StringVar(&config.GlobalConfig.Email.SMTPPort, "email_smtp_port", "587", "email smtp port")

	confCmd.StringVar(&config.GlobalConfig.Recognizer.URL, "recognizer_url", "https://api.platerecognizer.com", "recognizer url")
	confCmd.StringVar(&config.GlobalConfig.Recognizer.Token, "recognizer_token", "f548497ecd62d281f5ee97e497dd67426bb8a586", "recognizer token")
}
