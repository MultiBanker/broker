package grpc

import (
	"google.golang.org/grpc"

	"github.com/MultiBanker/broker/src/manager"
)

// Routing монтирует необходимые grpc-ресурсы для
// обработки клиентских запросов
func Routing(server *grpc.Server, man manager.Managers) {

}
