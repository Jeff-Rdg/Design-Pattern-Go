package main

import (
	"fmt"
	"github.com/google/uuid"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/factory"
	"patterns/1_FactoryMethod/requests"
)

func main() {
	inputPix := requests.TransferInput{
		Value:           50,
		AccountSender:   "123456-1",
		AccountReceiver: "654321-1",
		Type:            enums.PIX,
		PixKey:          uuid.New(),
	}

	inputTed := requests.TransferInput{
		Value:           25,
		AccountSender:   "123456-1",
		AccountReceiver: "654321-1",
		Type:            enums.TED,
	}

	inputErr := requests.TransferInput{
		Value:           25,
		AccountSender:   "123456-1",
		AccountReceiver: "654321-1",
		Type:            "invalid",
	}

	transfer, err := factory.NewTransfer(inputPix)
	if err != nil {
		fmt.Println(err)
	} else {
		transfer.Create()
	}

	transfer2, err := factory.NewTransfer(inputTed)
	if err != nil {
		fmt.Println(err)
	} else {
		transfer2.Create()
	}

	transfer3, err := factory.NewTransfer(inputErr)
	if err != nil {
		fmt.Println(err)
	} else {
		transfer3.Create()
	}
}
