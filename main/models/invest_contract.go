package models

type invest_contract struct {

	Id              int
	Master_id       int
	Description     string
	Amount_min      int
	Amount_max      int
	Duration_min    int
	Duration_max    int
	Duration_type   int
	For_register    int
	Rate            float32
	Early_terminate int
	Risk_rank       int
}