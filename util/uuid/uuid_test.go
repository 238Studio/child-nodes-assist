package uuid_test

import (
	"github.com/238Studio/child-nodes-assist/util/uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	t.Log(uuid.GenerateUUID())
}
