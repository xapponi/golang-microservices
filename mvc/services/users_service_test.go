package services

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/xapponi/golang-microservices/mvc/domain"
	"github.com/xapponi/golang-microservices/mvc/utils"
)

// var (
// 	userDaoMock usersDaoMock
// )

// func init() {
// 	userDaoMock := usersDaoMock{}
// 	domain.UserDao = &userDaoMock
// }

type usersDaoMock struct {
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return m.getUserFunction(userId)
}
func Test_usersService_GetUser(t *testing.T) {

	type args struct {
		userId int64
	}
	tests := []struct {
		name  string
		f     func(userId int64) (*domain.User, *utils.ApplicationError)
		args  args
		want  *domain.User
		want1 *utils.ApplicationError
	}{
		// TODO: Add test cases.
		{
			"invalid userId",
			func(userId int64) (*domain.User, *utils.ApplicationError) {
				return nil, &utils.ApplicationError{
					Message: "user 0 was not found",
					Status:  http.StatusNotFound,
					Code:    "not_found",
				}
			},
			args{userId: 0},
			nil,
			&utils.ApplicationError{
				Message: "user 0 was not found",
				Status:  http.StatusNotFound,
				Code:    "not_found",
			},
		},
		{
			"test get user no error",
			func(userId int64) (*domain.User, *utils.ApplicationError) {
				return &domain.User{
					Id:        123,
					FirstName: "Xander",
					LastName:  "Apponi",
					Email:     "myemail@gmail.com",
				}, nil
			},
			args{userId: 123},
			&domain.User{
				Id:        123,
				FirstName: "Xander",
				LastName:  "Apponi",
				Email:     "myemail@gmail.com",
			},
			nil,
		},
	}

	userDaoMock := usersDaoMock{}
	domain.UserDao = &userDaoMock

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usersService{}
			userDaoMock.getUserFunction = tt.f
			got, got1 := u.GetUser(tt.args.userId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("usersService.GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
