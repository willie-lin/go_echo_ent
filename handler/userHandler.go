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
	"time"
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
			SetPassword(user.Password).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
		_, err = uc.Save(context.Background())
		if err != nil {
			return fmt.Errorf("failed createing the group: %V", err)
		}
		fmt.Println(uc)

		//u := client.User.Create()
		//u.SetName(user.Name)
		//u.SetUsername(user.Username)
		//u.SetAge(user.Age).SetEmail("").SetPassword("qqqqq").SetCreatedAt(time.Now()).SetUpdatedAt(time.Now())
		//_, err = u.Save(context.Background())
		//if err != nil {
		//	log.Fatal(err)
		//}

		//u := client.User.Create()
		//u.SetName("aaa")
		//u.SetUsername("bbb")
		//u.SetAge(11).SetEmail("").SetPassword("qqqqq").SetCreatedAt(time.Now()).SetUpdatedAt(time.Now())
		//_, err = u.Save(context.Background())
		//if err != nil {
		//	log.Fatal(err)
		//}

		//client.User.
		//	Create().
		//	SetAge(30).
		//	SetName("a8m").
		//	Save(context.Background())
		//if err != nil {
		//	return fmt.Errorf("failed creating user: %v", err)
		//}
		//log.Println("user was created: ", u)

		//u, err := client.User.Create().SetName(user.Name).SetAge(user.Age).Save(context.Background())
		//fmt.Println(u)
		//if err != nil {
		//	fmt.Errorf("failed creating the User: %v", err)
		//}
		//log.Println("user was created: ", &u)

		return c.JSON(http.StatusOK, &user)

	}
}

func CreateUser1(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
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
