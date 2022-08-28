package util

import "time"

const REDIS_REDIRECTION_KEY_PREFIX = "CACHE_REDIRECTION_"
const FONT_COMIC_PATH = "assets/comic.ttf"
const CAPTHCA_TEXT_LENGTH = 6

var REDIS_REDIRECTION_TTL, _ = time.ParseDuration("5m")
