// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"wizz-home-page/app/model/internal"

	"golang.org/x/crypto/bcrypt"
)

// User is the golang structure for table user.
type User internal.User

// Fill with you ideas below.

// 注册请求参数，用于前后端交互参数格式约定
type UserApiSignUpReq struct {
	Username string `v:"required|length:5,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

// 注册输入参数
type UserServiceSignUpReq struct {
	Username string
	Password string
}

type UserApiLoginReq struct {
	Username string `v:"required|length:5,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

type UserServiceLoginReq struct {
	Username string
	Password string
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

func (user *UserServiceSignUpReq) EncryptPassword() error {
	if cipher, err := EncryptPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = cipher
		return nil
	}
}

func EncryptPassword(str string) (string, error) {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(str), PassWordCost); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func CheckPassword(plain string, cipher string) bool {
	// g.Log().Line().Debug(plain)
	// g.Log().Line().Debug(cipher)
	return bcrypt.CompareHashAndPassword([]byte(cipher), []byte(plain)) == nil
}
