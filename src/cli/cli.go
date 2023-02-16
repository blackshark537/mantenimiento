package cli

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/blackshark537/mantenimiento/src/config"
	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"

	"github.com/fatih/color"

	"github.com/urfave/cli/v2"
)

var instance = color.MagentaString("[CLI]:")

var (
	table  string
	filter string
	port   string
	data   string
)

const (
	empresas          = "empresas"
	contactos         = "contactos"
	geopoints         = "geopoints"
	areas_types       = "areas_types"
	areas             = "areas"
	equipos_types     = "equipos_types"
	equipos           = "equipos"
	componentes_types = "componentes_types"
	componentes       = "componentes"
	users             = "users"
)

var tables = []string{
	users,
	empresas,
	contactos,
	geopoints,
	areas_types,
	areas,
	equipos,
	equipos_types,
	componentes_types,
	componentes,
}

var Commands []*cli.Command = []*cli.Command{
	{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Start the server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       "3000",
				Usage:       "Server port listener",
				Destination: &port,
			},
		},
		Action: serverStart,
	},
	{
		Name:    "migrate",
		Aliases: []string{"m"},
		Usage:   "Execute Auto Migration",
		Action:  migrate,
	},
	{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "Show a list of a given table",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "from",
				Usage:       "Database Table",
				Aliases:     []string{"f"},
				Destination: &table,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "where",
				Usage:       `Apply Filters Ex: '{"key": "value"}'`,
				Aliases:     []string{"w"},
				Destination: &filter,
				Value:       "",
			},
		},
		Action: listTable,
	},
	{
		Name:    "insert",
		Aliases: []string{"i"},
		Usage:   "Create a new register",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "into",
				Usage:       "Database Table",
				Aliases:     []string{"to"},
				Destination: &table,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "data",
				Usage:       "Data to insert",
				Aliases:     []string{"d"},
				Destination: &data,
				Value:       "",
			},
		},
		Action: createEntity,
	},
	{
		Name:    "clear",
		Aliases: []string{"clr"},
		Usage:   "Clear all registers from a given table",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "from",
				Usage:       "Database Table",
				Aliases:     []string{"f"},
				Destination: &table,
				Required:    true,
			},
		},
		Action: clearEntity,
	},
}

func New() *cli.App {

	AppName := color.MagentaString(fmt.Sprintf("Dataprod - Mantenimiento"))
	AppArch := color.MagentaString(fmt.Sprintf("Core Arch: %v", runtime.GOARCH))
	AppDesc := color.MagentaString("A simple CLI program to manage your Mantenimiento system")

	website := color.CyanString("WEBSITE: https://dataprod.cloud")
	support := color.GreenString("SUPPORT: support@dataprod.cloud")
	cli.AppHelpTemplate = fmt.Sprintf(`
	                       _____________________________________________________
                      |                                                     |
             _______  |                                                     |
            / _____ | |           %s                  |
           / /(__) || |           %s                          |
  ________/ / |OO| || |                                                     |
 |         |-------|| |                                                     |
(|         |     -.|| |_______________________                              |
 |  ____   \       ||_________||____________  |             ____      ____  |
/| / __ \   |______||     / __ \   / __ \   | |            / __ \    / __ \ |\
\|| /  \ |_______________| /  \ |_| /  \ |__| |___________| /  \ |__| /  \|_|/
   | () |                 | () |   | () |                  | () |    | () |
    \__/                   \__/     \__/                    \__/      \__/
	%s

	%s
	
	%s
	
	`, AppName, AppArch, cli.AppHelpTemplate, website, support)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	return &cli.App{
		Name:     AppName,
		Version:  "v" + config.VERSION,
		Usage:    AppDesc,
		Commands: Commands,
	}
}

func migrate(ctx *cli.Context) error {
	Core.Migrate()
	return nil
}

func serverStart(ctx *cli.Context) error {
	action := Core.Action{}
	return action.Exec()
}

func listTable(ctx *cli.Context) error {
	action := Core.Action{}

	switch table {
	case users:
		action.OfType = actions.ListUser
		break
	case empresas:
		action.OfType = actions.ListEmpresa
		break
	case contactos:
		action.OfType = actions.ListContacto
		break
	case geopoints:
		action.OfType = actions.ListGeo
		break
	case areas_types:
		action.OfType = actions.ListAreaType
		break
	case areas:
		action.OfType = actions.ListArea
		break
	case equipos_types:
		action.OfType = actions.ListEquipoType
		break
	case equipos:
		action.OfType = actions.ListEquipo
		break
	case componentes_types:
		action.OfType = actions.ListComponenteType
		break
	case componentes:
		action.OfType = actions.ListComponente
		break
	default:
		return typesAvailable()
	}
	action.Payload = []byte(filter)
	return action.Exec()
}

func createEntity(cli *cli.Context) error {
	action := Core.Action{}
	switch table {
	case users:
		action.OfType = actions.CreateUser
		break
	case empresas:
		action.OfType = actions.CreateEmpresa
		break
	case contactos:
		action.OfType = actions.CreateContacto
		break
	case geopoints:
		action.OfType = actions.CreateGeo
		break
	case areas_types:
		action.OfType = actions.CreateAreaType
		break
	case areas:
		action.OfType = actions.CreateArea
		break
	case equipos_types:
		action.OfType = actions.CreateEquipoType
		break
	case equipos:
		action.OfType = actions.CreateEquipo
		break
	case componentes_types:
		action.OfType = actions.CreateComponenteType
		break
	case componentes:
		action.OfType = actions.CreateComponente
		break
	default:
		return typesAvailable()
	}
	action.Payload = []byte(data)
	err := action.Exec()
	if err == nil {
		fmt.Printf("%s creado con exito!\n", table)
	}
	return err
}

func clearEntity(ctx *cli.Context) error {
	action := Core.Action{}

	switch table {
	case users:
		action.OfType = actions.ClearUsers
		break
	case empresas:
		action.OfType = actions.ClearEmpresas
		break
	case contactos:
		action.OfType = actions.ClearContactos
		break
	case geopoints:
		action.OfType = actions.ClearGeo
		break
	case areas_types:
		action.OfType = actions.ClearAreaTypes
		break
	case areas:
		action.OfType = actions.ClearAreas
		break
	case equipos_types:
		action.OfType = actions.ClearEquipoType
		break
	case equipos:
		action.OfType = actions.ClearEquipos
		break
	case componentes_types:
		action.OfType = actions.ClearCompoenteTypes
		break
	case componentes:
		action.OfType = actions.ClearComponentes
		break
	default:
		return typesAvailable()
	}
	err := action.Exec()
	if err == nil {
		fmt.Printf("Todos los registros de la tabla %s, han sido borrados con exito!\n", table)
	}
	return err
}

func typesAvailable() error {
	var tablesAvailable = ""
	for i, el := range tables {
		tablesAvailable += fmt.Sprintf("%d - %v\n", i+1, el)
	}
	msg := fmt.Sprintf(
		"%s\n%v\n%v\n",
		color.RedString("Error: Table Not Available."),
		color.YellowString("Availables Tables Are:"),
		color.GreenString(fmt.Sprintf("%v", tablesAvailable)),
	)
	return errors.New(msg)
}
