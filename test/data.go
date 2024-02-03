package data

type P4P_fighter struct{
	Id int `json:"id"`
	first_name string `json:"first_name"`
	last_name string `json:"last_name"`
	weight_category string `json:"weight_category"`
}

var Fighters []P4P_fighter{
	{Id:1, first_name: "Islam", last_name: "Makhachev", weight_category: "lightweight"},
	{Id:2, first_name: "Jon", last_name: "Jones", weight_category: "heavyweight"},
	{Id:3, first_name: "Alexander", last_name: "Volkanovski", weight_category: "featherweight"},
	{Id:4, first_name: "Leon", last_name: "Edvards", weight_category: "welterweight"},
	{Id:5, first_name: "Alex", last_name: "Pereira", weight_category: "light heavyweight"},
	{Id:6, first_name: "Charles", last_name: "Oliveira", weight_category: "lightweight"},
	{Id:7, first_name: "Sean", last_name: "O'Malley", weight_category: "bantamweight"},
	{Id:8, first_name: "Dricus", last_name: "Du Plessis", weight_category: "middleweight"},
	{Id:9, first_name: "Alexander", last_name: "Pantoja", weight_category: "flyweight"},
	{Id:10, first_name: "Israel", last_name: "Adesanya", weight_category: "middleweight"},
	{Id:11, first_name: "Sean", last_name: "Stricland", weight_category: "middleweight"},
	{Id:12, first_name: "Aljamain", last_name: "Sterling", weight_category: "featherweight"},
	{Id:13, first_name: "tom", last_name: "Aspinall", weight_category: "heavyweight"},
	{Id:14, first_name: "Max", last_name: "Holloway", weight_category: "featherweight"},
	{Id:15, first_name: "Kamaru", last_name: "Usman", weight_category: "middleweight"},

}