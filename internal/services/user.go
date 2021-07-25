package services

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{db}
}

const TokenLifeTime = 3 * time.Hour

//CreateToken creates referral token for each user [23XX][01]
func CreateToken(userID string) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("ab0cdefghi9i8jklmn123opqrst547uv9wxyz1234ABCDEFGHIJKLMNOfg5386PQRSTUVWXYZ56789")
	b := make([]rune, 4)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters)-1)]
	}

	randByte4 := string(b[:4])

	return fmt.Sprintf("%s%s", randByte4, userID)
}

func hashAndSalt(password string) string {

	pwd := []byte(password)
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)

}

//RegisterNewUser registers new users
func (u User) RegisterNewUser(credentials models.RegistrationCredential) (*models.LoginResponse, error) {
	var response models.LoginResponse

	var usr schema.User
	usr.Username = credentials.Username
	usr.Email = credentials.Email
	usr.ID = 0
	usr.HashedPassword = hashAndSalt(credentials.Password)
	usr.Token = CreateToken(credentials.Password)
	usr.TokenExpiration = time.Now().Add(TokenLifeTime).Format("2006-01-02 3:04PM") //3 hours to token expiration

	result := u.DB.Table(schema.User{}.TableName()).Where("username = ?", usr.Username).Create(&usr).First(&response)
	if result.Error != nil {
		return nil, result.Error
	}

	return &response, nil
}

//LoginUser logins in users
func (u User) LoginUser(credentials models.LoginCredential) (*models.LoginResponse, error) {

	var response models.LoginResponse

	var usr schema.User
	result := u.DB.Table(schema.User{}.TableName()).Where("username = ?", credentials.Username).
		First(&usr).
		First(&response)

	if result.RecordNotFound() {

		return nil, fmt.Errorf("password or username incorrect-")

	}

	if result.Error != nil {
		return nil, result.Error
	}

	er := bcrypt.CompareHashAndPassword([]byte(usr.HashedPassword), []byte(credentials.Password))
	if er != nil {
		return nil, fmt.Errorf("password or username incorrect")
	}

	result = u.DB.Table(schema.User{}.TableName()).Where("username = ?", usr.Username).
		Update(map[string]interface{}{"token_expiration": time.Now().Add(TokenLifeTime).Format("2006-01-02 3:04PM")}).Find(&response) //token expires TokenLifeTime after login
	if result.Error != nil {
		return nil, result.Error
	}

	return &response, nil
}
