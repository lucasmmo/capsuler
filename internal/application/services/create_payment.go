package services

import (
	"capsuler/internal/domain/payment/model"

	"github.com/stripe/stripe-go/v82/checkout/session"
)

type Repository interface {
	Save(payment *model.Payment) error
	GetById(id string) (*model.Payment, error)
	Remove(id string) error
	Count() int
}

type CreatePayment struct {
	client            *session.Client
	paymentRepository Repository
}

type PaymentId string

func (c *CreatePayment) Create(ownerId string, sessionId string) (PaymentId, error) {
	return
}
