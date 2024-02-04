package transfer

import (
	"fmt"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/interfaces"
	"patterns/1_FactoryMethod/requests"
	"time"
)

type Ted struct {
	ID              uint                 `json:"id"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	Value           float64              `json:"value"`
	AccountSender   string               `json:"account_sender"`
	AccountReceiver string               `json:"account_receiver"`
	Type            enums.TransferType   `json:"type"`
	Status          enums.TransferStatus `json:"status"`
}

func NewTed(input requests.TransferInput) interfaces.Transfer {
	return &Ted{
		ID:              interfaces.Increment(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Value:           input.Value,
		AccountSender:   input.AccountSender,
		AccountReceiver: input.AccountReceiver,
		Type:            input.Type,
		Status:          enums.CREATED,
	}
}

func (p *Ted) Create() {
	fmt.Printf("Transação Ted criada!\n %s", p.ToString())
	time.Sleep(time.Second * 2)
	p.Status = enums.PROCESSING
	fmt.Printf("Processando...\n %s", p.ToString())
	time.Sleep(time.Second * 2)
	p.Status = enums.ANALYSIS
	fmt.Printf("Status da transação: %s", p.ToString())
}

func (p *Ted) ToString() string {
	return fmt.Sprintf(
		"Valor: %v\n"+
			"Status: %s\n"+
			"Conta pagadora: %s\n"+
			"Conta recebedora: %s\n\n", p.Value, p.Status, p.AccountSender, p.AccountReceiver)
}
