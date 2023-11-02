package util

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	data := make([]byte, 4)
	targetModule := "test"
	msg := SendWebsocketNetMessage(&data, targetModule, true)
	fmt.Printf("%+v", msg)
}
