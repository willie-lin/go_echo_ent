package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"go_echo_ent/datasource"
	"go_echo_ent/ent"
	"io/ioutil"
	"log"
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

		u, err := client.User.Create().SetName(user.Name).SetAge(user.Age).Save(context.Background())
		fmt.Println(u)
		if err != nil {
			fmt.Errorf("failed creating the User: %v", err)
		}
		log.Println("user was created: ", &u)

		return c.JSON(http.StatusOK, &user)

	}
}

func CreateUsers(user *ent.User, ctx context.Context, client *ent.Client) (*ent.User, error) {

	u, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
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
