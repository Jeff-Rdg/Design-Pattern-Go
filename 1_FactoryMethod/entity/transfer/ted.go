package transfer

import (
	"fmt"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/interfaces"
	"patterns/1_FactoryMethod/requests"
	"time"
)

type Ted struct {
	ID                uint               `json:"id"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	Value             float64            `json:"value"`
	TaxNumberSender   string             `json:"tax_number_sender"`
	TaxNumberReceiver string             `json:"tax_number_receiver"`
	Type              enums.TransferType `json:"type"`
}

func NewTed(input requests.Input) interfaces.Transfer {
	return &Ted{
		ID:                interfaces.Increment(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Value:             input.Value,
		TaxNumberSender:   input.TaxNumberSender,
		TaxNumberReceiver: input.TaxNumberReceiver,
		Type:              input.Type,
	}
}

func (p *Ted) Processing() string {
	return "Processando transação Ted..."
}

func (p *Ted) Send() string {
	return fmt.Sprintf("Transação Ted realizada com sucesso:\nvalor: %v\nRecebedor: %s\n", p.Value, p.TaxNumberReceiver)
}

func (p *Ted) Flow() {
	fmt.Println(p.Processing())
	time.Sleep(time.Second * 3)
	fmt.Println(p.Send())
}
