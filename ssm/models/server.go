package models

type Server struct {
	Ip          string
	Port        int
	Auth        string
	Location    string
	Description string
	Amount      int
	Remain      int
}
