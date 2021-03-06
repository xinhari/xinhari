// Package service encapsulates the client, server and other interfaces to provide a complete xinhari service.
package service

import (
	"xinhari.com/xinhari/client"
	"xinhari.com/xinhari/server"
)

// Service is an interface for a xinhari service
type Service interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// Run the service
	Run() error
	// The service implementation
	String() string
}
