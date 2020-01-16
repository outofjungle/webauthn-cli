package main

import (
	"fmt"
	"sync"
)

type userdb struct {
	users map[string]*User
	mu    sync.RWMutex
}

var db *userdb

func DB() *userdb {

	if db == nil {
		db = &userdb{
			users: make(map[string]*User),
		}
	}

	return db
}

func (db *userdb) GetUser(name string) (*User, error) {

	db.mu.Lock()
	defer db.mu.Unlock()
	user, ok := db.users[name]
	if !ok {
		return &User{}, fmt.Errorf("error getting user '%s': does not exist", name)
	}

	return user, nil
}

func (db *userdb) PutUser(user *User) {

	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.name] = user
}

func (db *userdb) DeleteUser(name string) {

	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.users, name)
}

func (db *userdb) ListUsers() []string {
	users := make([]string, 0, len(db.users))
	for u := range db.users {
		users = append(users, u)
	}
	return users
}
