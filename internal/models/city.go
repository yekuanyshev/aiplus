package models

type City struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
