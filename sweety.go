package sweetygo

import (
	"fmt"
	"net/http"
	"sync"
)

// HandlerFunc context handler func
type HandlerFunc func(*Context)

// SweetyGo is Suuuuuuuuper Sweetie!
type SweetyGo struct {
	Tree                    *Trie
	Pool                    sync.Pool
	NotFoundHandler         HandlerFunc
	MethodNotAllowedHandler HandlerFunc
	Middlewares             []HandlerFunc
}

// New SweetyGo App
func New() *SweetyGo {
	tree := &Trie{
		component: "/",
		methods:   make(map[string]HandlerFunc),
	}
	sg := &SweetyGo{Tree: tree,
		NotFoundHandler:         NotFoundHandler,
		MethodNotAllowedHandler: MethodNotAllowedHandler,
		Middlewares:             make([]HandlerFunc, 0),
	}
	sg.Pool = sync.Pool{
		New: func() interface{} {
			return NewContext(nil, nil, sg)
		},
	}
	return sg
}

// USE middlewares for SweetyGo
func (sg *SweetyGo) USE(middlewares ...HandlerFunc) {
	for i := range middlewares {
		if middlewares[i] != nil {
			sg.Middlewares = append(sg.Middlewares, middlewares[i])
		}
	}
}

// RunServer at the given addr
func (sg *SweetyGo) RunServer(addr string) {
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	http.ListenAndServe(addr, sg)
}
