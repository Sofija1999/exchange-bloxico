package app

import (
	"github.com/Bloxico/exchange-gateway/internal/config"
	"github.com/Bloxico/exchange-gateway/internal/database"
	"github.com/Bloxico/exchange-gateway/internal/log"
)

type App struct {
	Config config.Config
	Logger log.Logger

	DB *database.DB
	// AMQP     *amqp.Client
}
