package titan_client

import "testing"

func TestNewLog(t *testing.T) {
	err := NewLog().SetLevel("DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	logger.Debugf("this debug")
	logger.Info("this info")
	logger.Warn("this warn")
	logger.Error("this error")
}
