package comet

type IRouter interface{
	PreHandle(request IRequest)
	Handle(request IRequest)
	AfterHandle(request IRequest)
}

type BaseRouter struct {}

func (br *BaseRouter)PreHandle(req IRequest){}
func (br *BaseRouter)Handle(req IRequest){}
func (br *BaseRouter)AfterHandle(req IRequest){}


