package micro

import (
	"github.com/ebelanja/go-micro/client"
	"github.com/ebelanja/go-micro/debug/trace"
	"github.com/ebelanja/go-micro/server"
	"github.com/ebelanja/go-micro/store"

	// set defaults
	gcli "github.com/ebelanja/go-micro/client/grpc"
	memTrace "github.com/ebelanja/go-micro/debug/trace/memory"
	gsrv "github.com/ebelanja/go-micro/server/grpc"
	memoryStore "github.com/ebelanja/go-micro/store/memory"
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
