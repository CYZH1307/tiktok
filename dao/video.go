package dao

type User struct {
	Id 			int64
	Name 		string
	Password	string
}

func GetUserById(id int64) (User, error) {
	user := User{}
	if id == 0 {
		return user, nil
	}
	err := DB.Where("id = ?", id).First(&user).Error
	Handle(err)
	
	return user, err
}

func GetUserByName(name string) (User, error) {
	user := User{}
	err := DB.Where("name = ?", name).First(&user).Error
	return user, err
}

func InsertUser(user *User) error {
	err := DB.Create(user).Error
	Handle(err)
	return err
}

