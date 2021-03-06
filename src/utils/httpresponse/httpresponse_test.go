package httpresponse

import (
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert.Nil(t, nil)
}

func TestExtractCraneError(t *testing.T) {
	var err error
	c, ok := extractCraneError(err)
	assert.False(t, ok)
	assert.Nil(t, c)

	var cranerror cranerror.CraneError
	cranerror.Code = "code"
	c, ok = extractCraneError(&cranerror)
	assert.True(t, ok)
	assert.NotNil(t, c)
}

func TestParseHttpCodeAndErrCode(t *testing.T) {
	hCode, errCode := parseHttpCodeAndErrCode("400-10000")
	assert.Equal(t, hCode, 400)
	assert.Equal(t, errCode, 10000)
}
