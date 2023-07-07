package sistema

type Sistema interface {
	//
	InicializarSistema(string, string) string
	//
	Centralidad(string) string
	//
	CaminoEscalas(string) string
	//
	CaminoMas(string) string
	//
	NuevaAerolinea(string) string
	//
	Itenerario(string) string
	//
	ExportarKML(string) string
}
