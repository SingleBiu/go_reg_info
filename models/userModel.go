/*
 * @Author: SingleBiu
 * @Date: 2024-11-17 12:14:19
 * @LastEditors: SingleBiu
 * @LastEditTime: 2024-11-17 13:15:18
 * @Description: file content
 */
package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserTable struct {
	ID       int    `orm:"pk;auto;"`
	Name     string `orm:"size(4);column(name)"`
	Number   string `orm:"size(12);column(number)"`
	UID_YUAN string `orm:"size(11);column(UID1)"`
	UID_BENG string `orm:"size(11);column(UID2)"`
	UID_JUE  string `orm:"size(11);column(UID3)"`
}

func init() {
	//设置数据库基本信息	ORM必须注册一个别名为default的数据库作为默认使用
	//创建 default数据库 以及 项目数据库
	//[ CREATE DATABASE 'default'; ]
	//[ CREATE DATABASE registerInfo;]
	err := orm.RegisterDataBase("default", "mysql", "root:zdcanonvv@tcp(127.0.0.1:3306)/registerInfo?charset=utf8")
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
