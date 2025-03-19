package gorm

import "time"

type PgConfig struct {
	LogQuery      bool
	SingularTable bool
}

type PgOptions struct {
	Host        string
	User        string
	Password    string
	DBName      string
	SSLMode     string
	Port        int
	PoolSize    int
	MaxIdleTime time.Duration
}

func (opt *PgOptions) setDefaultValues() {
	const maxIdleSec = 60

	if opt.PoolSize == 0 {
		opt.PoolSize = 10
	}

	if opt.MaxIdleTime == 0 {
		opt.MaxIdleTime = maxIdleSec * time.Second
	}

	if opt.SSLMode == "" {
		opt.SSLMode = "disable"
	}
}
