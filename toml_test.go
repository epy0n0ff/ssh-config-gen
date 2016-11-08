package sshconfig

import "testing"
import "github.com/stretchr/testify/assert"

func TestLoadTomlSuccess(t *testing.T) {
	c, err := LoadToml("testdata/hosts.toml")
	assert.NoError(t, err)
	assert.NotNil(t, c)

	assert.Equal(t, 22, c.Port)
	assert.Equal(t, "hogehoge", c.User)
	assert.Equal(t, "/Users/hoge/.ssh/id_rsa", c.IdentityFile)
}

func TestLoadTomlFail(t *testing.T) {
	c, err := LoadToml("testdata/notfound.toml")
	assert.Error(t, err)
	assert.Nil(t, c)
}
