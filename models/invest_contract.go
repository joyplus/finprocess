package models


type Invest_Contract struct {
	Id              int
	Name            string
	Master_id       int
	Description     string
	Amount_min      string
	Amount_max      string
	Duration_min    string
	Duration_max    string
	Duration_type   int
	For_register    int
	Rate            float32
	Early_terminate int
	Risk_rank       int
}

