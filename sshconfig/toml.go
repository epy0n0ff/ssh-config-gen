package sshconfig

import (
	"github.com/BurntSushi/toml"
)

type tomlSshConfig struct {
	User         string
	Port         int
	IdentityFile string
	Hosts        [][]string
}

type tomlConfig struct {
	SshConfig *tomlSshConfig
}

func LoadToml(filepath string) (*tomlSshConfig, error) {
	var c tomlConfig
	_, err := toml.DecodeFile(filepath, &c)
	if err != nil {
		return nil, err
	}

	return c.SshConfig, nil
}
