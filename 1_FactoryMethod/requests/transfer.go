package requests

import (
	"github.com/google/uuid"
	"patterns/1_FactoryMethod/enums"
)

type TransferInput struct {
	Value           float64            `json:"value"`
	AccountSender   string             `json:"account_sender"`
	AccountReceiver string             `json:"account_receiver"`
	Type            enums.TransferType `json:"type"`
	PixKey          uuid.UUID          `json:"pix_key"`
}
