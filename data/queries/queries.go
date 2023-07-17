package queries

import (
	"database/sql"
	"github.com/highdream0828/smallapp/data/dbspeeds"
	"github.com/highdream0828/smallapp/data/models"
)

func GetUserByEmail(email string) (models.User, error) {

	var user models.User
	row := dbspeeds.DB.QueryRow("SELECT * FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil  
		}
		return user, err   
	}

	return user, nil
}

func CreateUser(user models.User) (int, error) {

	result := dbspeeds.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

	id, err := result.LastInsertId()   
	if err != nil {
		return 0, err  
	}   

	return int(id), nil
}