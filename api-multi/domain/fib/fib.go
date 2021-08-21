package fib

import (
	"fmt"
	"time"

	"github.com/thanhftu/api-multi/db/pg/fibnumdb"
)

const (
	InserQuery     = "INSERT INTO fib(index,value) VALUES($1,$2) RETURNING *;"
	getQuery       = "SELECT * FROM fib WHERE index=$1;"
	getQueryAll    = "SELECT * FROM fib;"
	deleteQuery    = "DELETE FROM fib WHERE id=$1;"
	getLatestQuery = "SELECT * FROM fib ORDER BY id DESC LIMIT 1"
)

type FibNumber struct {
	ID         int64
	Index      int64
	Value      int64
	Created_at time.Time
}

func (fib *FibNumber) SAVE() error {
	stmt, err := fibnumdb.Client.Prepare(InserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	errSave := stmt.QueryRow(fib.Index, fib.Value).Scan(&fib.ID, &fib.Index, &fib.Value, &fib.Created_at)
	if errSave != nil {
		return err
	}
	fmt.Println(fib.ID)
	return nil
}

func (fib *FibNumber) GET() error {
	stmt, err := fibnumdb.Client.Prepare(getQuery)
	if err != nil {
		return err
	}
	result := stmt.QueryRow(fib.Index)
	if err := result.Scan(&fib.ID, &fib.Index, &fib.Value, &fib.Created_at); err != nil {
		return err
	}
	return nil
}

func (fib *FibNumber) GETLATEST() error {
	stmt, err := fibnumdb.Client.Prepare(getLatestQuery)
	if err != nil {
		return err
	}
	result := stmt.QueryRow()
	if err := result.Scan(&fib.ID, &fib.Index, &fib.Value); err != nil {
		return err
	}
	return nil
}

func GETALL() ([]FibNumber, error) {
	rows, err := fibnumdb.Client.Query(getQueryAll)
	if err != nil {
		return nil, err
	}
	Fibs := make([]FibNumber, 0)
	for rows.Next() {
		fib := FibNumber{}
		if err := rows.Scan(&fib.ID, &fib.Index, &fib.Value, &fib.Created_at); err != nil {
			return nil, err
		}
		Fibs = append(Fibs, fib)
	}

	return Fibs, nil
}

func (fib *FibNumber) DELETE() error {
	stmt, err := fibnumdb.Client.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, errExe := stmt.Exec(fib.ID)
	if errExe != nil {
		return err
	}
	return nil
}
