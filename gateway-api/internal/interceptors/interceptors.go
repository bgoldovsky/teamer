package interceptors

import (
	"context"
	"time"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggingInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)

	logger.Log.
		WithFields(logrus.Fields{"method": method, "duration": time.Since(start), "error": err}).
		Infof("invoked gRPC method")

	return err
}
