package rw

type Reader interface {
	RwRead(p *string) (n int, err error)
}
