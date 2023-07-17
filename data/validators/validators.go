package validators

import (
	"errors"
	"github.com/highdream0828/smallapp/data/dbspeeds"
)


func ValidateUser(user models.User) error {

    if user.Name == "" {
        return errors.New("name is required")
    }
    
    if !isValidEmail(user.Email) {
        return errors.New("invalid email address") 
    }
    
    if user.Password == "" {
        return errors.New("password is required")    
    }
    
    return nil
}

func ValidateLogin(credentials models.Credentials) error {
   
    if credentials.Email == "" {
        return errors.New("email is required")    
    }
    
    if credentials.Password == "" {
        return errors.New("password is required")    
    }
    
    return nil    
}