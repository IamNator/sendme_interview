package services

import (
	"errors"

	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/models"
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	DB *gorm.DB
}

func NewTransaction(db *gorm.DB) Transaction {
	return Transaction{db}
}

func (t Transaction) DebitUser(debit models.DebitUser) (*schema.Wallet, error) {

	var wallet schema.Wallet
	result := t.DB.Table(wallet.TableName()).FirstOrCreate(&wallet)
	if er := result.Error; er != nil {
		return nil, er
	}

	if bal := (wallet.Balance - debit.Amount); bal > -1 {
		wallet.Balance = bal
		result := t.DB.Table(wallet.TableName()).Where("user_id = ?", debit.UserID).Update(&wallet).First(&wallet)
		if er := result.Error; er != nil {
			return nil, er
		}
	} else {
		return nil, errors.New("insuffiecient funds")
	}

	var transactionLog schema.Transaction
	transactionLog.Type = "debit"
	transactionLog.UserID = uint(debit.UserID)
	transactionLog.Amount = debit.Amount
	//	transactionLog.TransactionID = uint(debit.TransactionID)

	result = t.DB.Table(transactionLog.TableName()).Save(&transactionLog)
	if er := result.Error; er != nil {
		return nil, er
	}

	return &wallet, nil
}

func (t Transaction) CreditUser(credit models.CreditUser) (*schema.Wallet, error) {

	var wallet schema.Wallet
	result := t.DB.Table(wallet.TableName()).FirstOrCreate(&wallet)
	if er := result.Error; er != nil {
		return nil, er
	}

	if bal := (wallet.Balance + credit.Amount); bal > -1 {
		wallet.Balance = bal
		result := t.DB.Table(wallet.TableName()).Where("user_id = ?", credit.UserID).Update(&wallet).First(&wallet)
		if er := result.Error; er != nil {
			return nil, er
		}
	} else {
		return nil, errors.New("insuffiecient funds")
	}

	var transactionLog schema.Transaction
	transactionLog.Type = "credit"
	transactionLog.UserID = uint(credit.UserID)
	transactionLog.Amount = credit.Amount
	//	transactionLog.TransactionID = uint(debit.TransactionID)

	result = t.DB.Table(transactionLog.TableName()).Save(&transactionLog)
	if er := result.Error; er != nil {
		return nil, er
	}

	return &wallet, nil
}

func (t Transaction) WalletBalance(userID uint) (*schema.Wallet, error) {

	var wallet schema.Wallet
	result := t.DB.Table(wallet.TableName()).Where("user_id = ?").First(&wallet)
	if er := result.Error; er != nil {
		return nil, er
	}

	return &wallet, nil
}

func (t Transaction) TransactionHistory(userID uint) ([]*schema.Transaction, error) {

	var transHist []*schema.Transaction
	result := t.DB.Table(schema.Transaction{}.TableName()).Where("user_id = ?").Find(&transHist)
	if er := result.Error; er != nil {
		return nil, er
	}

	return transHist, nil
}
