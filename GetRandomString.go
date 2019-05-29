package utils

import (
	"math/rand"
	"time"
)

var str = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func GetRandomString() string {
	var result string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result += str[r.Intn(62)]
	}
	return result
}

// beego拦截器添加processId,不用再每个controller单独生成processId
/*
package controllers

import (
	"github.com/whaley/kids-education/kids_user_center/utils"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"time"
)

func init() {
	var AddProcessId = func(ctx *context.Context) {
		processId := time.Now().Format("20060102150405") + utils.GetRandString()
		ctx.Input.SetData("processId", processId)
	}
	beego.InsertFilter("/*", beego.BeforeRouter, AddProcessId)
}

后续可直接从request中拿：processId := this.GetString("processId")
*/

// github.com/ngaut/log通过SetPrefix添加processId,不用每行log都添加processId字段
/*
log.Logger().SetPrefix(fmt.Sprintf("processId=%v  ", processId))
*/

// gin亦可在拦截器中添加processId
/*
func GenerateProcessId() gin.HandlerFunc {
	return func(c *gin.Context) {
		processId := time.Now().Format("20060102150405") + utils.GetRandString()
		c.Set("processId", processId)
		c.Next()
	}
}
后续可直接从request中拿：processId := c.GetString("processId")
*/
