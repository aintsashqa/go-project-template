package slogext

import "sync"

var (
	ctxMu    sync.RWMutex
	ctxAttrs []contextAttr
)

type contextAttr struct {
	key  any
	name string
}

func AddContext(key any, name string) {
	ctxMu.Lock()
	defer ctxMu.Unlock()

	ctxAttrs = append(ctxAttrs, contextAttr{key: key, name: name})
}
