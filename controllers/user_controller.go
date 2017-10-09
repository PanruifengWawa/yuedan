package controllers

import (
	"yuedan/enums"
	"yuedan/models"
	"yuedan/utils"

	"github.com/astaxie/beego"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Login", c.Login)
}

// Post ...
// @Title register
// @Description user register
// @Param	StudentId	body	string	true	"student id"
// @Param	StudentPassword	body	string	true	"student password"
// @Param	Phone	body	string	true	"phone"
// @Param	School	body	string	true	"school name"
// @Success success { "Status": 0, "Error": "", "Data": "成功", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "参数错误", "Data": "用户已存在", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	if err := c.ParseForm(&v); err == nil {
		if v.StudentId == "" || v.StudentPassword == "" || v.Phone == "" || v.School == "" {
			c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "缺少参数")
		} else {
			if _, err := models.GetUserByStudentId(v.StudentId); err != nil {
				v.StudentPassword = utils.GetMD5Coding(v.StudentPassword)
				v.Password = v.StudentPassword
				v.UserType = enums.User
				if _, err := models.AddUser(&v); err == nil {
					c.Data["json"] = utils.GenerateDataWrapper(enums.Success, "成功")
				} else {
					c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
				}

			} else {
				c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "用户已存在")
			}
		}

	} else {
		c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "参数绑定错误")
	}
	c.ServeJSON()
}

// Login ...
// @Title login
// @Description user login
// @Param	StudentId	body	string	true	"student id"
// @Param	Password	body	string	true	"password"
// @Success success { "Status": 0, "Error": "", "Data": "c0b7a6dd-552c-482e-b569-bcca89de5146", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "账号或密码错误", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router /login [post]
func (c *UserController) Login() {
	studentId := c.GetString("StudentId")
	password := c.GetString("Password")
	if studentId == "" || password == "" {
		c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "缺少参数")
	} else {
		if user, err := models.GetUserByStudentId(studentId); err == nil {
			if user.Password == utils.GetMD5Coding(password) {

				token, tokenError := models.GetTokenByUserId(user.Id)

				var resultError error
				if tokenError != nil {
					token = &models.Token{}

					token.UserId = user
					token.TokenStr = utils.GetUUId()
					token.UserType = user.UserType
					_, resultError = models.AddToken(token)
				} else {
					token.TokenStr = utils.GetUUId()
					resultError = models.UpdateTokenById(token)
				}
				if resultError == nil {
					c.Data["json"] = utils.GenerateDataWrapper(enums.Success, token.TokenStr)
				} else {
					c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
				}

			} else {
				c.Data["json"] = utils.GenerateDataWrapper(enums.AuthError, "账号或密码错误")
			}

		} else {
			c.Data["json"] = utils.GenerateDataWrapper(enums.AuthError, "账号或密码错误")
		}
	}

	c.ServeJSON()

}
