package stats

import (
	"fmt"
	"time"
)

type URLStats struct {
	Count int
	Code  string
	Logs  []time.Time
}

func (s *URLStats) String() string {
	return fmt.Sprintf("{'code': %s, 'clicks': %d, 'accessLogs': %v}", s.Code, s.Count, s.Logs)
}
