package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	// "github.com/labstack/echo/middleware"
)

type User struct {
	ID    int    `json:"id" xml:"id" form:"id" query:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// Middleware
func createCustomMiddleware(name string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer fmt.Printf("middleware-%s: defer\n", name)
			fmt.Printf("middleware-%s: before\n", name)
			err := next(c)
			fmt.Printf("middleware-%s: after\n", name)
			return err
		}
	}
}

// APIのCRUD

// GET
func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Users")
}

// GET id
func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, &User{
		ID:    id,
		Name:  "I am " + strconv.Itoa(id),
		Email: strconv.Itoa(id) + "@localhost.com",
	})
}

// POST
func addUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.JSON(http.StatusOK, &User{
		ID:    1,
		Name:  name,
		Email: email,
	})
}

// PUT
func editUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// DELETE
func deleteUser(c echo.Context) error {
	return nil
}

// セッションに値を保持する。
func login(e *echo.Echo) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			//Path:でsessionの有効な範囲を指定｡指定無しで全て有効になる｡
			Path: "/",
			//有効な時間
			MaxAge: 86400 * 7,
			//trueでjsからのアクセス拒否
			HttpOnly: true,
		}
		//テキトウな値
		sess.Values["foo"] = "bar"
		//ログインしました
		sess.Values["auth"] = true
		//状態保存
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.String(http.StatusOK, "logged in")
	}
}

// セッションのAuthをfalseにし、認証していないことを保持する。
func logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		//ログアウト
		sess.Values["auth"] = false
		//状態を保存
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.String(http.StatusOK, "logged out")
	}
}

// セッションの情報を取得するサンプル。
// 認証されていなければ401を返す。
func secret() echo.HandlerFunc {
	return func(c echo.Context) error {
		//sessionを見る
		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error")
		}
		//ログインしているか
		if b, _ := sess.Values["auth"]; b != true {
			return c.String(http.StatusUnauthorized, "401")
		} else {
			return c.String(http.StatusOK, sess.Values["foo"].(string))
		}
	}
}

func main() {
	e := echo.New()

	// e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		fmt.Printf("Path = %s\n", c.Path())
	// 		skipPaths := []string{"/", "/users", "login", "logout", "secret"}
	// 		for _, s := range skipPaths {
	// 			if c.Path() == s {
	// 				return true
	// 			}
	// 		}
	// 		return false
	// 	},
	// 	KeyLookup: "header:X-WLB-Authorization",
	// 	Validator: func(key string, c echo.Context) (bool, error) {
	// 		fmt.Printf("Auth key = %s\n", key)
	// 		return true, nil
	// 	},
	// }))

	e.Use(createCustomMiddleware("Use-1"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// for CRUD
	e.GET("/users", getUsers, createCustomMiddleware("Route"))
	e.GET("/users/:id", getUser)
	e.POST("/users", addUser)
	e.PUT("/users/:id", editUser)

	// for Session
	e.Use(session.Middleware(sessions.NewFilesystemStore("", []byte("secret_session"))))
	e.GET("/login", login(e))
	e.GET("logout", logout())
	e.GET("secret", secret())

	e.Logger.Fatal(e.Start(":1323"))
}
