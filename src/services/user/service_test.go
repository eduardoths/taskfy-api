package user_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserPost struct {
	Name            string
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

type UserService struct{}

func (this *UserService) Signup(user *UserPost) string {
	if user.Password == user.ConfirmPassword {
		return "valid_token"
	}
	return ""
}

func TestUserSignup(t *testing.T) {
	cases := []struct {
		description string
		input       UserPost
		want        string
	}{
		{
			description: "should return token when password matches confirmation",
			input: UserPost{
				Name:            "Valid name",
				Username:        "valid_username",
				Email:           "valid_email@mail.com",
				Password:        "valid_password",
				ConfirmPassword: "valid_password",
			},
			want: "valid_token",
		},
		{
			description: "should not return token when passwords don't match",
			input: UserPost{
				Name:            "Valid name",
				Username:        "valid_username",
				Email:           "valid_email@mail.com",
				Password:        "valid_password",
				ConfirmPassword: "invalid_password",
			},
			want: "",
		},
	}

	for i, scenario := range cases {
		t.Run(fmt.Sprintf("Case %d: %s", i, scenario.description), func(t *testing.T) {
			serv := UserService{}
			actual := serv.Signup(&scenario.input)
			assert.Equal(t, scenario.want, actual)
		})
	}
}
