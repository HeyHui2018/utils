package utils

import (
	"fmt"
	"github.com/astaxie/beego"
)

func LogWithPId4BeegoInfo(processId, format string, v ...interface{}) {
	beego.BeeLogger.SetLogFuncCallDepth(3)
	beego.BeeLogger.Info(fmt.Sprintf("processId = %v,", processId)+format, v...)
}

func LogWithPId4BeegoWarn(processId, format string, v ...interface{}) {
	beego.BeeLogger.SetLogFuncCallDepth(3)
	beego.BeeLogger.Warn(fmt.Sprintf("processId = %v,", processId)+format, v...)
}
