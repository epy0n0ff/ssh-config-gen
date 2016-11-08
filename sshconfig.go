package sshconfig

import (
	"bytes"
	"text/template"
)

type templateConfig struct {
	Host, HostName, SshUser, IdentityFile string
	SshPort                               int
}

func GetSshConfig(filepath string, config *tomlSshConfig) (*bytes.Buffer, error) {
	t := template.Must(template.ParseFiles(filepath))
	b := bytes.NewBufferString("")

	for _, host := range config.Hosts {
		var tpl = &templateConfig{
			Host:         host[0],
			HostName:     host[1],
			SshUser:      config.User,
			SshPort:      config.Port,
			IdentityFile: config.IdentityFile,
		}
		err := t.Execute(b, tpl)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}
