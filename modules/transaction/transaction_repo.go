package transaction

import (
	"altaStore/business/transaction"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type Transaction struct {
	ID               uint   `gorm:"id;primaryKey;autoIncrement"`
	CartID           uint   `json:"cart_id" validate:"required" gorm:"not null"`
	InvoiceNumber    string `json:"invoice_number" validate:"required" gorm:"not null"`
	Status           string `json:"status" validate:"required" gorm:"type:varchar(25); not null"`
	TotalTransaction uint   `json:"total_transaction" validate:"gt=0" gorm:"not null"`
	ShippingFee      uint   `json:"shipping_fee" validate:"gte=0" gorm:"not null"`
	PayAt            *time.Time
	CancelAt         *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func newTransactionTable(transaction transaction.Transaction) *Transaction {

	return &Transaction{
		ID:               0,
		CartID:           transaction.CartID,
		InvoiceNumber:    transaction.InvoiceNumber,
		Status:           transaction.Status,
		TotalTransaction: transaction.TotalTransaction,
		ShippingFee:      transaction.ShippingFee,
		PayAt:            &time.Time{},
		CancelAt:         &time.Time{},
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        &time.Time{},
	}

}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (col *Transaction) TransactionToService() transaction.Transaction {
	var transaction transaction.Transaction

	transaction.ID = col.ID
	transaction.CartID = col.CartID
	transaction.InvoiceNumber = col.InvoiceNumber
	transaction.Status = col.Status
	transaction.TotalTransaction = col.TotalTransaction
	transaction.ShippingFee = col.ShippingFee
	transaction.PayAt = col.PayAt
	transaction.CancelAt = col.CancelAt
	transaction.CreatedAt = col.CreatedAt
	transaction.UpdatedAt = col.UpdatedAt
	transaction.DeletedAt = col.DeletedAt

	return transaction
}

func (repo *GormRepository) CreateTransaction(transaction transaction.Transaction) error {
	transactionData := newTransactionTable(transaction)

	err := repo.DB.Create(transactionData).Error

	if err != nil {
		return err
	}

	return nil
}
