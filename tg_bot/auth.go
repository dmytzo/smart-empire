package tg_bot

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var auth = newAuth()

type User struct {
	Username string `yaml:"username"`
	ChatId   int64  `yaml:"chat_id"`
	Active   bool   `yaml:"active"`
}

func (u *User) updateStatus(isActive bool) {
	u.Active = isActive
}

type Auth struct {
	Users    []*User `yaml:"users"`
	ApiToken string  `yaml:"api_token"`
}

func (a *Auth) getApiToken() string {
	return a.ApiToken
}

func (a *Auth) getUserByUsername(username string) (*User, error) {
	for _, user := range a.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("access denied")
}

func (a *Auth) updateUserStatus(username string, isActive bool) {
	user, _ := a.getUserByUsername(username)
	if user.Active == isActive {
		return
	}
	user.updateStatus(isActive)
	updateAuth(a)
}

func (a *Auth) getActiveUsers() []*User {
	var activeUsers []*User
	for _, user := range a.Users {
		if user.Active {
			activeUsers = append(activeUsers, user)
		}
	}
	return activeUsers
}

func GetAuth() *Auth {
	return auth
}

func newAuth() *Auth {
	newAuthType := &Auth{}
	file, err := os.Open("configs/tg.yaml")
	if err != nil {
		log.Fatal("Auth file was not found")
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err = d.Decode(&newAuthType); err != nil {
		log.Fatal("Can not decode auth file")
	}
	return newAuthType
}

func updateAuth(auth *Auth) {
	file, err := os.OpenFile("configs/tg.yaml", os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal("Auth file was not found")
	}
	defer file.Close()
	b, _ := yaml.Marshal(&auth)
	file.Write(b)
	auth = newAuth()
}
