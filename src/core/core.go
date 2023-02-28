package core

import (
	"errors"

	Commands "github.com/blackshark537/mantenimiento/src/core/Commands"
	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	left_port "github.com/blackshark537/mantenimiento/src/core/Ports/Left"
	right_port "github.com/blackshark537/mantenimiento/src/core/Ports/Right"
	Query "github.com/blackshark537/mantenimiento/src/core/Query"
	"gorm.io/gorm"
)

type API interface {
	ForRoot(port string) error
}

type Action struct {
	OfType  int
	Payload []byte
	Uid     string
}

func InjectDatabase(db *gorm.DB) {
	right_port.InjectDB(db)
}

func InjectApi(api API) {
	left_port.InjectApi(api)
}

func (a *Action) Exec() error {
	if a.Uid == "" {
		return errors.New("Error field Uid is required")
	}
	return Commands.Exec(a.OfType, a.Payload, a.Uid)
}

func (a *Action) Query() (interface{}, error) {
	if a.Uid == "" {
		return nil, errors.New("Error field Uid is required")
	}
	return Query.Exec(a.OfType, a.Payload, a.Uid)
}

func Migrate() {
	/* usr := entities.User{}
	usr.Migrate()

	emp := entities.Empresa{}
	emp.Migrate()

	geo := entities.Geoposicion{}
	geo.Migrate()

	cto := entities.Contacto{}
	cto.Migrate() */

	areaType := entities.AreaType{}
	areaType.Migrate()

	area := entities.Area{}
	area.Migrate()

	eqpType := entities.EquipoType{}
	eqpType.Migrate()

	eqp := entities.Equipo{}
	eqp.Migrate()

	cmpType := entities.ComponenteType{}
	cmpType.Migrate()

	cmp := entities.Componente{}
	cmp.Migrate()

	/* alm := entities.Alimento{}
	alm.Migrate()

	ing := entities.Ingrediente{}
	ing.Migrate()

	spl := entities.Suplidor{}
	spl.Migrate() */
}
