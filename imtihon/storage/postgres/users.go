package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	pakage "my_mod/Pakage"
	"my_mod/model"

	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewProRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (p *UserRepo) UserCreate(user model.Users) error {

	_, err := p.Db.Exec("insert into users(id,name,email,birthday,password,create_at,update_at) values($1,$2,$3,$4,$5,$6,$7)",
		uuid.NewString, user.Name, user.Email, user.Birthday, user.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (pa *UserRepo) UserRead(user model.Users) ([]model.Users, error) {
	rows, err := pa.Db.Query("select Id,Name, Email, Birthday, Password from users;")
	if err != nil {
		return nil, err
	}

	var p []model.Users
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}
		p = append(p, user)
	}
	return p, nil
}

func (u *UserRepo) UserUpdate(usersUpdateFilter model.UpdateUser) error {
	var params []string
	var args []interface{}
	query := `
	SELECT id
	FROM users
	WHERE delete_at IS NULL AND id = $1
	`

	if err := u.Db.QueryRow(query, usersUpdateFilter.UserId).Err(); err != nil {
		return fmt.Errorf("users by this id not found: %v", err)
	}

	query = `
	UPDATE users SET 
	`

	if usersUpdateFilter.Name != nil {
		params = append(params, fmt.Sprintf(" name = $%d ", len(args)+1))
		args = append(args, *usersUpdateFilter.Name)
	}

	if usersUpdateFilter.Email != nil {
		params = append(params, fmt.Sprintf(" email = $%d ", len(args)+1))
		args = append(args, *usersUpdateFilter.Email)
	}

	if usersUpdateFilter.Birthday != nil {
		params = append(params, fmt.Sprintf("birthday = $%d", len(args)+1))
		args = append(args, *usersUpdateFilter.Birthday)
	}

	if usersUpdateFilter.Password != nil {
		params = append(params, fmt.Sprintf("password = $%d", len(args)+1))
		args = append(args, *usersUpdateFilter.Password)
	}

	params = append(params, fmt.Sprintf("update_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}
	fmt.Println("salom")
	fmt.Println(len(args))

	args = append(args, usersUpdateFilter.UserId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND delete_at IS NULL", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err := u.Db.Exec(query, args...)

	fmt.Println(query)
	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

func (u *UserRepo) UserDelete(id string) error {
	_rows, err := u.Db.Exec(`update users set
	delete_at = date_part('epoch', current_timestamp)::INT
   where id = $1 and delete_at = 0`, id)
	if err != nil {
		return err
	}
	rowsaff,err:=_rows.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsaff==0 {
		return err
		
	}

	return nil
}

func (u *UserRepo) GetAllUsers(f model.UserGetAll) ([]model.Users ,error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)
	fmt.Println("salom")

	query := `SELECT id, name,email, birthday,password
	FROM users WHERE true `

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if f.Birthday !=nil{
		params["bi"] = f.Birthday
		filter += " and birthday = :birthday "
	}


	
	if f.Offset > 0 {
		params["offset"] = f.Offset
		filter += " OFFSET :offset"
	}
	
	if f.Limit > 0 {
		params["limit"] = f.Limit
		filter += " LIMIT :limit"
	}
	query = query + filter

	query, arr = pakage.ReplaceQueryParams(query, params)
	fmt.Println(query, arr)
	rows, err := u.Db.Query(query, arr...)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.Id, &user.Name, &user.Email,&user.Birthday,&user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
