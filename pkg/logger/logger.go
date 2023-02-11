package logger

import (
	"github.com/rs/zerolog/log"
)

// todo: im
var Logger = log.Logger.With().Caller().Logger()
