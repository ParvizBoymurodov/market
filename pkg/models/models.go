package models

import (
	"time"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Dsn  string `json:"dsn"`
	Key  string `json:"key"`
}

type Prices struct { //для создания таблицы Прайс
	ID int64
	Name string
	Category string
	Price int
	Removed bool
}

type Clients struct { // для создания таблицы Клиент
	ID int64
	Client string
	City string
	Region string
}

type Sales struct { // для добавления таблицы Продаж
	ID int64
	Date time.Time
	Product string
	Category string
	Count int
	Sum int
	Client string
}




















//type InsertPrices struct { // для добавления продукта
//	Name string
//	Category string
//	Price decimal.Decimal
//}
//
//type InsertSales struct { // для добавления продажи в бд
//	Date time.Time
//	Product string
//	Price decimal.Decimal
//	Client string
//}
