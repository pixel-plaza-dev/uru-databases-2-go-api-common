package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

type (
	// Handler interface
	Handler interface {
		CreateAuthenticatedEndpoint(mapper *pbtypesrest.Mapper, handler gin.HandlerFunc) (
			string,
			gin.HandlerFunc,
			gin.HandlerFunc,
		)
		CreateUnauthenticatedEndpoint(mapper *pbtypesrest.Mapper, handler gin.HandlerFunc) (string, gin.HandlerFunc)
	}

	// DefaultHandler struct
	DefaultHandler struct {
		authentication    auth.Authentication
		grpcInterceptions *map[grpc.Method]grpc.Interception
	}
)

// NewDefaultHandler creates a new default response handler
func NewDefaultHandler(
	authentication auth.Authentication,
	grpcInterceptions *map[grpc.Method]grpc.Interception,
) *DefaultHandler {
	return &DefaultHandler{authentication: authentication, grpcInterceptions: grpcInterceptions}
}

// CreateAuthenticatedEndpoint creates the authenticated endpoint
func (d *DefaultHandler) CreateAuthenticatedEndpoint(mapper *pbtypesrest.Mapper, handler gin.HandlerFunc) (
	string,
	gin.HandlerFunc,
	gin.HandlerFunc,
) {
	// Create the endpoint
	return mapper.Path(), d.authentication.Authenticate(
		mapper,
		d.grpcInterceptions,
	), handler
}

// CreateUnauthenticatedEndpoint creates the unauthenticated endpoint
func (d *DefaultHandler) CreateUnauthenticatedEndpoint(mapper *pbtypesrest.Mapper, handler gin.HandlerFunc) (
	string,
	gin.HandlerFunc,
) {
	// Create the endpoint
	return mapper.Path(), handler
}
