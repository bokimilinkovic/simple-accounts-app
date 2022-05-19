package viper

import (
	"github.com/spf13/viper"
)

// New returns a new viper configuration and applies optional settings to it.
func New(opts ...Option) *viper.Viper {

	v := viper.New()
	for _, opt := range opts {
		opt(v)
	}

	return v
}
