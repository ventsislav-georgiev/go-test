package main

type client struct {
}

func (*client) GetResponseModes() {
}

type Client interface {
}

type ResponseModeClient interface {
	GetResponseModes()
}

func main() {
	var client Client = &client{}
	responseModeClient, ok := client.(ResponseModeClient)
	if !ok {
		print("fail", client)
	} else {
		print("ok", responseModeClient)
	}
}
