package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Procedure_20231214_145503 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Procedure_20231214_145503{}
	m.Created = "20231214_145503"

	migration.Register("Procedure_20231214_145503", m)
}

// Run the migrations
func (m *Procedure_20231214_145503) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TRIGGER insert_update_delete_trigger
AFTER INSERT OR UPDATE OR DELETE ON home_pages_setting_table
FOR EACH ROW EXECUTE FUNCTION insert_update_delete();
`)
}

// Reverse the migrations
func (m *Procedure_20231214_145503) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
