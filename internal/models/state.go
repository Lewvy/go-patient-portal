package models

import (
	"github.com/Lewvy/markable/internal/config"
	"github.com/Lewvy/markable/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}
