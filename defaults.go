package micro

import (
	"xinhari.com/client"
	"xinhari.com/debug/trace"
	"xinhari.com/server"
	"xinhari.com/store"

	// set defaults
	gcli "xinhari.com/client/grpc"
	memTrace "xinhari.com/debug/trace/memory"
	gsrv "xinhari.com/server/grpc"
	memoryStore "xinhari.com/store/memory"
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
