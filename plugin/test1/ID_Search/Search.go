package ID_Search

import "time"

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0"
	Qps       = 50
)

var rateLimiter = time.Tick(
	time.Second / Qps)

type Searching struct {
	Request *Request
}

func NewSearch() *Searching {
	return &Searching{}
}

func (g *Searching) Visit() error {
	return g.Request.Do()
}
