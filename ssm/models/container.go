package models

type Container struct{
	Id		string
	Name	[]string
	Image	string
	Status	string
	Ports	[]Port
}

type Port struct{
	PrivatePort		int
	PublicPort		int
	Type 			string
}