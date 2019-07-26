package utils

import (
	"fmt"
	"github.com/astaxie/beego"
)

func LogWithTId4BeegoInfo(processId, format string, v ...interface{}) {
	beego.BeeLogger.SetLogFuncCallDepth(3)
	beego.BeeLogger.Info(fmt.Sprintf("processId = %v,", processId)+format, v...)
}

func LogWithTId4BeegoWarn(processId, format string, v ...interface{}) {
	beego.BeeLogger.SetLogFuncCallDepth(3)
	beego.BeeLogger.Warn(fmt.Sprintf("processId = %v,", processId)+format, v...)
}

/*
"github.com/ngaut/log"
在各协程起始处添加如下语句：
log.Logger().SetPrefix("123333 ")
效果如下：
123333 2019/07/26 16:20:35 logWithTraceId.go:25: [info] hello world,it's testing,id = 123333
*/
