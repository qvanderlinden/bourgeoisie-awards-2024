package main

import (
	"context"
	"database/sql"
	"io/fs"
	"log"
	"os"
	"path"
	"slices"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	handleError(err)

	pool, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	handleError(err)
	defer pool.Close()

	tx, err := pool.BeginTx(context.Background(), &sql.TxOptions{})
	handleError(err)

	seedsDir := "./database/seeds/data"
	entries, err := os.ReadDir(seedsDir)
	slices.SortFunc(entries, func(a fs.DirEntry, b fs.DirEntry) int {
		if a.Name() < b.Name() {
			return -1
		}
		return 1
	})
	handleError(err)
	for _, entry := range entries {
		info, err := entry.Info()
		handleError(err)

		filepath := path.Join(seedsDir, info.Name())
		bytes, err := os.ReadFile(filepath)
		handleError(err)

		script := string(bytes)
		_, err = tx.Exec(script)
		handleError(err)
	}
	tx.Commit()
}
