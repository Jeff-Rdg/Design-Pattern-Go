package factory

import (
	"errors"
	"patterns/1_FactoryMethod/entity/transfer"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/interfaces"
	"patterns/1_FactoryMethod/requests"
)

func NewTransfer(input requests.TransferInput) (interfaces.Transfer, error) {
	switch input.Type {
	case enums.PIX:
		return transfer.NewPix(input), nil
	case enums.TED:
		return transfer.NewTed(input), nil
	default:
		return nil, errors.New("Invalid transfer_type")
	}
}
