package conf

import (
	"strings"

	"bitbucket.org/ironstar/tokaido-cli/system/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config the application's configuration
type Config struct {
	Port              string
	Config            string
	Project           string
	Path              string
	Force             bool
	Debug             bool
	Version           bool
	CustomCompose     bool
	SystemdPath       string
	LaunchdPath       string
	CreateSyncService bool
}

// LoadConfig loads the config from a file if specified, otherwise from the environment
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("TOK")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("CreateSyncService", true)
	viper.SetDefault("SystemdPath", fs.HomeDir()+"/.config/systemd/user/")
	viper.SetDefault("LaunchdPath", fs.HomeDir()+"/Library/LaunchAgents/")

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./.tok/")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/.tok/")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return populateConfig(new(Config))
}
