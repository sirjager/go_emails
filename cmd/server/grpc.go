package server

import (
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/sirjager/go_emails/cfg"
	"github.com/sirjager/go_emails/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	rpc "github.com/sirjager/rpcs/emails/go"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func RunGRPCServer(srvr *server.RPCEmailsServer, logger zerolog.Logger, config cfg.Config, errs chan error) {
	listener, err := net.Listen("tcp", ":"+config.Grpc)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to listen grpc tcp server")
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				GRPCLogger(logger),
				grpc_prometheus.UnaryServerInterceptor,
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				GRPCStreamLogger(logger),
				grpc_prometheus.StreamServerInterceptor,
			),
		),

		grpc.MaxRecvMsgSize(1024*1024), // bytes * Kilobytes * Megabytes
	)

	rpc.RegisterEmailsServer(grpcServer, srvr)

	reflection.Register(grpcServer)

	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())

	logger.Info().Msgf("started gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to server gRPC server")
	}
}
