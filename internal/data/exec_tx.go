package data

import (
	"context"
	"database/sql"
	"fmt"
)

func (m Models) ExecInTx(fn func(*Models) error) error {
	tx, err := m.db.BeginTx(
		context.Background(),
		&sql.TxOptions{},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	q := NewModelsWithTx(tx)
	err = fn(&q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf(
				"tx err: %v, rb err: %v",
				err,
				rbErr,
			)
		}
		return err
	}
	return tx.Commit()

}
