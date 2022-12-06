package cont

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello, world"
	srv := Server(&StubStore{data})
	request := httptest.NewRequest(http.MethodGet)
}
