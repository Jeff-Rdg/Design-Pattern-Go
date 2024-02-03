package requests

import "patterns/FactoryMethod/entity/enums"

type Input struct {
	Value             float64            `json:"value"`
	TaxNumberSender   string             `json:"tax_number_sender"`
	TaxNumberReceiver string             `json:"tax_number_receiver"`
	Type              enums.TransferType `json:"type"`
}
