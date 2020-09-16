package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"go_echo_ent/datasource"
	"go_echo_ent/ent"
	"go_echo_ent/ent/user"
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
		u := new(ent.User)

		result, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll :", err)
			return err
		}
		err = json.Unmarshal(result, &u)
		if err != nil {
			fmt.Println("json.Unmarshal ", err)
			return err
		}

		us, err := client.User.
			Update().
			Where(user.UsernameEQ(u.Username)).
			SetAge(u.Age).Save(context.Background())
		if err != nil {
			return fmt.Errorf("failed querying user: %v", err)
		}
		log.Println("user returned: ", us)
		return c.JSON(http.StatusOK, us)

	}

}

// 查询所有用户
func GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		client, err := datasource.Clients()

		if err != nil {
			panic(err)
		}
		//user := new(ent.User)
		users, err := client.User.Query().All(context.Background())

		return c.JSON(http.StatusOK, users)
	}
}

// 根据用户名查询用户
func GetUserByName() echo.HandlerFunc {
	return func(c echo.Context) error {

		client, err := datasource.Clients()

		if err != nil {
			panic(err)
		}

		un := c.FormValue("name")

		fmt.Println(un)
		us, err := client.User.Query().Where(user.NameEQ(un)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("failed querying user: %v", err)
		}
		log.Println("user returned: ", us)
		return c.JSON(http.StatusOK, us)

	}
}

// 根据用户名查询用户
func GetUserByUserName() echo.HandlerFunc {
	return func(c echo.Context) error {

		client, err := datasource.Clients()
		//user := new(ent.User)

		if err != nil {
			panic(err)
		}
		un := c.FormValue("username")
		fmt.Println(un)
		us, err := client.User.Query().Where(user.UsernameEQ(un)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("failed querying user: %v", err)
		}
		log.Println("user returned: ", us)
		return c.JSON(http.StatusOK, us)
	}
}

//  根据用户邮箱进行查询用户
func GetUserByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

//func getUserById()  {
//}

func deleteUser() {

}
