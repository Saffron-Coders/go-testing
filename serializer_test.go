package serializer_test

import (
	"errors"
	"strconv"
	"testing"
	"time"

	jsonserializer "github.com/davetweetlive/jsonserializer"
)

var fileName = "tmp/user.json"

func TestCreateUsers(t *testing.T) {
	t.Parallel()
	users, err := jsonserializer.CreateUsers()
	if err != nil {
		t.Errorf("error shoube be nil, but got %v", err)
	}

	if len(users) != 3 {
		t.Errorf("users length should be 3, but got %v", len(users))
	}

	if users[1].FirstName != "Irak" {
		t.Errorf("the firstname should be Irak but got %v", users[1].FirstName)
	}
}

func TestWriteStructToJSONFile(t *testing.T) {
	t.Parallel()
	users, err := jsonserializer.CreateUsers()
	if err != nil {
		t.Errorf("error shoube be nil, but got %v", err)
	}

	err = jsonserializer.WriteStructToJSONFile(users, fileName)
	if err != nil {
		t.Errorf("Couldn't write to JSON file")
	}
}

func TestFromJSONFileToStruct(t *testing.T) {
	t.Parallel()
	users, err := jsonserializer.FromJSONFileToStruct(fileName)
	if err != nil {
		t.Errorf("error shoube be nil, but got %v", err)
	}
	if len(users) != 3 {
		t.Errorf("users length should be 3, but got %v", len(users))
	}

	if users[1].FirstName != "Irak" {
		t.Errorf("the firstname should be Irak but got %v", users[1].FirstName)
	}
}

func TestCreateUsersSubTest(t *testing.T) {
	users, err := jsonserializer.CreateUsers()
	if err != nil {
		t.Errorf("error shoube be nil, but got %v", err)
	}

	batch := []struct {
		batch int
		users []*jsonserializer.User
	}{
		{batch: 1, users: users},
		{batch: 2, users: []*jsonserializer.User{
			{
				Id:        1,
				FirstName: "ABC",
				LastName:  "Some Last Name",
				Email:     "abc@mail.com",
				CreatedAt: time.Now(),
			},
			{
				Id:        3,
				FirstName: "PQR",
				LastName:  "Alphabet",
				Email:     "pa@gmail.com",
				CreatedAt: time.Now(),
			},
		},
		},
		{
			batch: 3,
			users: []*jsonserializer.User{
				{
					Id:        3,
					FirstName: "XYZ",
					LastName:  "Alphabet",
				},
			},
		},
	}

	if batch == nil {
		t.Errorf("Errro %v", errors.New("batch shouldn't be empty"))
	}

	for i, tc := range batch {
		t.Run(strconv.Itoa(tc.batch), func(t *testing.T) {
			t.Parallel()
			if tc.batch < 1 || tc.batch > 3 {
				t.Errorf("Batch can't be more then three or less than equal to zeor, got %v", tc.batch)
				if tc.users[i].FirstName == "" {
					t.Errorf("firstname can't be empty")
				}
				if tc.users[i].LastName == "" {
					t.Errorf("lastname can't be empty")
				}
				if tc.users[i].Email == "" {
					t.Errorf("email can't be empty")
				}
			}

			if len(tc.users) == 0 {
				t.Error("user's list shouldn't be zero, got ", len(tc.users))
			}

		})
	}
}
