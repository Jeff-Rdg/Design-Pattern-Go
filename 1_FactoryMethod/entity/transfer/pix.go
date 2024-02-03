package transfer

import (
	"fmt"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/interfaces"
	"patterns/1_FactoryMethod/requests"
	"time"
)

type Pix struct {
	ID                uint               `json:"id"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	Value             float64            `json:"value"`
	TaxNumberSender   string             `json:"tax_number_sender"`
	TaxNumberReceiver string             `json:"tax_number_receiver"`
	Type              enums.TransferType `json:"type"`
}

func NewPix(input requests.Input) interfaces.Transfer {
	return &Pix{
		ID:                interfaces.Increment(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Value:             input.Value,
		TaxNumberSender:   input.TaxNumberSender,
		TaxNumberReceiver: input.TaxNumberReceiver,
		Type:              input.Type,
	}
}

func (p *Pix) Processing() string {
	return "Processando transação Pix..."
}

func (p *Pix) Send() string {
	return fmt.Sprintf("Transação Pix realizada com sucesso:\nvalor: %v\nRecebedor: %s\n", p.Value, p.TaxNumberReceiver)
}

func (p *Pix) Flow() {
	fmt.Println(p.Processing())
	time.Sleep(time.Second * 3)
	fmt.Println(p.Send())
}
