package repository

import (
	"Book_Cart_Project/models"
	"database/sql"
	"log"
)

//function to add user details to user database
func (r Repository) UserSignup(user models.User) models.User {

	query := `INSERT INTO users (
				first_name, 
				last_name, 
				password, 
				email, 
				phone_number) 
				VALUES($1, $2, $3, $4, $5) 
				RETURNING 
				user_id,
				is_active;`

	//Makes query
	err := r.DB.QueryRow(query,
		user.First_Name,
		user.Last_Name,
		user.Password,
		user.Email,
		user.Phone_Number).Scan(
		&user.User_ID,
		&user.Is_Active)
	//logFatal(err)
	if err != nil {
		log.Println("User signup error", err)
	}
	return user
}

// To check whether u.User exists
func (r Repository) DoesUserExists(user models.User) bool {

	query := `SELECT 
		email 
		FROM 
		users 
		WHERE 
		email = $1`

	err := r.DB.QueryRow(query,
		user.Email).Scan(
		&user.User_ID)
	// returns false if the user email does not exists
	return (err != sql.ErrNoRows)
}

func (r Repository) UserLogin(user models.User) (models.User, error) {

	query := `SELECT 
		user_id,
		first_name,
		last_name,
		password,
		email,
		phone_number,
		is_active, 
		created_at
		FROM 
		users 
		WHERE email = $1`

	err := r.DB.QueryRow(query, user.Email).Scan(
		&user.User_ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Password,
		&user.Email,
		&user.Phone_Number,
		&user.Is_Active,
		&user.CreatedAt)

	return user, err
}

func (r Repository) AdminLogin(admin models.User) (models.User, error) {

	query := `SELECT 
	user_id,
	first_name,
	last_name,
	password,
	email,
	phone_number,
	is_active, 
	created_at
	FROM 
	users 
	WHERE email = $1`

	//var name sql.NullString //expecting the value being null

	err := r.DB.QueryRow(query,
		admin.Email).Scan(
		&admin.User_ID,
		&admin.First_Name, //&name, //example check for null value
		&admin.Last_Name,
		&admin.Password,
		&admin.Email,
		&admin.Phone_Number,
		&admin.IsAdmin,
		&admin.CreatedAt)

	return admin, err
}

func (r Repository) BlockUser(user models.User) (models.User, error) {
	query := `UPDATE users 
				SET is_active =$1
				WHERE email = $2
				RETURNING 
				user_id,
				first_name, 
				last_name, 
				email,
				phone_number,
				is_active, 
				created_at;`

	err := r.DB.QueryRow(query,
		user.Is_Active,
		user.Email).Scan(
		&user.User_ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Phone_Number,
		&user.Is_Active,
		&user.CreatedAt)
	return user, err
}

//Admin can View all the users
func (r Repository) ViewUser() ([]models.User, error) {

	var users []models.User

	query := `SELECT 
		user_id,
		first_name, 
		last_name, 
		email, 
		phone_number, 
		is_active, 
		is_admin,	
		created_at		
		FROM 
		users;`
	row, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	//Loopingg through the rows
	for row.Next() {
		var user models.User
		if err := row.Scan(
			&user.User_ID,
			&user.First_Name,
			&user.Last_Name,
			&user.Email,
			&user.Phone_Number,
			&user.Is_Active,
			&user.IsAdmin); err != nil {
			return users, err
		}
		if !user.IsAdmin {
			users = append(users, user)
		}

	}

	if err = row.Err(); err != nil {
		return users, err
	}
	return users, nil
}
