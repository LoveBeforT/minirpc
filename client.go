package minirpc

type Client interface {
}

type clientImpl struct {
}

func NewClient() Client {
	return &clientImpl{}
}
