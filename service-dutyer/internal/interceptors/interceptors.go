package interceptors

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (h interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Log.
				WithField("recovered", r).
				WithField("stackTrace", string(debug.Stack())).
				Error("panic recovered")
			err = errors.New("internal service error")
		}
	}()

	start := time.Now()
	h, err = handler(ctx, req)

	logger.Log.
		WithFields(logrus.Fields{"method": info.FullMethod, "duration": time.Since(start), "error": err}).
		Infof("gRPC request")

	return
}
