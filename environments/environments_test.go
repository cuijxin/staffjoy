package environments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultWorks(t *testing.T) {
	assert := assert.New(t)
	explicitDefaultConfig, err := GetConfig(DefaultEnv)
	assert.NoError(err)
	implicitDefaultConfig, err := GetConfig("")
	assert.NoError(err)
	assert.Equal(explicitDefaultConfig, implicitDefaultConfig)
}
