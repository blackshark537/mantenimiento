package commands

import (
	"encoding/json"
	"errors"

	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	left "github.com/blackshark537/mantenimiento/src/core/Ports/Left"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/fatih/color"
)

func Exec(ofType int, payload []byte) error {
	switch ofType {
	case actions.ServerStart:
		return left.Serve()

		// Users
	case actions.ListUser:
		e := entities.User{}
		return e.List(payload)
	case actions.CreateUser:
		e := entities.Entity[entities.User]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// Empresas
	case actions.ListEmpresa:
		e := entities.Empresa{}
		return e.List(payload)
	case actions.CreateEmpresa:
		e := entities.Entity[entities.Empresa]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// Contacto
	case actions.ListContacto:
		e := entities.Contacto{}
		return e.List(payload)
	case actions.CreateContacto:
		e := entities.Entity[entities.Contacto]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// Geoposicion
	case actions.ListGeo:
		e := entities.Geoposicion{}
		return e.List(payload)
	case actions.CreateGeo:
		e := entities.Entity[entities.Geoposicion]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	// Area
	case actions.ListAreaType:
		e := entities.AreaType{}
		return e.List(payload)
	case actions.CreateAreaType:
		e := entities.Entity[entities.AreaType]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// Area
	case actions.ListArea:
		e := entities.Area{}
		return e.List(payload)
	case actions.CreateArea:
		e := entities.Entity[entities.Area]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// EquipoType
	case actions.ListEquipoType:
		e := entities.EquipoType{}
		return e.List(payload)
	case actions.CreateEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// Equipo
	case actions.ListEquipo:
		e := entities.Equipo{}
		return e.List(payload)
	case actions.CreateEquipo:
		e := entities.Entity[entities.Equipo]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// ComponenteType
	case actions.ListComponenteType:
		e := entities.ComponenteType{}
		return e.List(payload)
	case actions.CreateComponenteType:
		e := entities.Entity[entities.ComponenteType]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

		// Compoente
	case actions.ListComponente:
		e := entities.Componente{}
		return e.List(payload)
	case actions.CreateComponente:
		e := entities.Entity[entities.Componente]{}
		json.Unmarshal(payload, &e.Entity)
		return e.Create()

	default:
		return errors.New(color.RedString("Error: Command not implemented"))
	}
}
