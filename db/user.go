package db

import (
	"fmt"
)

func (d *Database) CreateUser(user User) error {
	userTable := ConvertUserToUserTable(user)
	result := d.db.Create(&userTable)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}

func (d *Database) ReadUserByUsername(username string) (UserTable, error) {
	var user UserTable
	result := d.db.First(&user, "username=?", username)
	if result.Error != nil {
		return user, fmt.Errorf("failed to read user: %w", result.Error)
	}
	return user, nil
}

func (d *Database) ReadUser(ID int) (UserTable, error) {
	var user UserTable
	result := d.db.First(&user, "ID=?", ID)
	if result.Error != nil {
		return user, fmt.Errorf("failed to read user: %w", result.Error)
	}
	return user, nil
}

func (d *Database) UpdateUser(ID int, newInfo UserTable) error {

	result := d.db.Model(&UserTable{}).Where("ID=?", ID).Updates(UserTable{Username: newInfo.Username, FirstName: newInfo.FirstName, LastName: newInfo.LastName, Password: newInfo.Password, Gender: newInfo.Gender, DateOfBirth: newInfo.DateOfBirth})
	if result.Error != nil {
		return fmt.Errorf("failed to Update user: %w", result.Error)
	}
	return nil
}

func (d *Database) DeleteUser(ID int) error {
	var user UserTable
	result := d.db.First(&user, "ID=?", ID)
	if result.Error != nil {
		return fmt.Errorf("failed to find user to delete: %w", result.Error)
	}
	result = d.db.Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	return nil
}
