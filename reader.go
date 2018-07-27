package gonf

type Reader interface {
	Read(path string) ([]byte, error)
}