package util_test

import (
	"errors"
	"testing"

	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
)

func TestError(t *testing.T) {
	err := util.NewError(_const.NotException, _const.NoError, false, errors.New("test"))

	t.Log(util.ErrToStruct(err))
	t.Log(util.ExtractErrorLevel(err))
	t.Log(util.ExtractErrorType(err))
	t.Log(util.ExtractErrorMessage(err))
	t.Log(util.ExtractStackTrace(err))
	t.Log(util.ExtractIsNeedSpecialHandel(err))
}
