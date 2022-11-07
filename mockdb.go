package main

/*
Package db - пакет враппер для работы с базами

Состав пакета:

	db.go - набор общих интерфейсов для использования в коде
	mgo.go - реализация для драйвера globalsign/mgo
	mock.go - набор mock-структур для проведения тестирования
*/

import (
	"github.com/chrisloarryn/mongodb-boltdb-mock/db"
)

//New - экспортируемая функция-обёртка, для упрощения создания переменной интерфейсного типа
func New(self db.Handler) db.Handler {
	return self
}

