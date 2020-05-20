package settings

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	usingConfigFile bool
)

func InitFromFile(filePath string) error {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	usingConfigFile = true
	return nil
}

// abc.def <==> PREFIX_ABC_DEF
func AutomaticEnv(prefix string) {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}
