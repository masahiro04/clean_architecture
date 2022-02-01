package userDao

import (
	"clean_architecture/golang/domains"
	userModel "clean_architecture/golang/domains/user"
	"time"

	"database/sql"
	"log"
)

type rw struct {
	store *sql.DB
}

func New(db *sql.DB) *rw {
	return &rw{
		store: db,
	}
}

type UserDto struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rw rw) GetAll() (*domains.Users, error) {
	var users []domains.User
	rows, err := rw.store.Query(GetAllSql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userDto UserDto

		if err = rows.Scan(
			&userDto.ID,
			&userDto.Name,
			&userDto.Email,
			&userDto.CreatedAt,
			&userDto.UpdatedAt,
		); err != nil {
			log.Fatal(err)
			return nil, err
		}

		id, _ := userModel.NewId(userDto.ID)
		name, _ := userModel.NewName(userDto.Name)
		email, _ := userModel.NewEmail(userDto.Email)
		newUser := domains.BuildUser(id, name, email, userDto.CreatedAt, userDto.UpdatedAt)

		users = append(users, newUser)
	}
	usersData := domains.NewUsers(users)
	return &usersData, nil
}

func (rw rw) GetById(id int) (*domains.User, error) {
	var userDto UserDto

	result := rw.store.QueryRow(GetByIdSql, id)
	err := result.Scan(
		&userDto.ID,
		&userDto.Name,
		&userDto.Email,
		&userDto.CreatedAt,
		&userDto.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := userModel.NewId(userDto.ID)
	name, _ := userModel.NewName(userDto.Name)
	email, _ := userModel.NewEmail(userDto.Email)
	newUser := domains.BuildUser(_id, name, email, userDto.CreatedAt, userDto.UpdatedAt)
	return &newUser, nil
}

//
// func (rw rw) Create(newBlog domains.Blog) (*domains.Blog, error) {
// 	var id int
// 	err := rw.store.QueryRow(
// 		CreateSql,
// 		newBlog.Title.Value, newBlog.Body.Value, time.Now(), time.Now()).Scan(&id)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
//
// 	_id, _ := blogModel.NewId(id)
// 	title, _ := blogModel.NewTitle(newBlog.Title.Value)
// 	body, _ := blogModel.NewBody(newBlog.Body.Value)
// 	blog := domains.BuildBlog(_id, title, body, newBlog.CreatedAt, newBlog.UpdatedAt)
// 	return &blog, nil
// }
//
// func (rw rw) CreateTx(newBlog domains.Blog, tx *sql.Tx) (*domains.Blog, error) {
// 	var id int
// 	err := tx.QueryRow(
// 		CreateSql,
// 		newBlog.Title, newBlog.Body, time.Now(), time.Now()).Scan(&id)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
//
// 	_id, _ := blogModel.NewId(newBlog.ID.Value)
// 	title, _ := blogModel.NewTitle(newBlog.Title.Value)
// 	body, _ := blogModel.NewBody(newBlog.Body.Value)
// 	blog := domains.BuildBlog(_id, title, body, newBlog.CreatedAt, newBlog.UpdatedAt)
// 	return &blog, nil
// }
//
// //
// func (rw rw) Update(id int, blog domains.Blog) (*domains.Blog, error) {
// 	_, err := rw.store.Exec(
// 		UpdateSql,
// 		id, blog.Title.Value, blog.Body.Value, time.Now())
//
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
//
// 	_id, _ := blogModel.NewId(id)
// 	title, _ := blogModel.NewTitle(blog.Title.Value)
// 	body, _ := blogModel.NewBody(blog.Body.Value)
// 	newBlog := domains.BuildBlog(_id, title, body, blog.CreatedAt, blog.UpdatedAt)
// 	return &newBlog, nil
// }
//
// func (rw rw) Delete(id int) error {
// 	if _, err := rw.store.Exec(DeleteSql, id, time.Now(), time.Now()); err != nil {
// 		return err
// 	}
// 	return nil
// }
