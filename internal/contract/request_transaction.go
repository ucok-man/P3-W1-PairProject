package contract

type ReqTransactionUpdate struct {
	Id          string  `json:"string" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}

type ReqTransactionDelete struct {
	Id string `json:"string" validate:"required"`
}

type ReqTransactionInsert struct {
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}
