package config

import "github.com/stanislav-zeman/gocut/internal/gocut"

type Config struct {
	Commands []gocut.Command `yaml:"commands"`
}
