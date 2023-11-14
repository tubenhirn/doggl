package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	assert.Equal(t, "doggl", rootCmd.Use)
	assert.Equal(t, "Doogl - A simple toggl cli.", rootCmd.Short)
	assert.Equal(t, "", cfgFile)
	assert.Equal(t, "", apitoken)
}
