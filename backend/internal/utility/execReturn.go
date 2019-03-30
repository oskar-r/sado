package utility

import (
	"database/sql"
	"errors"
	"log"
)

func ExecReturn(res sql.Result, err error) error {
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	i, err := res.RowsAffected()
	if i == 0 {
		log.Printf("[ERROR] no record created")
		return errors.New("no record created")
	}
	return err
}
