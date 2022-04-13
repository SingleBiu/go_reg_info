package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"regInfo/models"
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
	Phone := c.GetString("Phone")
	Career1 := c.GetString("career1")
	Career2 := c.GetString("career2")
	println(Name, Number, Phone, Career1, Career2)
	//2.对数据进行校验
	if Name == "" || Number == "" || Phone == "" || Career1 == "" || Career2 == "" {
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
	user.Phone = Phone
	user.Career1 = Career1
	user.Career2 = Career2
	exist := o.Read(&user, "number")
	if exist == nil {
		user.Name = Name
		user.Phone = Phone
		user.Career1 = Career1
		user.Career2 = Career2
		_, err := o.Update(&user)
		if err != nil {
			c.Ctx.WriteString("数据更新失败")
			c.Redirect("/Check", 302)
			return
		} else {
			c.Ctx.WriteString("数据更新成功")
			c.TplName = "success.html"
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
	c.TplName = "success.html"
}

func (c *UsrController) ShowRegister() {
	c.TplName = "Check.html"
}
