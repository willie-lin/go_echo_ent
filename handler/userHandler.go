package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"go_echo_ent/datasource"
	"go_echo_ent/ent"
	"io/ioutil"
	"net/http"
)

/**
UserHandler
*/

// 创建新用户
func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		client, err := datasource.Clients()

		if err != nil {
			panic(err)
		}
		user := new(ent.User)

		result, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll :", err)
			return err
		}
		err = json.Unmarshal(result, &user)
		if err != nil {
			fmt.Println("json.Unmarshal ", err)
			return err
		}
		fmt.Println("user:", &user)
		fmt.Println("user:", user.Name)

		uc := client.User.Create()

		uc.SetAge(user.Age).SetName(user.Name).SetUsername(user.Username).SetEmail(user.Email).
			SetPassword(user.Password)
			//SetCreatedAt(time.Now()).
			//SetUpdatedAt(time.Now())
		_, err = uc.Save(context.Background())
		if err != nil {
			return fmt.Errorf("failed createing the group: %V", err)
		}
		fmt.Println(uc)

		return c.JSON(http.StatusOK, &user)

	}
}

// 更新用户
func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		client, err := datasource.Clients()

		if err != nil {
			panic(err)
		}
		user := new(ent.User)

		result, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll :", err)
			return err
		}
		err = json.Unmarshal(result, &user)
		if err != nil {
			fmt.Println("json.Unmarshal ", err)
			return err
		}

		us, err := client.User.UpdateOne(user).SetAge(user.Age).Save(context.Background())
		if err != nil {
			return fmt.Errorf("failed update the user: %V", err)

		}
		return c.JSON(http.StatusOK, us)

	}

}

// 根据用户ID查询用户
func GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {

		//client, err := datasource.Clients()
		//
		//if err != nil {
		//	panic(err)
		//}

		//un := c.FormValue("username")
		un := c.Param("username")

		//user := new(ent.User)
		////
		//result, err := ioutil.ReadAll(c.Request().Body)
		//if err != nil {
		//	fmt.Println("ioutil.ReadAll :", err)
		//	return err
		//}
		//err = json.Unmarshal(result, &user)
		//if err != nil {
		//	fmt.Println("json.Unmarshal ", err)
		//	return err
		//}
		////fmt.Println("user:", &user)
		//fmt.Println("user:", user.Username)

		fmt.Println("un:", un)
		//u, err := client.User.Query().Where(name)

		return c.JSON(http.StatusOK, un)

	}
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
