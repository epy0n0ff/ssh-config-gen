package sshconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSshConfigSuccess(t *testing.T) {
	var sshconfig = `Host hoge1.example.com
  HostName 127.0.0.1
  User hogehoge
  Port 22
  UserKnownHostsFile /dev/null
  StrictHostKeyChecking no
  PasswordAuthentication no
  IdentityFile /Users/hoge/.ssh/id_rsa
  IdentitiesOnly yes
  LogLevel FATAL

Host hoge2.example.com
  HostName 127.0.0.2
  User hogehoge
  Port 22
  UserKnownHostsFile /dev/null
  StrictHostKeyChecking no
  PasswordAuthentication no
  IdentityFile /Users/hoge/.ssh/id_rsa
  IdentitiesOnly yes
  LogLevel FATAL

`
	toml, err := LoadToml("testdata/hosts.toml")
	assert.NoError(t, err)
	b, err := GetSshConfig("testdata/sshconfig.tpl", toml)
	assert.NoError(t, err)
	assert.Equal(t, sshconfig, b.String())
}
