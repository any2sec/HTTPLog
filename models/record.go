package models

type Record struct {
	Random		 	string		`json:"random" gorm:"size:255"`		//random
	RequestIP		string		`json:"request_ip" gorm:"size:20"`
	CreatedAt   	string	 	`json:"createAt" gorm:"size:100"`
}

type RespData struct {
	HTTPStatusCode string
	Msg            string
}
