package Models

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserRepository interface {
	Create(User) (bool, error)
	GetByEmail(email string) (User, error)
	GetAll() (UsersFile, error)
	Update(newUser User) (User, error)
	Delete(id string) (bool, error)
}

var localJsonFile = "./Data/users.json"

func InitialiseUsersFile() (UsersFile, error) {
	var usersFile UsersFile
	//we check if the file is created or not
	if _, err := os.Stat(localJsonFile); os.IsNotExist(err) {
		usersFile = make(UsersFile, 0)
	} else {
		//if the file exists, read its contents
		usersData, err := os.ReadFile(localJsonFile)
		if err != nil {
			return UsersFile{}, err
		}
		err = json.Unmarshal(usersData, &usersFile)
		if err != nil {
			return UsersFile{}, err
		}
	}
	return usersFile, nil
}
func SaveUsersFile(usersFile UsersFile) error {
	//convert the user data back to JSON
	jsonData, err := json.MarshalIndent(usersFile, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(localJsonFile, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

func Create(newUser User) (bool, error) {
	usersFile, err := InitialiseUsersFile()
	if err != nil {
		return false, err
	}
	//check the value of email in the local json whether it exists or not to avoid
	//double entry
	for _, user := range usersFile {
		if user.Email == newUser.Email {
			return true, nil
		}
	}
	//if user does not exist add them
	usersFile = append(usersFile, newUser)
	err = SaveUsersFile(usersFile)
	if err != nil {
		return false, err
	}
	return false, nil
}

func GetByEmail(email string) (User, error) {
	usersFile, err := InitialiseUsersFile()
	//check if users file is empty
	if len(usersFile) == 0 {
		return User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return User{}, err
	}
	for _, user := range usersFile {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func GetAll() (UsersFile, error) {
	usersFile, err := InitialiseUsersFile()
	if err != nil {
		return UsersFile{}, err
	}
	return usersFile, nil
}

func Update(newUser User) (User, error) {
	usersFile, err := InitialiseUsersFile()
	if err != nil {
		return User{}, err
	}

	// Check if users file is empty
	if len(usersFile) == 0 {
		return User{}, fmt.Errorf("user not found")
	}

	// Look for the specific user and update their information
	userFound := false
	for i, user := range usersFile {
		if user.Id == newUser.Id {
			usersFile[i] = newUser
			userFound = true
			break
		}
	}

	// If user not found, return an error
	if !userFound {
		return User{}, fmt.Errorf("user not found")
	}

	// Save the updated users file
	err = SaveUsersFile(usersFile)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func Delete(id string) (bool, error) {
	usersFile, err := InitialiseUsersFile()
	if err != nil {
		return false, err
	}
	//check if users file is empty
	if len(usersFile) == 0 {
		return false, fmt.Errorf("user not found")
	}
	newUsersFile := (usersFile)[:0]
	for _, user := range usersFile {
		if user.Id != id {
			newUsersFile = append(newUsersFile, user)
		}
	}
	err = SaveUsersFile(newUsersFile)
	if err != nil {
		return false, err
	}
	return true, nil

}
