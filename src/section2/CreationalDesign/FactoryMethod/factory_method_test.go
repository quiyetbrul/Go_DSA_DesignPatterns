package factor_method

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Error("A payment method of type Cash must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "cash") {
		t.Error("The cash payment method must contain the word cash")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Error("A payment method of type DebitCard must exist")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "debit card") {
		t.Error("The debit card payment method must contain the word debit card")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
