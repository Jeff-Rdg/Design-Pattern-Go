package enums

type TransferType string

const (
	PIX TransferType = "pix"
	TED TransferType = "ted"
)

type TransferStatus string

const (
	CREATED    TransferStatus = "criado"
	CONCLUDED  TransferStatus = "concluído"
	PROCESSING TransferStatus = "em processamento"
	FAILED     TransferStatus = "falhou"
	ANALYSIS   TransferStatus = "em análise"
)
