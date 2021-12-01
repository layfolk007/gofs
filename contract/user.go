package contract

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/no-src/gofs/util"
	"strings"
)

type User struct {
	userId   int
	userName string
	password string
}

func (user *User) String() string {
	return fmt.Sprintf("%s|%s", user.userName, user.password)
}

func (user *User) UserId() int {
	return user.userId
}

func (user *User) UserName() string {
	return user.userName
}

func (user *User) Password() string {
	return user.password
}

func NewUser(userId int, userName string, password string) (*User, error) {
	if userId <= 0 {
		return nil, errors.New("userId must greater than zero")
	}
	user := &User{
		userId:   userId,
		userName: userName,
		password: password,
	}
	err := isValidUser(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func isValidUser(user User) error {
	if len(user.UserName()) == 0 {
		return errors.New("userName can't be empty")
	}
	if len(user.Password()) == 0 {
		return errors.New("password can't be empty")
	}
	if strings.ContainsAny(user.UserName(), ",|") {
		return errors.New("userName can't contain ',' or '|' ")
	}
	if strings.ContainsAny(user.Password(), ",|") {
		return errors.New("password can't contain ',' or '|' ")
	}
	return nil
}

// ParseUsers parse users string to User List
// For example: user1|password1,user2|password2
func ParseUsers(userStr string) (users []*User, err error) {
	if len(userStr) == 0 {
		return users, nil
	}
	all := strings.Split(userStr, ",")
	userCount := 0
	for _, userStr := range all {
		userInfo := strings.Split(userStr, "|")
		if len(userInfo) == 2 {
			userName := strings.TrimSpace(userInfo[0])
			password := strings.TrimSpace(userInfo[1])
			if len(userName) > 0 && len(password) > 0 {
				userCount++
				user, err := NewUser(userCount, userName, password)
				if err != nil {
					return nil, err
				}
				users = append(users, user)
			}
		} else {
			return nil, fmt.Errorf("invalid user info => [%s]", userStr)
		}
	}
	return users, nil
}

func RandomUser(count, userLen, pwdLen int) []*User {
	var users []*User
	for i := 1; i <= count; i++ {
		user, err := NewUser(i, util.RandomString(userLen), util.RandomString(pwdLen))
		if err != nil {
			i--
		} else {
			users = append(users, user)
		}
	}
	return users
}

func ParseStringUsers(users []*User) (userStr string, err error) {
	if len(users) == 0 {
		return userStr, nil
	}
	var userResultList []string
	for _, user := range users {
		if user == nil {
			continue
		}
		err = isValidUser(*user)
		if err != nil {
			return userStr, err
		}
		userResultList = append(userResultList, user.String())
	}
	userStr = strings.Join(userResultList, ",")
	return userStr, nil
}

type SessionUser struct {
	UserId   int
	UserName string
	Password string
}

func MapperToSessionUser(user *User) *SessionUser {
	if user == nil {
		return nil
	}
	return &SessionUser{
		UserId:   user.UserId(),
		UserName: user.UserName(),
		Password: user.Password(),
	}
}

func init() {
	gob.Register(SessionUser{})
}
