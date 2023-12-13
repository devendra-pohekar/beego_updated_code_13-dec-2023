package models

import (
	"github.com/beego/beego/v2/client/orm"
)

func FormateDateFromTable(tableName, createdDateColumn, updatedDateColumn string) ([]orm.Params, error) {
	o := orm.NewOrm()
	var results []orm.Params
	_, err := o.Raw(`
		SELECT 
			TO_CHAR(created_date, 'DD-Mon-YYYY HH12:MI:SS AM') AS formatted_created_date,
			TO_CHAR(updated_date, 'DD-Mon-YYYY HH12:MI:SS AM') AS formatted_updated_date,
			TO_CHAR(created_date, 'HH12:MI:SS AM') AS created_time,
			TO_CHAR(updated_date, 'HH12:MI:SS AM') AS updated_time
		FROM 
			` + tableName).Values(&results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

// formattedDates, err := models.FormateDateFromTable("home_pages_setting_table", "created_date", "updated_date")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Use or print the formatted results as needed
// 	for _, row := range formattedDates {
// 		fmt.Printf("Formatted Created Date: %s\n", row["formatted_created_date"])
// 		fmt.Printf("Formatted Updated Date: %s\n", row["formatted_updated_date"])
// 		fmt.Printf("Created Time: %s\n", row["created_time"])
// 		fmt.Printf("Updated Time: %s\n", row["updated_time"])
// 	}
