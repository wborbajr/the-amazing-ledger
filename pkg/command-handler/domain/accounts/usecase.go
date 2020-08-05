package accounts

type UseCase interface {
	CreateAccount(AccountInput) error
}

type AccountInput struct {
	Type     string
	OwnerID  string
	Owner    string
	Name     string
	Metadata []string
}

type Account struct {
	OwnerID  string   `json:"owner_id"`
	Type     string   `json:"type"`
	Balance  int      `json:"balance"`
	Owner    string   `json:"owner"`
	Name     string   `json:"name"`
	Metadata []string `json:"metadata"`
}