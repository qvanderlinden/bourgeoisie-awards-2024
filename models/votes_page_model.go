package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Nominee struct {
	Id   string
	Name string
}

type Category struct {
	Id          string
	Name        string
	Description string
	Nominees    []*Nominee
	Vote        string
}

type VotesPageModel struct {
	Categories []*Category
}

func NewVotesPageModel(pool *sql.DB, userId string) (*VotesPageModel, error) {
	rows, err := pool.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*Category, 0)
	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	for _, category := range categories {
		rows, err := pool.Query(`
			SELECT 
				N.id, N.name
			FROM
				categories_nominees_bridge CN,
				nominees N
			WHERE
				N.id = CN.nominee_id AND
				CN.category_id = $1
			`, category.Id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		nominees := make([]*Nominee, 0)
		for rows.Next() {
			nominee := Nominee{}
			err := rows.Scan(&nominee.Id, &nominee.Name)
			if err != nil {
				return nil, err
			}
			nominees = append(nominees, &nominee)
		}
		category.Nominees = nominees

		row := pool.QueryRow(`
			SELECT 
				nominee_id
			FROM
				votes
			WHERE
				user_id = $1 AND
				category_id = $2
			`, userId, category.Id)
		if err != nil {
			return nil, err
		}

		row.Scan(&category.Vote)
	}

	return &VotesPageModel{Categories: categories}, nil
}
