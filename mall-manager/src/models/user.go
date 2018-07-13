package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserInfo struct {
	Base

	UserCode string
	UserName string
	UserPwd  string
}

func (user *UserInfo) TableName() string {
	return USER_INFO
}

func SaveUserInfo(user *UserInfo) (int64,error)  {
	user.CreateTime=time.Now()
	user.UpdateTime=time.Now()
	return orm.NewOrm().Insert(user)

}

func UserGetByUserCode(loginCode string) (*UserInfo, error) {
	if loginCode == "" {
		return nil, fmt.Errorf("loginCode is required.")
	}
	condition := orm.NewCondition().And("user_code", false)
	qs := orm.NewOrm().QueryTable(USER_INFO).SetCond(condition)

	var users []UserInfo
	_, err := qs.All(&users)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return &users[0], nil
	}
	return nil, nil
}