package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Procedure_20231213_190540 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Procedure_20231213_190540{}
	m.Created = "20231213_190540"

	migration.Register("Procedure_20231213_190540", m)
}

// Run the migrations
func (m *Procedure_20231213_190540) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TRIGGER after_delete_hpst
AFTER DELETE ON home_pages_setting_table
FOR EACH ROW EXECUTE FUNCTION move_to_backup_hpst();`)

}

// Reverse the migrations
func (m *Procedure_20231213_190540) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
