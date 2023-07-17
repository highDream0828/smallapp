package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"database/sql"
	"github.com/highdream0828/smallapp/data/dbspeeds"
	"github.com/highdream0828/smallapp/data/validators"
	"github.com/highdream0828/smallapp/data/queries"
	"github.com/highdream0828/smallapp/data/models"
	"golang.org/x/crypto/bcrypt"
)
// Hash password
func HashPassword(password string) string {
	// Hashes the password string with a work factor of 14
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// Handle error
    if err != nil {
        log.Fatal(err)  
		return  
    }
    return string(bytes) 
}

// Register new user
func Register(c echo.Context) error {
	// Parse request body
	var user models.User
	// Binding the &user struct from the request
	c.Bind(&user)
	
	// Validate input
	if err := validators.ValidateUser(user); err != nil {
		return c.JSON(400, err)
	}

	user.Password = HashPassword(user.Password)
	// Insert user into database without ORM   
	result, err := queries.CreateUser(user)   
	if err != nil {
		return c.JSON(500, err)
	}
	
	return c.JSON(200, result)
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		// Handle error
		return c.JSON(http.StatusUnauthorized, "Password does not match!") 
	} 
    // Generate JWT token     
    token, err := createToken(user)
    if err != nil {
         return c.JSON(500, err)
    }
        
    // Return token        
    return c.JSON(200, token)
}

func createToken(user models.User) (string, error) {
    // Create token 
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "name": user.Name,
        "email": user.Email,
        "exp": time.Now().Add(time.Minute * 30).Unix(),
    })
    // Generate encoded token and return 
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }  
   
    return tokenString, nil   
}
