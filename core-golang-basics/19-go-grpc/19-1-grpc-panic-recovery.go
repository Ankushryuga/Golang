package main

import (
  "context"
  "log"
  "net"

  "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
  healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
  healthsvc "google.golang.org/grpc/health"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

func panicRecoveryHandler(ctx context.Context, p any) (err error) {
  log.Printf("[GRPC PANIC] %v\n%s", p, recovery.CallStack(4))
  return status.Errorf(codes.Internal, "server panicked: %v", p)
}

func newGRPCServer() *grpc.Server {
  opts := []recovery.Option{
    recovery.WithRecoveryHandlerContext(panicRecoveryHandler),
  }
  return grpc.NewServer(
    grpc.UnaryInterceptor(recovery.UnaryServerInterceptor(opts...)),
    grpc.StreamInterceptor(recovery.StreamServerInterceptor(opts...)),
  )
}

type myServiceServer struct {
  // example service implementation
}

func (s *myServiceServer) PanickyCall(ctx context.Context, req *MyRequest) (*MyResponse, error) {
  panic("unexpected error in service")
}

func main() {
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("listen error: %v", err)
  }

  grpcServer := newGRPCServer()

  healthServer := healthsvc.NewServer()
  healthgrpc.RegisterHealthServer(grpcServer, healthServer)
  healthServer.SetServingStatus("", healthgrpc.HealthCheckResponse_SERVING)

  // Register your services
  RegisterMyServiceServer(grpcServer, &myServiceServer{})

  go func() {
    if err := grpcServer.Serve(lis); err != nil {
      log.Fatalf("serve error: %v", err)
    }
  }()
  log.Println("gRPC server started on :50051")
  select {} // block main
}
