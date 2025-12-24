// Manages a user defined set of resources

package pool

import (
	"io"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

func (p *Pool) New() {

}

func (p *Pool) Acquire() {

}

func (p *Pool) Release() {

}

func (p *Pool) Close() {

}
