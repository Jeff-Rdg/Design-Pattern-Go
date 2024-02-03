package factory

import (
	"errors"
	"patterns/1.FactoryMethod/entity/transfer"
	"patterns/1.FactoryMethod/enums"
	"patterns/1.FactoryMethod/interfaces"
	"patterns/1.FactoryMethod/requests"
)

func NewTransfer(input requests.Input) (interfaces.Transfer, error) {
	switch input.Type {
	case enums.PIX:
		return transfer.NewPix(input), nil
	case enums.TED:
		return transfer.NewTed(input), nil
	default:
		return nil, errors.New("Invalid transfer_type")
	}
}
