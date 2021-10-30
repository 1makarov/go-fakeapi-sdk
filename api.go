package fakeapi

type api struct {
	caller *transport
}

func newAPI(caller *transport) *api {
	return &api{caller}
}

func (api *api) call(url, raw, method string, body interface{}) ([]byte, int, error) {
	return api.caller.Call(url, raw, method, body)
}
