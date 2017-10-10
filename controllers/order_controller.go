package controllers

import (
	"strconv"
	"time"
	"yuedan/enums"
	"yuedan/models"
	"yuedan/utils"

	"github.com/astaxie/beego"
)

// OrderController operations for Order
type OrderController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrderController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetOnesOrder", c.GetOnesOrder)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description 发布约单
// @Param	token	header	String	true	"user token"
// @Param	title	body	string	true	"title"
// @Param	type	body	string	true	"type"
// @Param	date	body	string	true	"date 2017-10-20 10:20:00"
// @Param	place	body	string	true	"place"
// @Param	pay	body	string	true	"pay"
// @Param	content	body	string	true	"content"
// @Param	number	body	string	true	"number"
// @Param	sex	body	string	true	"sex"
// @Success success { "Status": 0, "Error": "", "Data": "成功", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "用户未登录", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router / [post]
func (c *OrderController) Post() {
	userId, _ := c.Ctx.Input.GetData("userId").(int)
	var order models.Order

	if bindErr := c.ParseForm(&order); bindErr == nil {
		if order.Title == "" || order.Type == "" || order.Date.IsZero() || order.Place == "" || order.Pay == "" || order.Content == "" || order.Number == 0 || order.Sex == "" {
			c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "缺少参数")
		} else {
			order.ReleaseDate = time.Now()
			order.UserId = &models.User{Id: userId}

			if _, saveErr := models.AddOrder(&order); saveErr == nil {
				c.Data["json"] = utils.GenerateDataWrapper(enums.Success, "成功")
			} else {
				c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
			}
		}

	} else {
		c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "参数绑定错误")
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description 根据id获取约单
// @Param	token	header	String	true	"user token"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success success { "Status": 0, "Error": "", "Data": { "Id": 1, "Title": "title", "Type": "type", "Date": "2017-10-20T10:20:00+08:00", "ReleaseDate": "2017-10-10T20:05:13+08:00", "Place": "place", "Pay": "pay", "Content": "content", "UserId": { "Id": 2, "StudentId": "", "StudentPassword": "", "Password": "", "Phone": "", "Photo": "", "Nickname": "大潘", "RealName": "", "Sex": "", "Birth": "0001-01-01T00:00:00Z", "Email": "", "School": "", "SchoolPart": "", "College": "", "Dormitory": "", "Specialty": "", "Education": "", "SchoolYear": "", "Identity": "", "Hobby": "", "IsSingle": "", "AncestralHome": "", "Lvl": 0, "Exp": 0, "UserType": 0 }, "Number": 3, "Sex": "F" }, "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "用户未登录", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router /:id [get]
func (c *OrderController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrderById(id)
	if err != nil {
		c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "约单不存在")
	} else {
		v.UserId = &models.User{Id: v.UserId.Id, Nickname: v.UserId.Nickname}
		c.Data["json"] = utils.GenerateDataWrapper(enums.Success, v)
	}
	c.ServeJSON()
}

// GetOnesOrder ...
// @Title Get One's Order List
// @Description 获取某人发布的约单列表
// @Param	token	header	String	true	"user token"
// @Param	PageIndex		body 	string	true		"PageIndex  1~"
// @Param	PageSize		body 	string	true		"PageSize 1~10"
// @Success success { "Status": 0, "Error": "", "Data": [ { "Content": "content", "Date": "2017-10-30T10:20:00+08:00", "Id": 5, "NickName": "", "Number": 3, "Pay": "pay", "Place": "place", "ReleaseDate": "2017-10-10T22:03:51+08:00", "Sex": "F", "Title": "title4", "Type": "type", "UserId": 3 } ], "PageSize": 10, "PageIndex": 1, "TotalPage": 1, "TotalCount": 3 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "用户未登录", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router /myList [get]
func (c *OrderController) GetOnesOrder() {
	userId, _ := c.Ctx.Input.GetData("userId").(int)
	pageIndexStr := c.GetString("PageIndex")
	pageSizeStr := c.GetString("PageSize")

	pageIndex, _ := strconv.ParseInt(pageIndexStr, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)
	if pageSize < 1 || pageSize > 10 || pageIndex < 1 {
		c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "PageIndex或者PageSize范围错误")
	} else {
		totalCount, countErr := models.GetOrderCountByUserId(userId)
		if countErr == nil {

			orders, err := models.GetOrderByUserId(userId, (pageIndex-1)*pageSize, pageSize)
			if err == nil {
				c.Data["json"] = utils.GenerateDataWrapperWithPage(enums.Success, orders, pageSize, pageIndex, totalCount)
			} else {
				c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
			}

		} else {
			c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
		}

	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description 删除约单
// @Param	token	header	String	true	"user token"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success success { "Status": 0, "Error": "", "Data": "成功", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @Failure error { "Status": 1, "Error": "权限错误", "Data": "用户未登录", "PageSize": 0, "PageIndex": 0, "TotalPage": 0, "TotalCount": 0 }
// @router /:id [delete]
func (c *OrderController) Delete() {
	userId, _ := c.Ctx.Input.GetData("userId").(int)

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	order, findErr := models.GetOrderById(id)
	if findErr == nil {
		if order.UserId.Id == userId {
			if deleteErr := models.DeleteOrder(id); deleteErr == nil {
				c.Data["json"] = utils.GenerateDataWrapper(enums.Success, "成功")
			} else {
				c.Data["json"] = utils.GenerateDataWrapper(enums.DBError, "数据库错误")
			}

		} else {
			c.Data["json"] = utils.GenerateDataWrapper(enums.AuthError, "不能删除别人的约单")
		}

	} else {
		c.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "约单不存在")
	}
	c.ServeJSON()

}
