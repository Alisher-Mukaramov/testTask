package server

type Endpoint struct {
	version string
	server  *Server
}

func NewEndpoint(version string, srv *Server) *Endpoint {
	return &Endpoint{
		version: version,
		server:  srv,
	}
}

func (e *Endpoint) Init() {
	//ver := fmt.Sprintf("/api/%s", e.version)
	//v := e.server.engine.Group(ver)

	// endpoints
	e.server.engine.PUT("/update", e.server.handler.Update)
	e.server.engine.GET("/state", e.server.handler.State)
	e.server.engine.GET("/get_names", e.server.handler.GetNames)

}
