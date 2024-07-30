package core

import (
	"github.com/luolayo/gin-study/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Some initialization done before testing
func init() {
	InitViper(enum.ConfigDevelopmentPath)
}

// TestGetGorm tests the GetGorm method
// Test whether the return of client is empty when the connection fails due to configuration errors or lack of configuration
func TestGetGorm_ConfigErr(t *testing.T) {
	testGormsClient := GetGorm()
	assert.Equal(t, false, testGormsClient.CheckGormConnection())
}
