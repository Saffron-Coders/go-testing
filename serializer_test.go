package serializer_test

import (
	"strconv"
	"testing"
	"time"

	jsonserializer "github.com/davetweetlive/jsonserializer"
	"github.com/stretchr/testify/require"
)

var fileName = "tmp/user.json"

func TestCreateUsers(t *testing.T) {
	t.Parallel()
	users, err := jsonserializer.CreateUsers()
	require.NoError(t, err)
	require.Equal(t, len(users), 3)
	require.Equal(t, users[1].FirstName, "Irak")
}

func TestWriteStructToJSONFile(t *testing.T) {
	t.Parallel()
	users, err := jsonserializer.CreateUsers()
	require.NoError(t, err)

	err = jsonserializer.WriteStructToJSONFile(users, fileName)
	require.NoError(t, err)
}

func TestFromJSONFileToStruct(t *testing.T) {
	t.Parallel()
	users, err := jsonserializer.FromJSONFileToStruct(fileName)
	require.NoError(t, err)
	require.Equal(t, len(users), 3)
	require.Equal(t, users[2].FirstName, "Imaginery")
}

func TestCreateUsersSubTest(t *testing.T) {
	users, err := jsonserializer.CreateUsers()
	require.NoError(t, err)
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

	require.NotEqual(t, batch, nil)

	for i, tc := range batch {
		t.Run(strconv.Itoa(tc.batch), func(t *testing.T) {
			t.Parallel()
			if tc.batch < 1 || tc.batch > 3 {

				require.NotEqual(t, tc.users[i].FirstName, "")

				require.NotEqual(t, tc.users[i].LastName, "")

				require.NotEqual(t, tc.users[i].Email, "")
			}
			require.NotEqual(t, len(tc.users), 0)
		})
	}
}
