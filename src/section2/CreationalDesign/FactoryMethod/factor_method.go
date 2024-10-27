package factor_method

import "errors"

/*
- delegates the creation of different type of payments
- the second-best known and used design pattern in the industry
- abstracts the user from the knowledge of the struct needed to achieve
- provides an interface that fits the user needs
- eases the process of downgrading or upgrading the implementation

- gain extra layer of encapsulation
- program grows in a controlled environment
- can delegate creation of families of objs to different pkg
*/

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

// CreatePaymentMethod returns a pointer to a PaymentMethod obj or an error
// if the method is not registered. We used "new" operator to return the pointer
// but we could also use &Type{} although "new" makes it more readable for
// newcommers to Go.
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New("Payment method not recognized")
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return "paid using cash"
}

func (d *DebitCardPM) Pay(amount float32) string {
	return "paid using debit card"
}
