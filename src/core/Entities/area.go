package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Area struct {
	Id          int          `json:"id" gorm:"primary_key"`
	Descripcion string       `json:"descripcion" gorm:"not null"`
	Capacidad   int          `json:"capacidad" gorm:"index; not null"`
	Largo       float32      `json:"largo" gorm:"not null"`
	Ancho       float32      `json:"ancho" gorm:"not null"`
	AreaTypeId  uint         `json:"area_type_id" gorm:"not null"`
	Equipos     []EquipoType `json:"equipos" gorm:"foreignKey:AreaId; constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	Uid         string       `json:"uid,omitempty" gorm:"index; not null"`
}

func (area *Area) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(area) == true {
		return
	}
	m.CreateTable(area)
}

func (area *Area) List(filter []byte) error {
	t := new(Entity[Area])
	t.Data = filter
	err := t.GetAll().Preload("Equipos").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Area", "Descipcion", "Capacidad", "No.Equipos", "Ancho", "Largo")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {

		tbl.AddRow(el.Id, el.AreaTypeId, el.Descripcion, el.Capacidad, len(el.Equipos), el.Ancho, el.Largo)
	}
	tbl.Print()
	return err
}
