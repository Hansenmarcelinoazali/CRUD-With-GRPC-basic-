package router

import (
	"fmt"
	"latihan/configs"

	// "latihan/proto/v1/books"
	"log"
	"net"

	"latihan/api/v1/books"
	hlpb "latihan/proto/v1/books"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGRPCServer creates the grpc server serve mux.
func NewGRPCServer(configs *configs.Configs, logger *logrus.Logger) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.Config.Server.Grpc.Port))
	if err != nil {
		return err
	}

	// register grpc service server
	grpcServer := grpc.NewServer()

	hlpb.RegisterBookServiceServer(grpcServer, books.New(configs, logger))

	// add reflection service
	reflection.Register(grpcServer)

	// running gRPC server
	log.Println("[SERVER] GRPC server is ready")

	grpcServer.Serve(lis)

	return nil
}
