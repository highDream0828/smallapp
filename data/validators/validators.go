package validators

import (
	"errors"
	"regexp"
	"github.com/highdream0828/smallapp/data/models"
)

func isValidEmail(email string) bool {
	// Validate Email
    emailRegExp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegExp.MatchString(email)
}

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