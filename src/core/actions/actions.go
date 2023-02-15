package actions

const (
	ServerStart = iota

	GetUser = iota
	ListUser
	CreateUser
	UpdateUser
	ChangePswd

	GetEmpresas = iota
	GetEmpresa
	ListEmpresa
	CreateEmpresa
	UpdateEmpresa
	DeleteEmpresa

	GetAllContactos = iota
	GetContacto
	ListContacto
	CreateContacto
	UpdateContacto
	DeleteContacto

	GetAllGeo = iota
	GetGeo
	ListGeo
	CreateGeo
	UpdateGeo
	DeleteGeo

	GetAllAreaType = iota
	GetAreaType
	ListAreaType
	CreateAreaType
	UpdateAreaType
	DeleteAreaType

	GetAllAreas = iota
	GetArea
	ListArea
	CreateArea
	UpdateArea
	DeleteArea

	GetAllEquipoType = iota
	GetEquipoType
	ListEquipoType
	CreateEquipoType
	UpdateEquipoType
	DeleteEquipoType

	GetAllEquipos = iota
	GetEquipo
	ListEquipo
	CreateEquipo
	UpdateEquipo
	DeleteEquipo

	GetAllComponentesType = iota
	GetComponenteType
	ListComponenteType
	CreateComponenteType
	UpdateComponenteType
	DeleteComponenteType

	GetAllComponentes = iota
	GetComponente
	ListComponente
	CreateComponente
	UpdateComponente
	DeleteComponente
)
