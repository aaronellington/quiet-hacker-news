package qhn

import "time"

// Config is the structure of the configuration options
type Config struct {
	Host                 string
	Port                 int           `env:"PORT"`
	PageSize             int           `env:"QHN_PAGE_SIZE"`
	RefreshIntervalHours time.Duration `env:"QHN_REFRESH_INTERVAL_HOURS"`
}
