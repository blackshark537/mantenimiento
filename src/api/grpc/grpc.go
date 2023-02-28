package grpc

import (
	"fmt"
	"log"

	"github.com/blackshark537/mantenimiento/src/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Conn() (*grpc.ClientConn, error) {
	env := config.GetEnv()
	host := fmt.Sprintf("%s:%s", env.GRPC_HOST, env.GRPC_PORT)
	rpc, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("[Error]: ", err)
	}
	return rpc, err
}
