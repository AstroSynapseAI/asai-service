package rest

type Routes interface {
	LoadRoutes(router *Rest)
	LoadMiddlewares(router *Rest)
}

type RestHandler interface {
	Run()
	SetRouter(router *Rest)
	Create(ctx *Context)
	Read(ctx *Context)
	ReadAll(ctx *Context)
	Update(ctx *Context)
	Destroy(ctx *Context)
}

type HandlerFunc func(ctx *Context)

