package services

import (
	"context"
	"testing"

	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/AhmedEnnaime/SnapEvent/pb"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	h, cleaner := setUp(t)
	defer cleaner(t)

	fooUser := models.User{
		Name:     "foo",
		Birthday: "2000-01-31",
		Email:    "foo@example.com",
		Password: "secret",
		Gender:   "MALE",
	}

	barUser := models.User{
		Name:     "bar",
		Birthday: "2004-05-10",
		Email:    "bar@example.com",
		Password: "secret",
		Gender:   "FEMALE",
	}

	tests := []struct {
		title    string
		req      *pb.CreateUserRequest
		expected *models.User
		hasError bool
	}{
		{
			"create fooUser: success",
			&pb.CreateUserRequest{
				Name:     "foo",
				Birthday: "2000-01-31",
				Email:    "foo@example.com",
				Password: "secret",
				Gender:   "MALE",
			},
			&fooUser,
			false,
		},
		{
			"create barUser: success",
			&pb.CreateUserRequest{
				Name:     "bar",
				Birthday: "2004-05-10",
				Email:    "bar@example.com",
				Password: "secret",
				Gender:   "FEMALE",
			},
			&barUser,
			false,
		},
		{
			"create fooUser: email empty",
			&pb.CreateUserRequest{
				Name:     "foo",
				Birthday: "2000-01-31",
				Email:    "",
				Password: "secret",
				Gender:   "MALE",
			},
			nil,
			true,
		},
		{
			"create fooUser: email already exists",
			&pb.CreateUserRequest{
				Name:     "jack",
				Birthday: "2000-01-31",
				Email:    "foo@example.com",
				Password: "secret",
				Gender:   "FEMALE",
			},
			nil,
			true,
		},
		{
			"create fooUser: invalid gender",
			&pb.CreateUserRequest{
				Name:     "foo",
				Birthday: "2000-01-31",
				Email:    "foo@example.com",
				Password: "secret",
				Gender:   "BISEXUAL",
			},
			nil,
			true,
		},
	}

	userServer := NewUserServer(h)

	for _, tt := range tests {
		c := context.Background()
		resp, err := userServer.CreateUser(c, tt.req)
		if (err != nil) != tt.hasError {
			t.Errorf("%s hasError %t, but got error: %v.", tt.title, tt.hasError, err)
			t.FailNow()
		}

		if !tt.hasError {
			if resp.User.Email != tt.expected.Email {
				t.Errorf("%q wrong Email, expected %q, got %q", tt.title, tt.expected.Email, resp.User.Email)
			}
		}
	}

}
