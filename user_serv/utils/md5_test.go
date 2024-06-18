package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, Md5Sum("123456"), "1d4e8b12763757ef90fd62a03d989e5e")
}
