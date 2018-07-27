package gonf

type ReaderMock struct {
	err error
}

func NewReaderMockError(e error) *ReaderMock{
	return &ReaderMock{err:e}
}

func (r *ReaderMock) Read(path string) ([]byte, error){

	if r.err != nil{
		return nil, r.err
	}

	return nil, nil
}