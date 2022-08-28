package util

import "time"

const REDIS_REDIRECTION_KEY_PREFIX = "CACHE_REDIRECTION_"

var REDIS_REDIRECTION_TTL, _ = time.ParseDuration("5m")
