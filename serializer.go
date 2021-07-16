package serializer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func WriteStructToJSONFile(listofusers []*User, fileName string) error {
	data, err := json.MarshalIndent(listofusers, "", "    ")
	if err != nil {
		log.Println("couldn't marshal to json")
		return err
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		log.Println("couldn't write to a file")
		return nil
	}
	return nil
}

func FromJSONFileToStruct(fileName string) ([]*User, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("couldn't read the json file")
		return nil, err
	}

	var users []*User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUsers() ([]*User, error) {
	users := []*User{
		&User{
			Id:        1,
			FirstName: "Dave",
			LastName:  "Augustus",
			Email:     "dave@mail.com",
			CreatedAt: time.Now(),
		},
		&User{
			Id:        2,
			FirstName: "Irak",
			LastName:  "Rigia",
			Email:     "ir@mail.com",
			CreatedAt: time.Now(),
		},
		&User{
			Id:        3,
			FirstName: "Imaginery",
			LastName:  "User",
			Email:     "iu@gmail.com",
			CreatedAt: time.Now(),
		},
	}

	return users, nil
}
