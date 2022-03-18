package option_test

import (
	"encoding/json"
	"fmt"

	"github.com/cyrusaf/option"
)

func ExampleOption() {
	type User struct {
		ID   string
		Name string
	}
	var GetUser = func(id string) option.Option[User] {
		if id == "1" {
			return option.Some(User{
				ID:   "1",
				Name: "cyrusaf",
			})
		}
		return option.None[User]()
	}

	type HTTPResponse struct {
		StatusCode int
		Body       string
	}
	var HandleGetUser = func(id string) HTTPResponse {
		userResult := GetUser(id)
		user, ok := userResult.Unwrap()
		if !ok {
			return HTTPResponse{
				StatusCode: 404,
				Body:       "user not found",
			}
		}
		b, err := json.Marshal(user)
		if err != nil {
			return HTTPResponse{
				StatusCode: 500,
				Body:       fmt.Sprintf("marshalling user: %s", err.Error()),
			}
		}

		return HTTPResponse{
			StatusCode: 200,
			Body:       string(b),
		}

	}

	fmt.Printf("%+v\n", HandleGetUser("1"))
	fmt.Printf("%+v\n", HandleGetUser("2"))

	// Output:
	// {StatusCode:200 Body:{"ID":"1","Name":"cyrusaf"}}
	// {StatusCode:404 Body:user not found}
}
