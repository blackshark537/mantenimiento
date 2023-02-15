package entities

import (
	"encoding/json"
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Componente struct {
	Id             uint           `json:"id" gorm:"primary_key"`
	EquipoId       uint           `json:"equipo_id" gorm:"not null"`
	Marca          string         `json:"marca" gorm:"index; not null"`
	Modelo         string         `json:"modelo" gorm:"index; not null"`
	Serie          string         `json:"serie" gorm:"index; not null"`
	Descripcion    string         `json:"descripcion"`
	VidaUtil       int            `json:"vida_util" gorm:"not null"`
	ComponenteType ComponenteType `json:"componente_type" gorm:"foreignKey:ComponenteId; not null; constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	Owner          string         `json:"owner" gorm:"not null"`
}

func (cmp *Componente) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(cmp) == true {
		return
	}
	m.CreateTable(cmp)
}

func (comp *Componente) List(filter []byte) error {
	t := new(Entity[Componente])
	json.Unmarshal(filter, comp)
	err := t.GetAll().Where(comp).Preload("ComponenteType").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Componente", "Marca", "Modelo", "Descripcion", "VidaUtil (hr)")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.ComponenteType.Nombre, el.Marca, el.Modelo, el.Descripcion, el.VidaUtil)
	}
	tbl.Print()

	return err
}
