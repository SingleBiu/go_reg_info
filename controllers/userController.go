package controllers

import (
	"regInfo/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type UsrController struct {
	beego.Controller
}

func (c *UsrController) Get() {
	c.TplName = "Check.html"
}

func (c *UsrController) HandleUsrPost() {
	//1.拿到数据
	Name := c.GetString("Name")
	Number := c.GetString("Number")
	UID_YUAN := c.GetString("UID1")
	UID_BENG := c.GetString("UID2")
	UID_JUE := c.GetString("UID3")
	println(Name, Number, UID_YUAN, UID_BENG, UID_JUE)
	//2.对数据进行校验
	if Name == "" || Number == "" || UID_YUAN == "" || UID_BENG == "" || UID_JUE == "" {
		c.Ctx.WriteString("未填写完全")
		c.Ctx.WriteString("请补全信息")
		c.Redirect("/Check", 302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()

	user := models.UserTable{}
	user.Name = Name
	user.Number = Number
	user.UID_YUAN = UID_YUAN
	user.UID_BENG = UID_BENG
	user.UID_JUE = UID_JUE
	//QQ号作为ID更新信息 查询到数据库中有该QQ号则更新信息
	exist := o.Read(&user, "number")
	if exist == nil {
		user.Name = Name
		user.UID_YUAN = UID_YUAN
		user.UID_BENG = UID_BENG
		user.UID_JUE = UID_JUE
		_, err := o.Update(&user)
		if err != nil {
			c.Ctx.WriteString("数据更新失败")
			c.Redirect("/Check", 302)
			return
		} else {
			c.Ctx.WriteString("数据更新成功")
			c.TplName = "Success.html"
			return
		}
	}
	_, err := o.Insert(&user)
	if err != nil {
		c.Ctx.WriteString("插入数据失败")
		c.Redirect("/Check", 302)
		return
	}
	//4.返回登陆界面
	c.TplName = "Success.html"
}

func (c *UsrController) ShowRegister() {
	c.TplName = "Check.html"
}
