package entities

import (
	"encoding/json"
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type EquipoType struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Nombre    string    `json:"nombre" gorm:"index; not null"`
	EquipoId  uint      `json:"equipo_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Owner     string    `json:"owner" gorm:"not null"`
}

func (equipo *EquipoType) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(equipo) == true {
		return
	}
	m.CreateTable(equipo)
}

func (equipo *EquipoType) List(filter []byte) error {
	t := new(Entity[EquipoType])
	json.Unmarshal(filter, equipo)
	err := t.GetAll().Where(equipo).Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Nombre", "EquipoId")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.Nombre, el.EquipoId)
	}
	tbl.Print()
	return err
}
