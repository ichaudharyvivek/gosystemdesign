package stats

import (
	"fmt"
	"time"
)

type URLStats struct {
	Code       string
	Clicks     int
	AccessLogs []time.Time
}

func (s *URLStats) String() string {
	return fmt.Sprintf("{'code': %s, 'clicks': %d, 'accessLogs': %v}", s.Code, s.Clicks, s.AccessLogs)
}
