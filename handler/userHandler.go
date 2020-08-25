package handler

/**
UserHandler
*/

// 创建新用户
//func createUser() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		//var user  ent.User
//		m := echo.Map{}
//		if err := c.Bind(&m); err != nil {
//			return err
//		}
//		u := ent.User{
//			Name: m["name"].(string),
//			Age: m["age"].(int),
//			Username: m["username"].(string),
//			Email: m["email"].(string),
//			Password : m["password"].(string),
//
//		}
//
//		//user = u
//		//_, err := ent.Client.User.Create()
//
//		return nil
//	}
//}

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
