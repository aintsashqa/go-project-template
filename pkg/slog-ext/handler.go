package slogext

import (
	"context"
	"log/slog"
)

type ContextHandler struct {
	next slog.Handler
}

func NewContextHandler(next slog.Handler) *ContextHandler {
	return &ContextHandler{
		next: next,
	}
}

func (hdlr *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return hdlr.next.Enabled(ctx, level)
}

func (hdlr *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	ctxMu.RLock()
	defer ctxMu.RUnlock()

	for _, ctxAttr := range ctxAttrs {
		if value := ctx.Value(ctxAttr.key); value != nil {
			r.AddAttrs(slog.Any(ctxAttr.name, value))
		}
	}

	return hdlr.next.Handle(ctx, r)
}

func (hdlr *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewContextHandler(hdlr.next.WithAttrs(attrs))
}

func (hdlr *ContextHandler) WithGroup(name string) slog.Handler {
	return NewContextHandler(hdlr.next.WithGroup(name))
}
