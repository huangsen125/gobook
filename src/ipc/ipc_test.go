package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{method, params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1 := client1.Call("ECHO", "From Client1")
	resp2 := client2.Call("ECHO", "From Client2")

	fmt.Println(resp1.Body, resp1.Code)
	if resp1.Body != "From Client1" || resp2.Body != "From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}
}
