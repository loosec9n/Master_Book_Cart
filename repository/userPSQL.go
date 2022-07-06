package repository

import (
	"Book_Cart_Project/models"
	"database/sql"
	"log"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//function to add user details to user database
func (r Repository) UserSignup(user models.User) models.User {

	query := `INSERT INTO users (
		first_name, 
		last_name, 
		password, 
		email, 
		phone_number) 
		VALUES($1, $2, $3, $4, $5) 
		RETURNING user_id;`

	//Makes query
	err := r.DB.QueryRow(query, user.First_Name, user.Last_Name, user.Password, user.Email, user.Phone_Number).Scan(&user.User_ID)
	logFatal(err)
	return user
}

// To check whether u.User exists
func (r Repository) DoesUserExists(user models.User) bool {

	query := `SELECT email FROM users 
				WHERE email = $11`

	err := r.DB.QueryRow(query, user.Email).Scan(&user.User_ID)
	// returns false if the user email does not exists
	return (err != sql.ErrNoRows)
}

func (r Repository) UserLogin(user models.User) (models.User, error) {

	query := `SELECT * FROM users 
				WHERE email = $1`

	err := r.DB.QueryRow(query, user.Email).Scan(&user.User_ID, &user.First_Name, &user.Last_Name,
		&user.Password, &user.Email, &user.Phone_Number, &user.IsAdmin, &user.CreatedAt)

	return user, err
}

func (r Repository) AdminLogin(admin models.User) (models.User, error) {

	query := `SELECT * FROM users 
				WHERE email = $1`
	var name sql.NullString //expecting the value being null
	err := r.DB.QueryRow(query, admin.Email).Scan(&admin.User_ID, &name, &admin.Last_Name,
		&admin.Password, &admin.Email, &admin.Phone_Number, &admin.IsAdmin, &admin.CreatedAt)

	if name.Valid { //giving a default value to the cell.
		admin.First_Name = name.String
	}
	return admin, err
}
