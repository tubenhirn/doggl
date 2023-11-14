package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeStringToDuration(t *testing.T) {
	duration, _ := timeStringToDuration("8h")
	assert.Equal(t, int64(28800), duration)

	duration, _ = timeStringToDuration("8h25m")
	assert.Equal(t, int64(30300), duration)

	duration, _ = timeStringToDuration("1m")
	assert.Equal(t, int64(60), duration)

	duration, _ = timeStringToDuration("")
	assert.Equal(t, int64(0), duration)

	duration, err := timeStringToDuration("28800")
	assert.Equal(t, int64(0), duration)
	assert.NotNil(t, err)
	assert.Equal(t, "time: missing unit in duration \"28800\"", err.Error())
}
