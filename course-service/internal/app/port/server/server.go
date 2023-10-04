package server

type AppServer interface {
	Start(addr string) error
	Stop() error
}
