package db

func CreateUser(user User) error {
	result := DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ReadAnotherUser(username string) (User, error) {
	var user User
	result := DB.First(&user, "username=?", username)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func ReadUser(ID string) (User, error) {
	var user User
	result := DB.First(&user, "ID=?", ID)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func UpdateUser(ID string, newInfo User) error {
	result := DB.Model(&User{}).Where("ID=?", ID).Updates(User{FirstName: newInfo.FirstName, LastName: newInfo.LastName, Gender: newInfo.Gender, DateOfBirth: newInfo.DateOfBirth})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(ID string) error {
	var user User
	result := DB.First(&user, "ID=?", ID)
	if result.Error != nil {
		return result.Error
	}
	result = DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
