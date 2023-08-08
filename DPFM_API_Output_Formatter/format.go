package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToPartnerFunction(rows *sql.Rows) (*PartnerFunction, error) {
	defer rows.Close()
	partnerFunction := PartnerFunction{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&partnerFunction.partnerFunction,
			&partnerFunction.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &partnerFunction, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partnerFunction, nil
	}

	return &partnerFunction, nil
}
