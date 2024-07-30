package core

import "testing"

func TestLogger(t *testing.T) {
	logger := NewLogger(DebugLevel)
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
