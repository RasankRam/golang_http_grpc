// https://habr.com/ru/companies/slurm/articles/798207/ да прост отсюда скопировал для контекста
package logger

import (
	"context"
	"log/slog"
	"slices"
)

type ctxKey string

var ContextKeys = make([]string, 0)

const (
	slogFields ctxKey = "slog_fields"
)

type ContextHandler struct {
	slog.Handler
}

func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, key := range ContextKeys {
		if attr, ok := ctx.Value(key).(slog.Attr); ok {
			r.AddAttrs(attr)
		}
	}

	return h.Handler.Handle(ctx, r)
}

func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if slices.Contains(ContextKeys, attr.Key) {
		ContextKeys = append(ContextKeys, attr.Key)
	}

	return context.WithValue(parent, attr.Key, attr)
}
