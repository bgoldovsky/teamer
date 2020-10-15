package interceptors

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Log.
				WithField("recovered", r).
				WithField("stackTrace", string(debug.Stack())).
				Error("panic recovered")
		}
	}()

	start := time.Now()
	h, err := handler(ctx, req)

	logger.Log.
		WithFields(logrus.Fields{"method": info.FullMethod, "duration": time.Since(start), "error": err}).
		Infof("gRPC request")

	return h, err
}
