package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type AreaType struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Nombre    string    `json:"nombre" gorm:"index; not null"`
	EmpresaId uint      `json:"empresa_id" gorm:"not null"`
	Areas     []Area    `json:"areas" gorm:"foreignKey:AreaTypeId; constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Uid       string    `json:"uid,omitempty" gorm:"index; not null"`
}

func (atp *AreaType) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(atp) == true {
		return
	}
	m.CreateTable(atp)
}

func (area *AreaType) List(filter []byte) error {
	t := Entity[AreaType]{}
	t.Data = filter
	err := t.GetAll().Preload("Areas").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Nombre", "Camtidad")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.Nombre, len(el.Areas))
	}
	tbl.Print()
	return err
}
