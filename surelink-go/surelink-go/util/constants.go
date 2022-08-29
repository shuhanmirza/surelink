package util

import "time"

const REDIS_REDIRECTION_KEY_PREFIX = "CACHE_REDIRECTION_"
const REDIS_CAPTCHA_KEY_PREFIX = "CACHE_CAPTCHA_"
const FONT_COMIC_PATH = "assets/comic.ttf"
const CAPTHCA_TEXT_LENGTH = 6

var REDIS_REDIRECTION_TTL, _ = time.ParseDuration("5m")
var REDIS_CAPTCHA_TTL, _ = time.ParseDuration("5m")
