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

func TestGetUsers(t *testing.T) {

	t.Parallel()
	h, cleaner := setUp(t)
	defer cleaner(t)

	users := []models.User{
		{
			Name:     "foo",
			Birthday: "2000-01-31",
			Email:    "foo@example.com",
			Password: "secret",
			Gender:   "MALE",
		},
		{
			Name:     "bar",
			Birthday: "2004-05-10",
			Email:    "bar@example.com",
			Password: "secret",
			Gender:   "FEMALE",
		},
	}

	// Insert the test users into the database
	for _, user := range users {
		if err := h.us.Create(&user); err != nil {
			t.Fatalf("Failed to insert user into database: %v", err)
		}
	}

	tests := []struct {
		title    string
		req      *pb.GetUsersRequest
		expected []models.User
		hasError bool
	}{
		{
			"get all users: success without pagination",
			&pb.GetUsersRequest{
				Page:  0,
				Limit: 0,
			},
			users,
			false,
		},
	}

	userServer := NewUserServer(h)

	for _, tt := range tests {
		c := context.Background()
		resp, err := userServer.GetUsers(c, tt.req)
		if (err != nil) != tt.hasError {
			t.Errorf("%s hasError %t, but got error: %v.", tt.title, tt.hasError, err)
			t.FailNow()
		}

		if !tt.hasError {
			receivedUsers := resp.GetUsers()
			if len(receivedUsers) != len(tt.expected) {
				t.Errorf("%s: expected %d users, got %d", tt.title, len(tt.expected), len(receivedUsers))
			}

			for i, receivedUser := range receivedUsers {
				expectedUser := tt.expected[i]
				if receivedUser.GetEmail() != expectedUser.Email {
					t.Errorf("%s: wrong Email for user %d, expected %q, got %q", tt.title, i, expectedUser.Email, receivedUser.GetEmail())
				}
			}
		}
	}

}

func TestGetUserById(t *testing.T) {
	t.Parallel()
	h, cleaner := setUp(t)
	defer cleaner(t)

	// Create a test user
	testUser := models.User{
		Name:     "test",
		Birthday: "1995-08-22",
		Email:    "test@example.com",
		Password: "secret",
		Gender:   "MALE",
	}

	// Insert the test user into the database
	if err := h.us.Create(&testUser); err != nil {
		t.Fatalf("Failed to insert user into database: %v", err)
	}

	tests := []struct {
		title    string
		req      *pb.GetUserId
		expected models.User
		hasError bool
	}{
		{
			"get existing user by ID: success",
			&pb.GetUserId{
				Id: uint32(testUser.ID),
			},
			testUser,
			false,
		},
		{
			"get non-existing user by ID",
			&pb.GetUserId{
				Id: 9999,
			},
			models.User{},
			true,
		},
	}

	userServer := NewUserServer(h)

	for _, tt := range tests {
		c := context.Background()
		resp, err := userServer.GetUserByID(c, tt.req)
		if (err != nil) != tt.hasError {
			t.Errorf("%s hasError %t, but got error: %v.", tt.title, tt.hasError, err)
			t.FailNow()
		}

		if !tt.hasError {
			receivedUser := resp.GetUser()
			if receivedUser.GetEmail() != tt.expected.Email {
				t.Errorf("%s: wrong Email, expected %q, got %q", tt.title, tt.expected.Email, receivedUser.GetEmail())
			}
		}
	}
}
