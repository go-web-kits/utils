package settings

import (
	"encoding/json"
	"strings"

	"github.com/go-web-kits/utils/mapx"
	"github.com/spf13/viper"
)

func Scan(key string, obj interface{}) error {
	if usingConfigFile {
		return viper.UnmarshalKey(key, obj)
	}
	return json.Unmarshal([]byte(String(key)), obj)
}

// Map("database.pg")
func Map(key string) map[string]interface{} {
	if usingConfigFile {
		return viper.GetStringMap(key)
	}

	result := map[string]interface{}{}
	err := json.Unmarshal([]byte(String(key)), &result)
	if err != nil {
		panic(err)
	}
	return result
}

// String("database.pg.host")
func String(key string) string {
	if usingConfigFile {
		return viper.GetString(key)
	}

	keys := strings.Split(key, ".")
	if len(keys) == 1 {
		return "" // config center
	}
	return mapx.Dig(Map(keys[0]), keys[1:]...).(string)
}

// Int("database.pg.connection_pool")
func Int(key string) int {
	if usingConfigFile {
		return viper.GetInt(key)
	}

	keys := strings.Split(key, ".")
	if len(keys) == 1 {
		return 0 // config center
	}
	return int(mapx.Dig(Map(keys[0]), keys[1:]...).(float64))
}
