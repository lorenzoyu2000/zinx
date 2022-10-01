package ziface

/*
	可自定义的对数据处理方法
*/
type IRouter interface {
	// 处理业务之前的方法
	PreHandle(request IRequest)
	// 处理业务的主方法
	Handle(request IRequest)
	// 处理业务之后的方法
	PostHandle(request IRequest)
}
