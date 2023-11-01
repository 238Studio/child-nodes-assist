package util_test

import (
	"errors"
	"log"
	"testing"

	"github.com/UniversalRobotDriveTeam/child-nodes-util/util"
)

func TestNewError(t *testing.T) {
	err := util.NewError(1, 1, false, errors.New("test"))

	log.Println(util.ErrToStruct(err))
	log.Println(util.ExtractStackTrace(err))
	log.Println(util.ExtractErrorMessage(err))
	log.Println(util.ExtractErrorType(err))
	log.Println(util.ExtractErrorLevel(err))
	log.Println(util.ExtractIsNeedSpecialHandle(err))
}
