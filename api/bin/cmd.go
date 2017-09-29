package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/tonkla/op3n/api/province"
)

func main() {
	importProvincesData()
}

func importProvincesData() {
	province.Import()
}
