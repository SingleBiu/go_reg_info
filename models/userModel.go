package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserTable struct {
	ID      int    `orm:"pk;auto;"`
	Name    string `orm:"size(4);column(name)"`
	Number  string `orm:"size(12);column(number)"`
	Phone   string `orm:"size(11);column(phone)"`
	Career1 string `orm:"size(6);column(career1)"`
	Career2 string `orm:"size(6);column(career2)"`
}

func init() {
	//设置数据库基本信息	ORM必须注册一个别名为default的数据库作为默认使用
	err := orm.RegisterDataBase("default", "mysql", "root:NZzd952107@tcp(127.0.0.1:3306)/registerinfo?charset=utf8")
	if err != nil {
		return
	}
	//映射model数据
	orm.RegisterModel(new(UserTable))
	//自动建表	参数2是否更新表(drop重建)  参数3为是否开启创建表
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		return
	}

}
