package aeutils

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/appengine/log"
)

var trimprefix = ""

// Set trimprefix to the path to the source code directory, so that we only log the filename and not the full path.
func init() {
	_, path, _, _ := runtime.Caller(1)

	trimprefix = path[:strings.LastIndex(path, "/")+1]
}

func logctx(skip int) (file string, line int) {
	_, file, line, _ = runtime.Caller(skip + 1)
	file = strings.TrimPrefix(file, trimprefix)

	return
}

func LogDebugf(ctx context.Context, format string, a ...interface{}) {
	log.Debugf(ctx, format, a...)
}

func LogDebugfd(ctx context.Context, format string, a ...interface{}) {
	file, line := logctx(1)

	log.Debugf(ctx, "%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func LogDebugJSON(ctx context.Context, v interface{}) {
	if b, err := json.MarshalIndent(v, "", "\t"); err != nil {
		log.Debugf(ctx, "json.MarshalIndent(): err=%v", err)
	} else {
		log.Debugf(ctx, string(b))
	}
}

func LogInfof(ctx context.Context, format string, a ...interface{}) {
	log.Infof(ctx, format, a...)
}

func LogInfofd(ctx context.Context, format string, a ...interface{}) {
	file, line := logctx(1)

	log.Infof(ctx, "%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func LogWarningf(ctx context.Context, format string, a ...interface{}) {
	log.Warningf(ctx, format, a...)
}

func LogWarningfd(ctx context.Context, format string, a ...interface{}) {
	file, line := logctx(1)

	log.Warningf(ctx, "%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func LogErrorf(ctx context.Context, format string, a ...interface{}) {
	log.Errorf(ctx, format, a...)
}

func LogErrorfd(ctx context.Context, format string, a ...interface{}) {
	file, line := logctx(1)

	log.Errorf(ctx, "%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func LogCriticalf(ctx context.Context, format string, a ...interface{}) {
	log.Criticalf(ctx, format, a...)
}

func LogCriticalfd(ctx context.Context, format string, a ...interface{}) {
	file, line := logctx(1)

	log.Criticalf(ctx, "%s:%d %s", file, line, fmt.Sprintf(format, a...))
}
