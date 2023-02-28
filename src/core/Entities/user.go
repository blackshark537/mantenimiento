package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Nombre    string    `json:"nombre" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"not null"`
	Uid       string    `json:"uid,omitempty" gorm:"index; not null; unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (usr *User) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(usr) == true {
		return
	}
	m.CreateIndex(usr, "nombre")
	m.CreateIndex(usr, "email")
	m.CreateTable(usr)
}

func (user *User) List(filter []byte) error {
	t := new(Entity[User])
	t.Data = filter
	err := t.GetAll().Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Nombre", "Email", "CreatedAt")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.Nombre, el.Email, el.CreatedAt)
	}
	tbl.Print()
	return err
}
