package transaction

type Repository interface {
	Create(t *Transaction) error
	GetLast(count int) ([]Transaction, error)
}
