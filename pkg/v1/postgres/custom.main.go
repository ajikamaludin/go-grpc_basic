package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/converter"
)

type CustomMain struct {
	UserId      string         `db:"user_id,omitempty"`
	Pass        string         `db:"pass,omitempty"`
	DelFlag     bool           `db:"del_flag,omitempty"`
	Description sql.NullString `db:"description,omitempty"`
	CreId       string         `db:"cre_id,omitempty"`
	CreTime     time.Time      `db:"cre_time,omitempty"`
	ModId       string         `db:"mod_id,omitempty"`
	ModTime     time.Time      `db:"mod_time,omitempty"`
	ModTs       int            `db:"mod_td,omitempty"`
}

func (c *Conn) CustomMainSelect(filter *CustomMain) ([]*CustomMain, error) {
	//create query syntax
	qsyntax := fmt.Sprintf(`SELECT * FROM %s`, constants.Table_Custom_Main)

	fil := reflect.ValueOf(*filter)

	for i := 0; i < fil.NumField(); i++ {
		if !fil.Field(i).IsZero() {
			if i == 0 {
				qsyntax = fmt.Sprintf(`%s WHERE`, qsyntax)
			}
			qsyntax = fmt.Sprintf(`%s %s = '%v'`, qsyntax, converter.CamelToSnakeCase(fil.Type().Field(i).Name), fil.Field(i))
		}
	}

	qsyntax = strings.TrimRight(qsyntax, "AND")
	qsyntax = fmt.Sprintf(`%s;`, qsyntax)

	fmt.Println(qsyntax)
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(qsyntax)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []*CustomMain

	for rows.Next() {
		var rowsScan CustomMain
		err := rows.Scan(&rowsScan.UserId, &rowsScan.Pass,
			&rowsScan.DelFlag, &rowsScan.Description, &rowsScan.CreId,
			&rowsScan.CreTime, &rowsScan.ModId, &rowsScan.ModTime, &rowsScan.ModTs)

		if err != nil {
			return nil, err
		}

		rowsScanArr = append(rowsScanArr, &rowsScan)
	}

	return rowsScanArr, nil
}
