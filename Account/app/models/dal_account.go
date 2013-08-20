package models

import (
	"labix.org/v2/mgo/bson"
	"errors"
	"fmt"
)

func (d *Dal) SaveUser(user *User) error {
	uc := d.session.DB(DbName).C(UserCollection)

	i, _ := uc.Find(bson.M{"username":user.UserName}).Count()
	if i != 0 {
		fmt.Println("username registed!!!!!!!!!!!!!")
		return errors.New("user name registed!!!")
	}

	i, _ = uc.Find(bson.M{"email":user.Email}).Count()
	if i != 0 {
		fmt.Println("email registed!!!!!!!!!!!!!")
		return errors.New("email name registed!!!")
	}

	err := uc.Insert(user)
	return err
}
