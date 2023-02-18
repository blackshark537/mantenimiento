package commands

import (
	"errors"

	entities "github.com/blackshark537/mantenimiento/src/core/Entities"
	left "github.com/blackshark537/mantenimiento/src/core/Ports/Left"
	actions "github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/fatih/color"
)

func Exec(ofType int, payload []byte) error {
	switch ofType {
	case actions.ServerStart:
		return left.Serve(string(payload[:]))

	// Area
	case actions.ListAreaType:
		e := entities.AreaType{}
		return e.List(payload)

	case actions.CreateAreaType:
		e := entities.Entity[entities.AreaType]{}
		e.Data = payload
		return e.Create()

	case actions.ClearAreaTypes:
		e := entities.Entity[entities.AreaType]{}
		return e.Clear("area_types")

		// Area
	case actions.ListArea:
		e := entities.Area{}
		return e.List(payload)

	case actions.CreateArea:
		e := entities.Entity[entities.Area]{}
		e.Data = payload
		return e.Create()

	case actions.ClearAreas:
		e := entities.Entity[entities.Area]{}
		return e.Clear("areas")

		// EquipoType
	case actions.ListEquipoType:
		e := entities.EquipoType{}
		return e.List(payload)

	case actions.CreateEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		e.Data = payload
		return e.Create()

	case actions.ClearEquipoType:
		e := entities.Entity[entities.EquipoType]{}
		return e.Clear("equipo_types")

		// Equipo
	case actions.ListEquipo:
		e := entities.Equipo{}
		return e.List(payload)

	case actions.CreateEquipo:
		e := entities.Entity[entities.Equipo]{}
		e.Data = payload
		return e.Create()

	case actions.ClearEquipos:
		e := entities.Entity[entities.Equipo]{}
		return e.Clear("equipos")

		// ComponenteType
	case actions.ListComponenteType:
		e := entities.ComponenteType{}
		return e.List(payload)

	case actions.CreateComponenteType:
		e := entities.Entity[entities.ComponenteType]{}
		e.Data = payload
		return e.Create()

	case actions.ClearCompoenteTypes:
		e := entities.Entity[entities.ComponenteType]{}
		return e.Clear("componente_types")

		// Compoente
	case actions.ListComponente:
		e := entities.Componente{}
		return e.List(payload)

	case actions.CreateComponente:
		e := entities.Entity[entities.Componente]{}
		e.Data = payload
		return e.Create()

	case actions.ClearComponentes:
		e := entities.Entity[entities.Componente]{}
		return e.Clear("componentes")
	default:
		return errors.New(color.RedString("Error: Command not implemented"))
	}
}
