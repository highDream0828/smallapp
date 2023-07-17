package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"database/sql"
	"github.com/highdream0828/smallapp/data/dbspeeds"
	"github.com/highdream0828/smallapp/data/validators"
	"github.com/highdream0828/smallapp/data/queries"
	"golang.org/x/crypto/bcrypt"
)
// Hash password
func HashPassword(password string) string {
	// Hashes the password string with a work factor of 14
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    handleError(err)  
    return string(bytes) 
}

// Register new user
func Register(c echo.Context) error {
	// Parse request body
	var user struct {
		Name string json:"name"
		Email string json:"email"
		Password string json:"password"
	}
	// Binding the &user struct from the request
	c.Bind(&user)
	
	// Validate input
	if err := validators.ValidateUser(user); err != nil {
		return c.JSON(400, err)
	}

	user.Password = hashPassword(user.Password)
	// Insert user into database without ORM   
	result, err := queries.CreateUser(user)   
	if err != nil {
		return c.JSON(500, err)
	}
	
	return c.JSON(200, result.RowsAffected)
}

func Login(c echo.Context) error {
  
    // Get and validate input credentials
    var creds struct {
        Email    string `form:"email" validate:"required,email"`
        Password string `form:"password" validate:"required"` 
    }   
    c.Bind(&creds)
    if err := validators.ValidateUser(creds); err != nil {
        return c.JSON(400, err)
    }
        
    // Query database for user 
    user, err := queries.GetUserByEmail(creds.email)
    if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(404, "User not found")
		}   
		return c.JSON(500, "Error getting user")  
	}
	
    // Validate password
    if !bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) {
        c.JSON(http.StatusUnauthorized, ErrInvalidCredentials) 
   		return
    }
        
    // Generate JWT token     
    token, err := createToken(user.ID)
    if err != nil {
         return c.JSON(500, err)
    }
        
    // Return token        
    return c.JSON(200,  {"token": token})
}
