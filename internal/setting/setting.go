package setting

import (
	"github.com/kelseyhightower/envconfig"
)

type Setting struct {
	MaxHistory int `envconfig:"MAX_HISTORY" default:"1000"`
}

func Load() (Setting, error) {
	var s Setting
	err := envconfig.Process("", &s)
	return s, err
}
