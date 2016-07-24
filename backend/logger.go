package backend

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type GAELogger struct {
}

func (*GAELogger) Debugf(ctx context.Context, format string, v ...interface{}) {
	log.Debugf(ctx, format, v)
}

func (*GAELogger) Infof(ctx context.Context, format string, v ...interface{}) {
	log.Infof(ctx, format, v)
}

func (*GAELogger) Errorf(ctx context.Context, format string, v ...interface{}) {
	log.Errorf(ctx, format, v)
}
