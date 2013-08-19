package models

import (
	"fmt"
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/robfig/revel"
)

const (
	DEBUG_PWD = true
)

/*
 * real struct which was persisted in database
 */
type User struct {
	UserName string
	Email    string
	NickName string
	HashPassword []byte
}

func (user *User) String() string {
	if !DEBUG_PWD {
		return fmt.Sprintf("User(username = %s, email = %s, nick name = %s)",
			user.UserName, user.Email, user.NickName)
	} else {
		return fmt.Sprintf("User(username = %s, email = %s, nick name = %s), pwd = %s",
			user.UserName, user.Email, user.NickName, user.HashPassword)
	}
}

/*
* TODO: This should be a method or a function???
 */
func (user *User) generatePwdByte(pwdStr string) error {
	user.HashPassword, _ = bcrypt.GenerateFromPassword([]byte(pwdStr), bcrypt.DefaultCost)
	return nil
}

func (user *User) SaveUser(regUser *RegUser) error {
	user.generatePwdByte(regUser.PasswordStr)
	fmt.Println("Save User success: ", user)
	return nil
}

/*
 * used for login
 */
type LoginUser struct {
	User
	PasswordStr string
}

func (loginUser *LoginUser) Validate(v *revel.Validation) {
	v.Check(loginUser.UserName,
		revel.Required{},
		revel.MinSize{6},
		revel.MaxSize{16},
	)

	v.Check(loginUser.PasswordStr,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{16},
	)

	//0: generate passing str
	//1: get pwd bytes from database
	//2: compare them
	//test here
	pwd := "testtest"
	//rPwd := "testtest"
	rPwd := "testtest"
	v.Required(pwd == rPwd).Message("user name or password is wrong!!!")
}

/*
 * used for register or update user
 */
type RegUser struct {
	User
	PasswordStr string
	ConfirmPwdStr string
}

func (regUser *RegUser) Validate(v *revel.Validation) {
	//Check workflow:
	//see @validation.go Check(obj interface{}, checks ...Validator)
	//Validator is an interface, v.Check invoke v.Apply for each validator.
	//Further, v.Apply invoke validator.IsSatisfied with passing obj.
	//Checking result is an object of ValidationResult. The field Ok of ValidationResult
	//would be true if checking success. Otherwise, Ok would be false, and another filed
	//Error of ValidationResult would be non-nil, an ValidationError filled with error message
	//should be assigned to Error.
	v.Check(regUser.UserName,
		revel.Required{},
		revel.MinSize{6},
		revel.MaxSize{16},
	)

	v.Check(regUser.NickName,
		revel.Required{},
		revel.MinSize{6},
		revel.MaxSize{16},
	)

	//validation provide an convenient method for checking Email.
	//revel has a const for email rexgep, Email will use the rex to check email string.
	v.Email(regUser.Email)

	v.Check(regUser.PasswordStr,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{16},
	)
	v.Check(regUser.ConfirmPwdStr,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{16},
	)
	//pwd and comfirm_pwd should be equal
	v.Required(regUser.PasswordStr == regUser.ConfirmPwdStr).Message("The passwords do not match.")
}
