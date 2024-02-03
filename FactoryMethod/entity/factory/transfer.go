package factory

import (
	"errors"
	"patterns/FactoryMethod/entity/enums"
	"patterns/FactoryMethod/entity/interfaces"
	"patterns/FactoryMethod/entity/transfer"
	"patterns/FactoryMethod/requests"
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
