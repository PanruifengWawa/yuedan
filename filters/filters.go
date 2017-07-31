package filters

import (
	"time"
	"yuedan/enums"
	"yuedan/models"
	"yuedan/utils"

	"github.com/astaxie/beego/context"
)

func UserFilter(ctx *context.Context) {
	tokenStr := ctx.Input.Header("token")
	//	var errorMessage string
	if tokenStr == "" {
		ctx.Output.JSON(utils.GenerateDataWrapper(enums.ParameterError, "认证参数为空"), true, true)
		return
	}

	token, err := models.GetTTokenByTokenStr(tokenStr)
	if err != nil {
		ctx.Output.JSON(utils.GenerateDataWrapper(enums.AuthError, "用户未登录"), true, true)
		return
	} else {
		loginDate := token.LoginDate
		minutes := time.Now().Sub(loginDate).Minutes()
		if minutes > 30 || minutes < -30 {
			ctx.Output.JSON(utils.GenerateDataWrapper(enums.AuthError, "登录过期，请重新登录"), true, true)
			return
		}
		ctx.Input.SetData("userId", token.UserId.Id)
	}

}
