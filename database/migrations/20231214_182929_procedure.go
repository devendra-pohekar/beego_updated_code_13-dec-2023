package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Procedure_20231214_182929 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Procedure_20231214_182929{}
	m.Created = "20231214_182929"

	migration.Register("Procedure_20231214_182929", m)
}

// Run the migrations
func (m *Procedure_20231214_182929) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Procedure_20231214_182929) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
