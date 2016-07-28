package models

type Server struct {
	Ip          string
	Port        int
	Auth        string
	Driver		string
	Location    string
	Title		string
	Description string
	Amount      int
	Have      	int
}
