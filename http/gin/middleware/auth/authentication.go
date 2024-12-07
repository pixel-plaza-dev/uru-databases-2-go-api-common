package auth

import (
	"github.com/gin-gonic/gin"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Authentication interface
type Authentication interface {
	Authenticate(
		mapper *pbtypesrest.Mapper,
		grpcInterceptions *map[pbtypesgrpc.Method]pbtypesgrpc.Interception,
	) gin.HandlerFunc
}
