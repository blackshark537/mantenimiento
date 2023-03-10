package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type ComponenteType struct {
	Id          int          `json:"id" gorm:"primary_key"`
	Nombre      string       `json:"nombre" gorm:"index; not null"`
	EquipoId    uint         `json:"equipo_id" gorm:"not null"`
	Componentes []Componente `json:"componentes" gorm:"foreignKey:ComponenteTypeId; not null; constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	Uid         string       `json:"uid,omitempty" gorm:"index; not null"`
}

func (cmp *ComponenteType) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(cmp) == true {
		return
	}
	m.CreateTable(cmp)
}

func (cmp *ComponenteType) List(filter []byte) error {
	t := new(Entity[ComponenteType])
	t.Data = filter
	err := t.GetAll().Preload("Componentes").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Nombre", "Camtidad")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.Nombre, len(el.Componentes))
	}
	tbl.Print()
	return err
}
