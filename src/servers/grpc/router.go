package grpc

import (
	"google.golang.org/grpc"

	"github.com/MultiBanker/broker/src/manager"
)

// setupResources монтирует необходимые grpc-ресурсы для
// обработки клиентских запросов
func Routing(server *grpc.Server, man manager.Abstractor) {

}
