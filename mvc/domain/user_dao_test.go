package domain

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/xapponi/golang-microservices/mvc/utils"
)

// func TestGetUserNoUserFound( t *testing.T) {
// 	user, err := GetUser(0)
// 	if user != nil {
// 		t.Error("we were not expecting a user with id 0")
// 	}
// 	if err == nil {
// 		t.Error("we were expecting an error when user id is 0")
// 	}
// }

// func TestGetUserUserFound(t *testing.T) {
// 	user, err := GetUser(123)
// 	if user.Id != 123 {
// 		t.Error("expecting user with id 123")
// 	}
// 	if err != nil {
// 		t.Error("expecting no error")
// 	}
// }
func TestGetUser(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name  string
		u     *userDao
		args  args
		want  *User
		want1 *utils.ApplicationError
	}{
		{
			"test user not found",
			&userDao{},
			args{userId: 0},
			nil,
			&utils.ApplicationError{
				Message: fmt.Sprintf("user %v was not found", 0),
				Status:  http.StatusNotFound,
				Code:    "not_found",
			},
		},
		{
			"test user found",
			&userDao{},
			args{userId: 123},
			&User{
				Id:        123,
				FirstName: "Xander",
				LastName:  "Apponi",
				Email:     "myemail@gmail.com",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.u
			got, got1 := u.GetUser(tt.args.userId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
