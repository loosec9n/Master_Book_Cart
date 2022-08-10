package repository

import (
	"Book_Cart_Project/models"
	"context"
	"database/sql"
	"log"
)

type wishList struct {
	Id       int     `json:"wishID"`
	Price    float64 `json:"wishPrice"`
	PrImage  string  `json:"wishImage"`
	PrName   string  `json:"wishName"`
	PrUserID int     `json:"wishUserID"`
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

func (r Repository) AddAddress(address models.Address) (models.Address, error) {

	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("error begining transaction")
		return models.Address{}, err
	}

	query := `INSERT INTO user_address(
		address_id,
		house_name,
		street_name,
		land_mark,
		city,
		add_state,
		pincode)
		VALUES($1,$2,$3,$4,$5,$6,$7)
		RETURNING 
		address_id,
		house_name, 
		city;`

	err = tx.QueryRow(query,
		address.AddressID,
		address.HouseName,
		address.StreetName,
		address.LandMark,
		address.City,
		address.State,
		address.Pincode).Scan(
		&address.UserAddress_ID,
		&address.HouseName,
		&address.City,
	)
	if err != nil {
		log.Println("error running the query", err)
		tx.Rollback()
		return models.Address{}, err
	}

	_, err = tx.Exec(`UPDATE 
						users 
					SET 
						user_address_id = $1 
					WHERE 
						user_id = $2;`, address.AddressID, address.UserAddress_ID)
	if err != nil {
		log.Println("error running update query", err)
		tx.Rollback()
		return models.Address{}, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error in commiting add address")
		return models.Address{}, err
	}

	return address, nil
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
		is_active 
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
	)

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
	is_active
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
	)

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
				is_active;`

	err := r.DB.QueryRow(query,
		user.Is_Active,
		user.Email).Scan(
		&user.User_ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Phone_Number,
		&user.Is_Active,
	)
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
		is_admin
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
			&user.IsAdmin,
			//&user.UserAddressID,
		); err != nil {
			return users, err
		}
		if !user.IsAdmin {
			users = append(users, user)
		}
	}
	if err = row.Err(); err != nil {
		return users, err
	}
	return users, err
}

func (r Repository) FindUserByEmail(user models.User) (models.User, error) {

	//var user models.User

	query := `SELECT 
		user_id,
		first_name,
		last_name,
		password,
		email,
		is_active 
		FROM 
		users 
		WHERE email = $1`

	err := r.DB.QueryRow(query, user.Email).Scan(
		&user.User_ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Password,
		&user.Email,
		&user.Is_Active,
	)

	return user, err
}

func (r Repository) ForgetPasswordUpdate(user models.User, pass string) (models.ForgotPasswordInput, error) {

	var usr models.ForgotPasswordInput

	query := `UPDATE
				users
				SET password = $1
				WHERE email = $2
				RETURNING
				email;`
	err := r.DB.QueryRow(query, pass, user.Email).Scan(
		&usr.Email,
	)
	return usr, err
}

func (r Repository) AddWishlist(wishlist models.Wishlist) (models.Wishlist, error) {
	query := `INSERT INTO wishlist(
		product_id,
		user_id)
		VALUES($1,$2)
		RETURNING
		wishlist_id,
		user_id,
		product_id;`

	err := r.DB.QueryRow(query,
		wishlist.Product_ID.Product_ID,
		wishlist.User_ID.User_ID,
	).Scan(
		&wishlist.Wishlist_ID,
		&wishlist.Product_ID.Product_ID,
		&wishlist.User_ID.User_ID,
	)

	return wishlist, err
}

func (r Repository) ViewWishlist(filter models.Filter, wlist models.Wishlist) ([]wishList, models.Metadata, error) {
	var lists []wishList

	query := `SELECT COUNT(*) OVER(),
	product.product_id,	
	product.product_price,
	product.product_name,
	users.user_id
	FROM
	wishlist
	INNER JOIN 
	product ON wishlist.product_id = product.product_id
	INNER JOIN
	users ON wishlist.user_id=users.user_id
	WHERE wishlist.user_id=$1
	LIMIT $2 OFFSET $3;`

	rows, err := r.DB.Query(query,
		wlist.User_ID.User_ID,
		filter.Limit(),
		filter.Offset())
	if err != nil {
		return nil, models.Metadata{}, err
	}
	defer rows.Close()

	var toatalRecords int

	for rows.Next() {
		var list wishList
		if err := rows.Scan(&toatalRecords,
			&list.Id,
			&list.Price,
			&list.PrName,
			&list.PrUserID,
		); err != nil {
			return lists, models.Metadata{}, err
		}
		lists = append(lists, list)
	}
	if err = rows.Err(); err != nil {
		return lists, models.Metadata{}, err
	}

	return lists, models.ComputeMetadata(toatalRecords, filter.Page, filter.PageSize), nil
}

func (r Repository) DeleteProductWishlist(delete models.Wishlist) (int64, error) {

	query := `DELETE FROM 
				wishlist
					WHERE product_id = $1 AND user_id = $2;`
	res, err := r.DB.Exec(query, delete.Product_ID.Product_ID, delete.User_ID.User_ID)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, err
}
