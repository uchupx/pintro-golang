package transport

import (
	"github.com/jmoiron/sqlx"
	config "github.com/uchupx/pintro-golang/config"
	data "github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/mysql"
	database "github.com/uchupx/pintro-golang/database"
)

type Transport struct {
	mysqlConn *sqlx.DB

	gameHandler *GameHandler

	gameRepository data.GameRepository
}

func (t Transport) GetMySQLConn(conf *config.Config) *sqlx.DB {
	if t.mysqlConn == nil {
		mysqlConfig := database.Config{
			HostName: conf.Host,
			Username: conf.Username,
			Database: conf.Database,
			Password: conf.Password,
		}

		conn, err := database.NewConnection(mysqlConfig)
		if err != nil {
			panic(err)
		}

		t.mysqlConn = conn
	}

	return t.mysqlConn
}

func (t Transport) GetGameRepo(conf *config.Config) data.GameRepository {
	if t.gameRepository == nil {
		repo := mysql.NewGameMysqlRepo(t.GetMySQLConn(conf))

		t.gameRepository = repo
	}

	return t.gameRepository
}

func (t Transport) GetGameHandler(conf *config.Config) *GameHandler {
	if t.gameHandler == nil {
		handler := GameHandler{
			GameRepository: t.GetGameRepo(conf),
		}

		t.gameHandler = &handler
	}

	return t.gameHandler
}
