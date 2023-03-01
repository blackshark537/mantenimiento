package entities

import (
	"fmt"
	"time"

	rightPort "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Suplidor struct {
	Id          int          `json:"id" gorm:"primary_key"`
	Nombre      string       `json:"nombre" gorm:"index; not null"`
	Direccion   string       `json:"direccion" gorm:"not null"`
	Contactos   []Contacto   `json:"contactos" gorm:"foreignKey:SuplidorId; constraint:OnDelete:CASCADE"`
	Componentes []Componente `json:"componentes" gorm:"foreignKey:SuplidorId; constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	Uid         string       `json:"uid,omitempty" gorm:"index; not null"`
}

func (spl *Suplidor) Migrate() {
	db := rightPort.GetDB()
	m := db.Migrator()
	if m.HasTable(spl) == true {
		return
	}
	m.CreateTable(spl)
}

func (cto *Suplidor) List(filter []byte) error {
	t := Entity[Suplidor]{
		Data: filter,
	}
	err := t.GetAll().Preload("Contactos").Limit(25).Find(&t.Entities).Error
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Nombre", "Direccion", "Telefono")
	fmt.Printf("%s %v Items\n", color.MagentaString("[Results]:"), len(t.Entities))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, el := range t.Entities {
		var phone string = ""
		if len(el.Contactos) > 0 {
			phone = el.Contactos[0].Numero
		}
		tbl.AddRow(el.Id, el.Nombre, el.Direccion, phone)
	}
	tbl.Print()
	return err
}
