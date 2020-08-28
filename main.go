package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go_echo_ent/datasource"
	"go_echo_ent/ent"
	"go_echo_ent/ent/user"
	"go_echo_ent/handler"
	"log"
	"net/http"
)

func main() {

	e := echo.New()

	fmt.Println("aaaaa")

	client, err := datasource.Clients()

	if err != nil {
		panic(err)
	}
	fmt.Println("dddd")

	fmt.Println(client)
	fmt.Println("eeee")
	ctx := context.Background()

	migrate := datasource.Migrate
	migrate(client, ctx)

	//CreateUser(ctx, client)
	//println(CreateUser(ctx, client_model))

	fmt.Println("bbbbb")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!!!")
	})
	e.POST("/ww", handler.CreateUser())
	e.GET("/user/:username", handler.GetUserById())

	e.Logger.Fatal(e.Start(":2020"))

}

//func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
//	u, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed creating user: %v", err)
//	}
//	log.Println("user was created: ", u)
//	return u, nil
//}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().Where(user.NameEQ("a8m")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil

}

//func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
//	// creating new car with model "Tesla"
//	tesla, err := client.Car.
//		Create().SetModel("Tesla").
//		SetRegisteredAt(time.Now()).Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed creating car: %v", err)
//	}
//
//	// creating new car with model "Ford".
//	ford, err := client.Car.
//		Create().
//		SetModel("Ford").
//		SetRegisteredAt(time.Now()).Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed creating car: %v", err)
//	}
//	log.Println("car was created: ", ford)
//
//	// create a new user, and add it the 2 cars.
//	a8m, err := client.User.
//		Create().
//		SetAge(30).
//		SetName("a8m").
//		AddCars(tesla, ford).
//		Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed creating user: %v", err)
//
//	}
//	log.Println("user was created: ", a8m)
//	return a8m, nil
//}
//
//func QueryCars(ctx context.Context, a8m *ent.User) error {
//	cars, err := a8m.QueryCars().All(ctx)
//	if err != nil {
//		return fmt.Errorf("failed querying user cars: %v", err)
//	}
//	log.Println("returned cars:", cars)
//
//	// what about filtering specific cars.
//	ford, err := a8m.QueryCars().
//		Where(car.ModelEQ("Ford")).
//		Only(ctx)
//	if err != nil {
//		return fmt.Errorf("failed querying user cars: %s", err)
//
//	}
//	log.Println(ford)
//	return nil
//}
//
//func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
//	cars, err := a8m.QueryCars().All(ctx)
//	if err != nil {
//		return fmt.Errorf("failed querying user cars: %v", err)
//	}
//	for _, ca := range cars {
//		owner, err := ca.QueryOwner().Only(ctx)
//		if err != nil {
//			return fmt.Errorf("failed querying car %q owner : %v", ca.Model, err)
//		}
//		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
//	}
//	return nil
//}
//
//func CreateGraph(ctx context.Context, client *ent.Client) error {
//
//	a8m, err := client.User.
//		Create().SetAge(30).
//		SetName("Ariel").
//		Save(ctx)
//	if err != nil {
//		return err
//	}
//	neta, err := client.User.Create().SetAge(28).SetName("Neta").Save(ctx)
//	if err != nil {
//		return err
//	}
//
//	_, err = client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).SetOwner(a8m).Save(ctx)
//	if err != nil {
//		return err
//	}
//	_, err = client.Car.Create().SetModel("Mazda").SetRegisteredAt(time.Now()).SetOwner(a8m).Save(ctx)
//	if err != nil {
//		return err
//	}
//	_, err = client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).SetOwner(neta).Save(ctx)
//	if err != nil {
//		return err
//	}
//	_, err = client.Group.Create().SetName("Gitlab").AddUsers(neta, a8m).Save(ctx)
//	if err != nil {
//		return err
//	}
//	_, err = client.Group.Create().SetName("Github").AddUsers(a8m).Save(ctx)
//	if err != nil {
//		return err
//	}
//	log.Println("The graph was created successfully!")
//	return nil
//
//}
//
//// 获取名为GitHub 组内的所有汽车
//func QueryGithub(ctx context.Context, client *ent.Client) error {
//
//	cars, err := client.Group.Query().Where(group.Name("Github")).QueryUsers().QueryCars().All(ctx)
//	if err != nil {
//		return fmt.Errorf("failed getting cars: %v", err)
//	}
//	log.Println("cars returned:", cars)
//	return nil
//}
//
//// 便遍历的源是用户Ariel：
//func QueryArielCars(ctx context.Context, client *ent.Client) error {
//
//	a8m, err := client.User.Query().Where(user.HasCars(), user.Name("Ariel")).Only(ctx)
//
//	cars, err := a8m.QueryGroups().QueryUsers().QueryCars().Where(car.Not(car.ModelEQ("Mazda"))).All(ctx)
//
//	if err != nil {
//		return fmt.Errorf("failed getting cars: %v", err)
//	}
//	log.Println("cars returned:", cars)
//	return nil
//}
//
//func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
//	groups, err := client.Group.
//		Query().
//		Where(group.HasUsers()).
//		All(ctx)
//	if err != nil {
//		return fmt.Errorf("failed getting groups: %v", err)
//	}
//	log.Println("groups returned:", groups)
//	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
//	return nil
//}
