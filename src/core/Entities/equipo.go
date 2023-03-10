package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Equipo struct {
	Id           int              `json:"id" gorm:"primary_key"`
	AreaId       uint             `json:"area_id" gorm:"not null"`
	Marca        string           `json:"marca" gorm:"index; not null"`
	Modelo       string           `json:"modelo" gorm:"index; not null"`
	Serie        string           `json:"serie" gorm:"index; not null"`
	Descripcion  string           `json:"descripcion"`
	EquipoTypeId uint             `json:"equipo_type_id" gorm:"not null"`
	Componentes  []ComponenteType `json:"componentes" gorm:"foreignKey:EquipoId; constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time        `json:"created_at" gorm:"autoCreateTime"`
	Uid          string           `json:"uid,omitempty" gorm:"index; not null"`
}

func (eqp *Equipo) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(eqp) == true {
		return
	}
	m.CreateTable(eqp)
}

func (equipo *Equipo) List(filter []byte) error {
	t := new(Entity[Equipo])
	t.Data = filter
	err := t.GetAll().Preload("Componentes").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Equipo", "Marca", "Modelo", "Serie", "No.Componentes")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.EquipoTypeId, el.Marca, el.Modelo, el.Serie, len(el.Componentes))
	}
	tbl.Print()

	return err
}
