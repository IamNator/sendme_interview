package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/models"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	DB map[string]schema.User
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

func (u user) RegisterNewUser(credentials models.RegistrationCredential) (*models.LoginResponse, error) {
	by, _ := json.Marshal(&credentials)

	var usr schema.User
	json.Unmarshal(by, &usr)
	usr.HashedPassword = hashAndSalt(credentials.Password)
	usr.ID = uint(r.Uint32())
	u.DB[usr.UserName] = usr

	var output models.LoginResponse
	json.Unmarshal(by, &output)
	output.Token = "someanf"

	return &output, nil
}

func (u user) LoginUser(credentials models.LoginCredential) (*models.LoginResponse, error) {
	if usr, ok := u.DB[credentials.Username]; ok {

		er := bcrypt.CompareHashAndPassword([]byte(usr.HashedPassword), []byte(credentials.Password))
		if er != nil {
			return nil, fmt.Errorf("password or username incorrect")
		}

		by, _ := json.Marshal(&usr)

		var output models.LoginResponse
		json.Unmarshal(by, &output)

		output.Token = "someanf"
		return &output, nil
	}

	return nil, fmt.Errorf("password or username incorrect")
}

func assertEqual(t *testing.T, want, got interface{}, msg string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("%s: want %v\n; got %v\n", msg, want, got)
	}
}

var testLogin = []struct {
	Name string
	User struct {
		RegistrationRequestBody models.RegistrationCredential
		LoginRequestBody        models.LoginCredential
		Want                    models.LoginResponse
	}
}{
	{
		// TODO: Add test cases.
		Name: "user with short username",
		User: struct {
			RegistrationRequestBody models.RegistrationCredential
			LoginRequestBody        models.LoginCredential
			Want                    models.LoginResponse
		}{
			RegistrationRequestBody: models.RegistrationCredential{
				UserName: "bobby",
				Email:    "bobby@gmail.com",
				Password: "hosty",
			},
			LoginRequestBody: models.LoginCredential{
				Username: "bobby",
				Password: "hosty",
			},
			Want: models.LoginResponse{
				Username: "bobby",
				Email:    "bobby@gmail.com",
				Token:    "someanf",
			},
		},
	},
}

//---------------------------------------

type transaction struct {
	DB struct {
		Wallet      map[int]schema.Wallet //userID ->  wallet
		Transaction map[int][]schema.Transaction
	}
}

type testTranxDB struct {
	DB    transaction
	Tests []struct {
		Name              string
		UserID            int
		DebitRequestBody  models.DebitUser
		CreditRequestBody models.CreditUser
	}
}

var source = rand.NewSource(time.Now().Unix())
var r = rand.New(source)

func (t *transaction) DebitUser(debit models.DebitUser) (*schema.Wallet, error) {

	if bal := (t.DB.Wallet[debit.UserID].Balance - debit.Amount); bal > -1 {
		wallet := t.DB.Wallet[debit.UserID]
		wallet.Balance = bal
		t.DB.Wallet[debit.UserID] = wallet

		t.DB.Transaction[debit.UserID] = append(t.DB.Transaction[debit.UserID], schema.Transaction{
			UserID: uint(debit.UserID),
			Amount: debit.Amount,
			Type:   "debit",
			ID:     uint(r.Uint32()),
		},
		)
		return &wallet, nil
	}
	return nil, fmt.Errorf("insuffient funds")
}

func (t *transaction) CreditUser(credit models.CreditUser) (*schema.Wallet, error) {

	if _, ok := t.DB.Wallet[credit.UserID]; !ok {
		return nil, fmt.Errorf("user does not exist")
	}

	if bal := (t.DB.Wallet[credit.UserID].Balance + credit.Amount); bal > -1 {
		wallet := t.DB.Wallet[credit.UserID]
		wallet.Balance = bal
		t.DB.Wallet[credit.UserID] = wallet

		// if _, ok := t.DB.Transaction[credit.UserID]; !ok {
		// 	print("how far\n")
		// 	t.DB.Transaction[credit.UserID] = []schema.Transaction{schema.Transaction{
		// 		UserID: uint(credit.UserID),
		// 		Amount: credit.Amount,
		// 		Type:   "credit",
		// 		ID:     uint(r.Uint32()),
		// 	}}

		// } else {
		t.DB.Transaction[credit.UserID] = append(t.DB.Transaction[credit.UserID], schema.Transaction{
			UserID: uint(credit.UserID),
			Amount: credit.Amount,
			Type:   "credit",
			ID:     uint(r.Uint32()),
		},
		)
		// }

		return &wallet, nil
	}
	return nil, fmt.Errorf("insuffient funds")
}

func (t *transaction) TransactionHistory(userID uint) ([]*schema.Transaction, error) {
	var output []*schema.Transaction
	if tranx, ok := t.DB.Transaction[int(userID)]; ok {

		for _, v := range tranx {
			output = append(output, &v)
		}

		return output, nil
	}

	return nil, fmt.Errorf("record not found")
}

func (t *transaction) WalletBalance(userID uint) (*schema.Wallet, error) {
	if wallet, ok := t.DB.Wallet[int(userID)]; ok {
		return &wallet, nil
	}

	return nil, fmt.Errorf("record not found")
}
