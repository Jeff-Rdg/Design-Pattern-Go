package transfer

import (
	"fmt"
	"github.com/google/uuid"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/interfaces"
	"patterns/1_FactoryMethod/requests"
	"time"
)

type Pix struct {
	ID              uint                 `json:"id"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	Value           float64              `json:"value"`
	AccountSender   string               `json:"account_sender"`
	AccountReceiver string               `json:"account_receiver"`
	Type            enums.TransferType   `json:"type"`
	PixKey          uuid.UUID            `json:"pix_key"`
	Status          enums.TransferStatus `json:"status"`
}

func NewPix(input requests.TransferInput) interfaces.Transfer {
	return &Pix{
		ID:              interfaces.Increment(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Value:           input.Value,
		AccountSender:   input.AccountSender,
		AccountReceiver: input.AccountReceiver,
		Type:            input.Type,
		PixKey:          input.PixKey,
		Status:          enums.CREATED,
	}
}

func (p *Pix) Create() {
	fmt.Printf("Transação pix criada!\n %s", p.ToString())
	time.Sleep(time.Second * 1)
	p.Status = enums.PROCESSING
	fmt.Printf("Processando...\n %s", p.ToString())

	time.Sleep(time.Second * 1)
	p.Status = enums.CONCLUDED
	fmt.Printf("Transação realizada com sucesso!\n %s", p.ToString())
}

func (p *Pix) ToString() string {
	return fmt.Sprintf(
		"Chave pix: %s\n"+
			"Valor: %v\n"+
			"Status: %s\n"+
			"Conta pagadora: %s\n"+
			"Conta recebedora: %s\n\n", p.PixKey, p.Value, p.Status, p.AccountSender, p.AccountReceiver)
}
