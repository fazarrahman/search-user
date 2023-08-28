package entity

type Salary struct {
	Id       string   `json:"_id"`
	Index    int64    `json:"index"`
	Guid     string   `json:"guid"`
	IsActive bool     `json:"isActive"`
	Balance  string   `json:"balance"`
	Tags     []string `json:"tags"`
	Friends  []Friend `json:"friends"`
}

type Friend struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SalaryCSV struct {
	Index    int64  `json:"index"`
	Guid     string `json:"guid"`
	IsActive bool   `json:"isActive"`
	Balance  string `json:"balance"`
	Tags     string `json:"tags"`
	Name     string `json:"name"`
}
