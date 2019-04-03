package data

import (
	"database/sql"
	"strings"

	"github.com/casbin/casbin/model"
	"github.com/lib/pq"
)

const (
	schema          = "policies"
	table           = "policies"
	ptype           = "ptype"
	sub             = "sub"
	obj             = "obj"
	oid             = "oid"
	act             = "act"
	numberOfColumns = 5
)

func (r *Repos) CreatePolicyDB() error {

	_, err := r.db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`CREATE TABLE IF NOT EXISTS ` + schema + `.` + table + ` (
		` + ptype + ` VARCHAR(10) NULL, 
		` + sub + ` VARCHAR(256) NULL, 
		` + obj + ` VARCHAR(256) NULL, 
		` + oid + ` VARCHAR(256) NULL, 
		` + act + ` VARCHAR(256) NULL,
		CONSTRAINT policies_un UNIQUE (` + ptype + `, ` + sub + `, ` + obj + `,` + oid + `, ` + act + ` )
		)`)

	if err != nil {
		panic(err)
	}

	return nil
}

func (r *Repos) DropPolicyDB() error {
	_, err := r.db.Exec(`DROP table ` + schema + `.` + table)
	if err != nil {
		panic(err)
	}
	return err
}

func (r *Repos) LoadPolicies(namespace string) ([]string, error) {

	rows, err := r.db.Query(`SELECT ` + ptype + `,` + sub + `,` + obj + `,` + oid + `,` + act + ` FROM ` + schema + `.` + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []string
	var (
		ptype sql.NullString
		v1    sql.NullString
		v2    sql.NullString
		v3    sql.NullString
		v4    sql.NullString
	)
	for rows.Next() {
		err := rows.Scan(&ptype, &v1, &v2, &v3, &v4)
		if err != nil {
			return nil, err
		}
		line := ""
		if ptype.Valid && ptype.String != "" {
			line = ptype.String
		}
		if v1.Valid && v1.String != "" {
			line += ", " + v1.String
		}
		if v2.Valid && v2.String != "" {
			line += ", " + v2.String
		}
		if v3.Valid && v3.String != "" {
			line += ", " + v3.String
		}
		if v4.Valid && v4.String != "" {
			line += ", " + v4.String
		}
		if line != "" {
			p = append(p, line)
		}
	}
	err = rows.Err()
	return p, err
}

var insertPolicy = `INSERT INTO ` + schema + `.` + table + ` VALUES(` + strings.TrimRight(strings.Repeat("?,", numberOfColumns), ",") + `)`

func (r *Repos) SaveAllPolicy(namespace string, model model.Model) error {

	stm, err := r.db.Prepare(insertPolicy)
	if err != nil {
		return err
	}
	defer stm.Close()

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			if err = writeTableLine(stm, ptype, rule); err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			if err = writeTableLine(stm, ptype, rule); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Repos) AddPolicy(i []interface{}) (int64, error) {
	var ri int64
	res, err := r.db.Exec(insertPolicy, i...)

	if err != nil {
		if driverErr, ok := err.(*pq.Error); ok {
			if driverErr.Code.Name() == "23505" {
				err = nil
				ri = 0
			} else {
				return 0, err
			}
		} else {
			return 0, err
		}
	} else {
		ri, err = res.RowsAffected()
	}
	return ri, err
}

const deletePolicy = `DELETE FROM ` + schema + `.` + table + ` WHERE 
		` + ptype + `=$1
		'` + sub + `=$2
		'` + obj + `=$3
		'` + oid + `=$4
		,` + act + `=$5`

func (r *Repos) DeletePolicy(i []interface{}) (int64, error) {
	res, err := r.db.Exec(createDeleteQuery(len(i)), i...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func createDeleteQuery(index int) string {
	q := `DELETE FROM ` + schema + `.` + table + ` WHERE `
	switch index {
	case 1:
		q = q + ptype + `=$1`
	case 2:
		q = q + ptype + `=$1 AND ` + sub + `=$2`
	case 3:
		q = q + ptype + `=$1 AND ` + sub + `=$2 AND ` + obj + `=$3`
	case 4:
		q = q + ptype + `=$1 AND ` + sub + `=$2 AND ` + obj + `=$3 AND ` + oid + `=$4`
	case 5:
		q = q + ptype + `=$1 AND ` + sub + `=$2 AND ` + obj + `=$3 AND ` + oid + `=$4 AND ` + act + `=$5`
	}

	return q
}

func writeTableLine(stm *sql.Stmt, ptype string, rule []string) error {
	params := make([]interface{}, 0, numberOfColumns)
	params = append(params, ptype)
	for _, v := range rule {
		params = append(params, v)
	}
	need := 5 - len(params)
	for i := 0; i < need; i++ {
		params = append(params, "")
	}
	if _, err := stm.Exec(params...); err != nil {
		return err
	}
	return nil
}
