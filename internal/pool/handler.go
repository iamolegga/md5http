package pool

//go:generate mockgen -destination=./mock/handler.go -package=mock . Handler
type Handler interface {
	Handle(string) error
}
