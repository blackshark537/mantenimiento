package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Geoposicion struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	EmpresaId uint      `json:"empresa_id" gorm:"index; not null"`
	Lat       float32   `json:"lat" gorm:"not null"`
	Lng       float32   `json:"lng" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Owner     string    `json:"owner" gorm:"not null"`
}

func (geo *Geoposicion) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(geo) == true {
		return
	}
	m.CreateTable(geo)
}

func (geo *Geoposicion) List(filter []byte) error {
	t := new(Entity[Geoposicion])
	t.Data = filter
	err := t.GetAll().Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "EmpresaId", "Lat", "Lng")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		tbl.AddRow(el.Id, el.EmpresaId, el.Lat, el.Lng)
	}
	tbl.Print()

	return err
}
