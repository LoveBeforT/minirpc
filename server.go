package minirpc

type Server interface {
}

type serverImpl struct {
}

func NewServer() Server {
	return &serverImpl{}
}
