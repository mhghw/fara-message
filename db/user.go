package db

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

func init() {
	UsersDB.users = sampleUsers()
}

var UsersDB userDB

type userDB struct {
	users map[string]User
	mu    sync.Mutex
}

type Gender int8

const (
	Male Gender = iota
	Female
)

type User struct {
	ID          string
	Username    string
	FirstName   string
	LastName    string
	Password    string
	Gender      Gender
	DateOfBirth time.Time
	CreatedTime time.Time
}

func (db *userDB) CreateUser(user User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, has := db.users[user.ID]; has {
		return ErrUserAlreadyExists
	}

	db.users[user.ID] = user

	return nil
}

func (db *userDB) GetUser(userID string) (User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	user, has := db.users[userID]
	if !has {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (db *userDB) UpdateUser(newUser User) error {
	if _, err := db.GetUser(newUser.ID); err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[newUser.ID] = newUser

	return nil
}

func sampleUsers() map[string]User {
	return map[string]User{
		"10001": {
			ID:          "10001",
			Username:    "hamid123",
			FirstName:   "Hamid",
			LastName:    "Javidi",
			Password:    "554dbf0b41b3cd068ee1fcfd6235466a263647b4", //qwerqwer
			Gender:      Male,
			DateOfBirth: time.Date(1997, time.January, 12, 0, 0, 0, 0, time.UTC),
			CreatedTime: time.Now(),
		},
		"10002": {
			ID:          "10002",
			Username:    "rezareza",
			FirstName:   "Reza",
			LastName:    "Gholami",
			Password:    "b6d2b6e8aad5b72946292dea96b5af2c6d3e94ab", //rezarezagholami
			Gender:      Male,
			DateOfBirth: time.Date(1978, time.May, 2, 0, 0, 0, 0, time.UTC),
			CreatedTime: time.Now(),
		},
		"10003": {
			ID:          "10003",
			Username:    "n.Panahi",
			FirstName:   "Neda",
			LastName:    "Panahi",
			Password:    "37ed856a34c507996c81a7e5b702dfc11a7bd416", //nedapanahi1234
			Gender:      Female,
			DateOfBirth: time.Date(2005, time.August, 19, 0, 0, 0, 0, time.UTC),
			CreatedTime: time.Now(),
		},
		"10004": {
			ID:          "10004",
			Username:    "hesami",
			FirstName:   "Sara",
			LastName:    "Hesami",
			Password:    "4caab017bc944efa3ff354d44dc44c29a85f78fd", //sarahfiwowoie
			Gender:      Female,
			DateOfBirth: time.Date(1978, time.November, 1, 0, 0, 0, 0, time.UTC),
			CreatedTime: time.Now(),
		},
	}
}
