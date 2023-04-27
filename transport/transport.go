package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	config "github.com/uchupx/pintro-golang/config"
	data "github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/mysql"
	database "github.com/uchupx/pintro-golang/database"
	"github.com/uchupx/pintro-golang/transport/payload"
)

type Transport struct {
	mysqlConn *sqlx.DB

	gameHandler      *GameHandler
	genreHandler     *GenreHandler
	publisherHandler *PublisherHandler

	gameRepository      data.GameRepository
	genreRepository     data.GenreRepository
	publisherRepository data.PublisherRepository

	gameResponseGenerator *payload.GameResponseGenerator
}

type CollectionsResponse struct {
	Perpage uint64                 `json:"perpage"`
	Page    uint64                 `json:"page"`
	Data    []payload.ResponseData `json:"data"`
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
			GameRepository:        t.GetGameRepo(conf),
			GameResponseGenerator: t.GetGameResponseGenerator(conf),
		}

		t.gameHandler = &handler
	}

	return t.gameHandler
}

func (t Transport) GetGenreRepo(conf *config.Config) data.GenreRepository {
	if t.genreRepository == nil {
		repo := mysql.NewGenreMysqlRepo(t.GetMySQLConn(conf))

		t.genreRepository = repo
	}

	return t.genreRepository
}

func (t Transport) GetGenreHandler(conf *config.Config) *GenreHandler {
	if t.genreHandler == nil {
		handler := GenreHandler{
			GenreRepository: t.GetGenreRepo(conf),
		}

		t.genreHandler = &handler
	}

	return t.genreHandler
}

func (t Transport) GetPublisherRepo(conf *config.Config) data.PublisherRepository {
	if t.publisherRepository == nil {
		repo := mysql.NewPublisherMysqlRepo(t.GetMySQLConn(conf))

		t.publisherRepository = repo
	}

	return t.publisherRepository
}

func (t Transport) GetPublisherHandler(conf *config.Config) *PublisherHandler {
	if t.publisherHandler == nil {
		handler := PublisherHandler{
			PublisherRepository: t.GetPublisherRepo(conf),
		}

		t.publisherHandler = &handler
	}

	return t.publisherHandler
}

func (t Transport) GetGameResponseGenerator(conf *config.Config) *payload.GameResponseGenerator {
	if t.gameResponseGenerator == nil {
		handler := payload.GameResponseGenerator{
			GenreRepository: t.GetGenreRepo(conf),
		}

		t.gameResponseGenerator = &handler
	}

	return t.gameResponseGenerator
}

// /////////////
func shouldBind(c *gin.Context, body interface{}) error {
	err := c.ShouldBindWith(body, FQueryBinding{})
	if err != nil {
		return err
	}
	return nil
}
