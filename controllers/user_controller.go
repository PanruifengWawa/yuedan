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
	c.Mapping("Update", c.Update)
	c.Mapping("GetUserDetails", c.GetUserDetails)
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

// Get User Details ...
// @Title Get User Details
// @Description Get User Details
// @Param	token	header	String	true	"user token"
// @Success success { "Status": 0, "Error": "", "Data": "c0b7a6dd-552c-482e-b569-bcca89de5146", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "账号或密码错误", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router /details [get]
func (c *UserController) GetUserDetails() {
	userId, _ := c.Ctx.Input.GetData("userId").(int)
	user, findUserErr := models.GetUserById(userId)
	if findUserErr == nil {
		c.Data["json"] = utils.GenerateDataWrapper(enums.Success, user)
	} else {
		c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
	}
	c.ServeJSON()

}

// Update ...
// @Title update user
// @Description update user
// @Param	token	header	String	true	"user token"
// @Param	Photo	body	string	false	"Photo"
// @Param	Nickname	body	string	false	"Nickname"
// @Param	Password	body	string	false	"Password"
// @Param	RealName	body	string	false	"RealName"
// @Param	Sex	body	string	false	"Sex"
// @Param	Birth	body	date	false	"Birth"
// @Param	Email	body	string	false	"Email"
// @Param	School	body	string	false	"School"
// @Param	SchoolPart	body	string	false	"SchoolPart"
// @Param	College	body	string	false	"College"
// @Param	Specialty	body	string	false	"Specialty"
// @Param	Dormitory	body	string	false	"Dormitory"
// @Param	Education	body	string	false	"Education"
// @Param	SchoolYear	body	string	false	"SchoolYear"
// @Param	Identity	body	string	false	"Identity"
// @Param	Hobby	body	string	false	"Hobby"
// @Param	IsSingle	body	string	false	"IsSingle"
// @Param	AncestralHome	body	string	false	"AncestralHome"
// @Success success { "Status": 0, "Error": "", "Data": "c0b7a6dd-552c-482e-b569-bcca89de5146", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "账号或密码错误", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router /update [post]
func (c *UserController) Update() {
	userId, _ := c.Ctx.Input.GetData("userId").(int)
	userInDB, findUserErr := models.GetUserById(userId)
	if findUserErr == nil {
		var user models.User
		if bindErr := c.ParseForm(&user); bindErr == nil {
			if user.Photo != "" {
				userInDB.Photo = user.Photo
			}
			if user.Nickname != "" {
				userInDB.Nickname = user.Nickname
			}
			if user.Password != "" {
				userInDB.Password = utils.GetMD5Coding(user.Password)
			}
			if user.RealName != "" {
				userInDB.RealName = user.RealName
			}
			if user.Sex != "" {
				userInDB.Sex = user.Sex
			}
			if !user.Birth.IsZero() {
				userInDB.Birth = user.Birth
			}
			if user.Email != "" {
				userInDB.Email = user.Email
			}
			if user.School != "" {
				userInDB.School = user.School
			}
			if user.SchoolPart != "" {
				userInDB.SchoolPart = user.SchoolPart
			}
			if user.College != "" {
				userInDB.College = user.College
			}
			if user.Specialty != "" {
				userInDB.Specialty = user.Specialty
			}
			if user.Dormitory != "" {
				userInDB.Dormitory = user.Dormitory
			}
			if user.Education != "" {
				userInDB.Education = user.Education
			}
			if user.SchoolYear != "" {
				userInDB.SchoolYear = user.SchoolYear
			}
			if user.Identity != "" {
				userInDB.Identity = user.Identity
			}
			if user.Hobby != "" {
				userInDB.Hobby = user.Hobby
			}
			if user.IsSingle != "" {
				userInDB.IsSingle = user.IsSingle
			}
			if user.AncestralHome != "" {
				userInDB.AncestralHome = user.AncestralHome
			}

			updateErr := models.UpdateUserById(userInDB)
			if updateErr != nil {
				c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
			} else {
				userInDB.StudentPassword = ""
				userInDB.Password = ""
				c.Data["json"] = utils.GenerateDataWrapper(enums.Success, userInDB)
			}
		} else {
			c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "数据绑定错误")
		}
	} else {
		c.Data["json"] = utils.GenerateDataWrapper(enums.AuthError, "用户不存在")
	}
	c.ServeJSON()

}
