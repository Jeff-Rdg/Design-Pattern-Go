package main

import (
	"fmt"
	"patterns/1_FactoryMethod/enums"
	"patterns/1_FactoryMethod/factory"
	"patterns/1_FactoryMethod/requests"
)

func main() {
	inputPix := requests.Input{
		Value:             50,
		TaxNumberSender:   "520.585.330-11",
		TaxNumberReceiver: "134.688.200-20",
		Type:              enums.PIX,
	}

	inputTed := requests.Input{
		Value:             25,
		TaxNumberSender:   "476.290.150-46",
		TaxNumberReceiver: "301.601.650-83",
		Type:              enums.TED,
	}

	inputErr := requests.Input{
		Value:             25,
		TaxNumberSender:   "870.676.640-47",
		TaxNumberReceiver: "235.103.660-34",
		Type:              "invalid",
	}

	transfer, err := factory.NewTransfer(inputPix)
	if err != nil {
		fmt.Println(err)
	} else {
		transfer.Flow()
	}

	transfer2, err := factory.NewTransfer(inputTed)
	if err != nil {
		fmt.Println(err)
	} else {
		transfer2.Flow()
	}

	transfer3, err := factory.NewTransfer(inputErr)
	if err != nil {
		fmt.Println(err)
	} else {
		transfer3.Flow()
	}
}
