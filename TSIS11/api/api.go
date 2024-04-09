package api

type Fighter struct {
	Id             int    `json:"id"`
	First_name     string `json:"firstName"`
	Last_name      string `json:"lastName"`
	WeightCategory string `json:"weight_category"`
}

var Fighters = []Fighter{
	{Id: 1, First_name: "Islam", Last_name: "Makhachev", WeightCategory: "lightweight"},
	{Id: 2, First_name: "Jon", Last_name: "Jones", WeightCategory: "heavyweight"},
	{Id: 3, First_name: "Alexander", Last_name: "Volkanovski", WeightCategory: "featherweight"},
	{Id: 4, First_name: "Leon", Last_name: "Edvards", WeightCategory: "welterweight"},
	{Id: 5, First_name: "Alex", Last_name: "Pereira", WeightCategory: "light heavyweight"},
	{Id: 6, First_name: "Charles", Last_name: "Oliveira", WeightCategory: "lightweight"},
	{Id: 7, First_name: "Sean", Last_name: "O'Malley", WeightCategory: "bantamweight"},
	{Id: 8, First_name: "Dricus", Last_name: "Du Plessis", WeightCategory: "middleweight"},
	{Id: 9, First_name: "Alexander", Last_name: "Pantoja", WeightCategory: "flyweight"},
	{Id: 10, First_name: "Israel", Last_name: "Adesanya", WeightCategory: "middleweight"},
	{Id: 11, First_name: "Sean", Last_name: "Stricland", WeightCategory: "middleweight"},
	{Id: 12, First_name: "Aljamain", Last_name: "Sterling", WeightCategory: "featherweight"},
	{Id: 13, First_name: "tom", Last_name: "Aspinall", WeightCategory: "heavyweight"},
	{Id: 14, First_name: "Max", Last_name: "Holloway", WeightCategory: "featherweight"},
	{Id: 15, First_name: "Kamaru", Last_name: "Usman", WeightCategory: "middleweight"},
}
