package util_test

import (
	"testing"

	"github.com/238Studio/child-nodes-assist/util"
)

func TestName(t *testing.T) {
	data := make([]byte, 4)
	targetModule := "test"
	msg := util.SendWebsocketNetMessage(&data, targetModule, true)
	t.Log(msg)
}
