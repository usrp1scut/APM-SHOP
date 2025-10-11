package dogapm

import (
	"context"

	"github.com/sirupsen/logrus"
)

type log struct {
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var Logger = &log{}

func (l *log) Info(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Info()
}

func (l *log) Warn(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Warn()

}

func (l *log) Debug(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Debug()
}

func (l *log) Error(ctx context.Context, action string, kv map[string]interface{}, err error) {
	kv["action"] = action
	logrus.WithFields(kv).WithError(err).Error()
}
