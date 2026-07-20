package fasthttp

import (
	"errors"
	"sync"
	"time"
)

var ErrNoAvailableClients = errors.New("fasthttp: no available clients")

type BalancingClient interface {
	DoDeadline(req *Request, resp *Response, deadline time.Time) error
	PendingRequests() int
}

type LBClient struct {
	noCopy noCopy

	HealthCheck func(req *Request, resp *Response, err error) bool

	Clients []BalancingClient

	cs []*lbClient

	Timeout time.Duration

	mu sync.RWMutex

	once sync.Once
}

const DefaultLBClientTimeout = time.Second

func (cc *LBClient) DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *LBClient) DoTimeout(req *Request, resp *Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *LBClient) Do(req *Request, resp *Response) error { _ = "STUB: not implemented"; return nil }

func (cc *LBClient) init() {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	if len(cc.Clients) == 0 {

		panic("BUG: LBClient.Clients cannot be empty")
	}
	for _, c := range cc.Clients {
		cc.cs = append(cc.cs, &lbClient{
			c:           c,
			healthCheck: cc.HealthCheck,
		})
	}
}

func (cc *LBClient) AddClient(c BalancingClient) int { _ = "STUB: not implemented"; return 0 }

func (cc *LBClient) RemoveClients(rc func(BalancingClient) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func (cc *LBClient) get() *lbClient { _ = "STUB: not implemented"; return nil }

type lbClient struct {
	c           BalancingClient
	healthCheck func(req *Request, resp *Response, err error) bool
	penalty     uint32

	total uint64
}

func (c *lbClient) DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *lbClient) PendingRequests() int { _ = "STUB: not implemented"; return 0 }

func (c *lbClient) isHealthy(req *Request, resp *Response, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *lbClient) incPenalty() bool { _ = "STUB: not implemented"; return false }

func (c *lbClient) decPenalty() { _ = "STUB: not implemented"; return }

const (
	maxPenalty = 300

	penaltyDuration = 3 * time.Second
)
