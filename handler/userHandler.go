package handler

import (
	"github.com/labstack/echo/v4"
	"go_echo_ent/ent"
)

/**
UserHandler
*/

// 创建新用户
func createUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var _ ent.User
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		u := ent.User{
			Name:     m["name"].(string),
			Age:      m["age"].(int),
			Username: m["username"].(string),
			Email:    m["email"].(string),
			Password: m["password"].(string),
		}

		_ = u
		//_, err := ent.Client.User.Create()

		return nil
	}
}

// 根据用户ID查询用户
func getUserById() {

}

// 根据用户名查询用户
func getUserByUserName() {

}

//  根据用户邮箱进行查询用户
func getUserByEmail() {

}

//func getUserById()  {
//}

func getAllUser() {

}

func deleteUser() {

}

func updateUser() {

}
