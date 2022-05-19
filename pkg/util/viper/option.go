package viper

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Option sets the optional setting to viper.Viper.
type Option func(*viper.Viper)

// WithEnvKeyReplacer sets the environment key replacer to viper.
// It uses strings replacer to reformat the expected environment
// variable name. For example to replace dots with underscores.
func WithEnvKeyReplacer(r *strings.Replacer) Option {
	return func(v *viper.Viper) {
		v.SetEnvKeyReplacer(r)
	}
}

// WithDefault sets a default value for a key.
func WithDefault(k string, v interface{}) Option {
	return func(v *viper.Viper) {
		v.SetDefault(k, v)
	}
}

// WithEnvPrefix sets prefix for every environment variable.
func WithEnvPrefix(p string) Option {
	return func(v *viper.Viper) {
		v.SetEnvPrefix(p)
	}
}

// WithAutomaticEnv checks ENV variables for all keys.
func WithAutomaticEnv() Option {
	return func(v *viper.Viper) {
		v.AutomaticEnv()
	}
}

// WithBindEnv binds a viper key to ENV variable.
func WithBindEnv(input ...string) Option {
	return func(v *viper.Viper) {
		v.BindEnv(input...)
	}
}

// WithEmptyEnv consider empty ENVs as valid.
func WithEmptyEnv() Option {
	return func(v *viper.Viper) {
		v.AllowEmptyEnv(true)
	}
}

// WithConfigFile sets configuration file to read.
func WithConfigFile(file string) Option {
	return func(v *viper.Viper) {
		v.SetConfigFile(file)
	}
}

// WithConfigFileFromEnv reads the configuration file name from
// the given environment variable with a default being used in
// case env var is not set.
func WithConfigFileFromEnv(env, def string) Option {
	return func(v *viper.Viper) {
		file := os.Getenv(env)
		if file == "" {
			file = def
		}
		v.SetConfigFile(file)
	}
}
