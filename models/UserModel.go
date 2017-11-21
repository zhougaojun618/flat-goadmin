package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int64	`pk:"auto"`
	UserName string `orm:"unique,size(30)" form:"userName" valid:"Required;MaxSize(20);MinSize(6)"`
	Password string `orm:"size(32)" form:"password"`
	Repassword string `orm:"-" form:"repassword"`
	Remark string `orm:"null;size(300)" form:"remark"`
	Status int `orm:"default(1)" form:"status"`
	CreateTime time.Time `orm:type(datetime)`
	Role []*Role `orm:"rel(m2m)"`
}

func  AddUser(user *User) (int64,error)  {
	user.CreateTime = time.Now()
	id,err := orm.NewOrm().Insert(user)
	return id,err
}

func TableName() (string){
	return "user"
}

func PageList(page,size int) ([]*User,int64){
	users := make([]*User, 0)
	qt := orm.NewOrm().QueryTable(TableName())
	total,_:= qt.Count()
	qt.OrderBy("-id").Limit(size,(page-1)*size).All(&users)
	return users,total
}