package services

import (
	"log"

	"github.com/dizars1776/library-lite/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	RoleId    int
	Email     string
	Name      string
	Surname   string
	Password  string
	CreatedAt string
}

type UserService struct {
	store *store.Store
}

func NewUserService(store *store.Store) *UserService {
	return &UserService{store: store}
}

func (u *UserService) getUser(email string) (*User, error) {
	stmt, err := u.store.DB().Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var user User

	err = stmt.QueryRow(email).Scan(&user.ID, &user.RoleId, &user.Email, &user.Name, &user.Surname, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &User{
		ID:        user.ID,
		RoleId:    user.RoleId,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *UserService) LoginUser(email string, password string) (*User, error) {
	dbUser, err := u.getUser(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
		return nil, err
	}

	return dbUser, nil
}

// func (u *UserService) LogoutUser(userId int) error {
// 	//TODO: save the deleted token? Maybe? Probably to be removed.
// 	return nil
// }

func (u *UserService) GetRoleName(roleId int) (roleName string, err error) {
	stmt, err := u.store.DB().Prepare("SELECT name FROM roles WHERE id = ?")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer stmt.Close()

	err = stmt.QueryRow(roleId).Scan(&roleName)

	return roleName, nil
}
