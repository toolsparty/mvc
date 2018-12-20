package mvc

import "time"

// app config interface
// see github.com/spf13/viper
type Config interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
}

type AppConfig struct {
	Config Config
	Router Router
	Logger LogFunc

	Controllers []Controller
	Models      []Model
	Views       []View
}
