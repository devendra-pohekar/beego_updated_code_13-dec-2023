package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Procedure_20231214_144757 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Procedure_20231214_144757{}
	m.Created = "20231214_144757"

	migration.Register("Procedure_20231214_144757", m)
}

// Run the migrations
func (m *Procedure_20231214_144757) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE OR REPLACE FUNCTION insert_update_delete()
RETURNS TRIGGER AS $$
BEGIN
       IF TG_OP = 'INSERT' THEN
        INSERT INTO backup_after_delete_hpst(page_setting_id,
        section,
        data_type,
        unique_code,
        setting_data,
        created_date,
        updated_date,
        created_by,
        updated_by,
        sample)
        VALUES(NEW.page_setting_id,
        NEW.section,
        NEW.data_type,
        NEW.unique_code,
        NEW.setting_data,
        NEW.created_date,
        NEW.updated_date,
        NEW.created_by,
        NEW.updated_by,
        NEW.sample);
        RAISE NOTICE 'A new row was inserted with id % and copied to backup_after_delete_hpst with source_id %', NEW.page_setting_id, NEW.page_setting_id;  
	END IF;	 
END;
$$ LANGUAGE plpgsql;`)
}

// Reverse the migrations
func (m *Procedure_20231214_144757) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
