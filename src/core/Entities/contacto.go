package entities

import (
	"fmt"
	"log"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const (
	CASA    = "CASA"
	OFICINA = "OFICINA"
	CELULAR = "CELULAR"
	OTRO    = "OTRO"
)

type Contacto struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	SuplidorId uint      `json:"suplidor_id" gorm:"not null"`
	Numero     string    `json:"numero" gorm:"index; not null"`
	Tipo       string    `json:"tipo" gorm:"type:phone_type; not null; default:CASA"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Uid        string    `json:"uid,omitempty" gorm:"index; not null"`
}

func (cto *Contacto) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(cto) == true {
		return
	}
	cto.CreateTypes()
	m.CreateTable(cto)
}

func (cto *Contacto) CreateTypes() {
	db := rightPort.GetDB()
	err := db.Exec("CREATE TYPE phone_type AS ENUM ('CASA','OFICINA', 'CELULAR', 'OTRO');").Error
	if err != nil {
		log.Fatal(err)
	}
}

func (cto *Contacto) List(filter []byte) error {
	t := new(Entity[Contacto])
	t.Data = filter
	err := t.GetAll().Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Suplidor Id", "Teléfono", "Tipo")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.SuplidorId, el.Numero, el.Tipo)
	}
	tbl.Print()
	return err
}
