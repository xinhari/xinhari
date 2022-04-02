package xinhari

import (
	"xinhari.com/xinhari/client"
	"xinhari.com/xinhari/debug/trace"
	"xinhari.com/xinhari/server"
	"xinhari.com/xinhari/store"

	// set defaults
	gcli "xinhari.com/xinhari/client/grpc"
	memTrace "xinhari.com/xinhari/debug/trace/memory"
	gsrv "xinhari.com/xinhari/server/grpc"
	memoryStore "xinhari.com/xinhari/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
