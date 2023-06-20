package data

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(int, *Account) error
	GetAccountByID(ID int) (*Account, error)
	GetAccount() (*[]Account, error)
}
