package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	config "github.com/uchupx/pintro-golang/config"
	data "github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/mysql"
	database "github.com/uchupx/pintro-golang/database"
	"github.com/uchupx/pintro-golang/helper/crypt"
	"github.com/uchupx/pintro-golang/transport/payload"
)

type Transport struct {
	mysqlConn    *sqlx.DB
	middleware   *Middleware
	cryptService crypt.CryptService

	gameHandler      *GameHandler
	genreHandler     *GenreHandler
	publisherHandler *PublisherHandler
	platformHandler  *PlatformHandler
	regionHandler    *RegionHandler
	userHandler      *UserHandler

	gameRepository          data.GameRepository
	genreRepository         data.GenreRepository
	publisherRepository     data.PublisherRepository
	platformRepository      data.PlatformRepository
	regionRepository        data.RegionRepository
	gamePublisherRepository data.GamePublisherRepository
	gamePlatformRepository  data.GamePlatformRepository
	userRepoitory           data.UserRepoitory

	gameResponseGenerator          *payload.GameResponseGenerator
	gamePublisherResponseGenerator *payload.GamePublisherResponseGenerator
	platfromResponseGenerator      *payload.PlatfromResponseGenerator
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

func (t Transport) GetPlatformRepo(conf *config.Config) data.PlatformRepository {
	if t.platformRepository == nil {
		repo := mysql.NewPlatformMysqlRepo(t.GetMySQLConn(conf))

		t.platformRepository = repo
	}

	return t.platformRepository
}

func (t Transport) GetPlatformHandler(conf *config.Config) *PlatformHandler {
	if t.platformHandler == nil {
		handler := PlatformHandler{
			PlatformRepository: t.GetPlatformRepo(conf),
		}

		t.platformHandler = &handler
	}

	return t.platformHandler
}

func (t Transport) GetGamePublisherRepo(conf *config.Config) data.GamePublisherRepository {
	if t.gamePublisherRepository == nil {
		repo := mysql.NewGamePublisherMysqlRepo(t.GetMySQLConn(conf))

		t.gamePublisherRepository = repo
	}

	return t.gamePublisherRepository
}

func (t Transport) GetUserRepository(conf *config.Config) data.UserRepoitory {
	if t.userRepoitory == nil {
		repo := mysql.NewUserMySQlRepo(t.GetMySQLConn(conf))

		t.userRepoitory = repo
	}

	return t.userRepoitory
}

func (t Transport) GetRegionRepo(conf *config.Config) data.RegionRepository {
	if t.regionRepository == nil {
		repo := mysql.NewRegionMysqlRepo(t.GetMySQLConn(conf))

		t.regionRepository = repo
	}

	return t.regionRepository
}

func (t Transport) GetGamePlatformRepo(conf *config.Config) data.GamePlatformRepository {
	if t.gamePlatformRepository == nil {
		repo := mysql.NewGamePlatformMysqlRepo(t.GetMySQLConn(conf))

		t.gamePlatformRepository = repo
	}

	return t.gamePlatformRepository
}

func (t Transport) GetRegionHandler(conf *config.Config) *RegionHandler {
	if t.regionHandler == nil {
		handler := RegionHandler{
			RegionRepository: t.GetRegionRepo(conf),
		}

		t.regionHandler = &handler
	}

	return t.regionHandler
}

func (t Transport) GetGameResponseGenerator(conf *config.Config) *payload.GameResponseGenerator {
	if t.gameResponseGenerator == nil {
		handler := payload.GameResponseGenerator{
			GenreRepository:                t.GetGenreRepo(conf),
			PublisherRepository:            t.GetPublisherRepo(conf),
			GamePublisherRepository:        t.GetGamePublisherRepo(conf),
			PlatformRepository:             t.GetPlatformRepo(conf),
			GamePlatformRepository:         t.GetGamePlatformRepo(conf),
			GamePublisherResponseGenerator: *t.GetGamePublisherResponseGenerator(conf),
		}

		t.gameResponseGenerator = &handler
	}

	return t.gameResponseGenerator
}

func (t Transport) GetGamePublisherResponseGenerator(conf *config.Config) *payload.GamePublisherResponseGenerator {
	if t.gamePublisherResponseGenerator == nil {
		handler := payload.GamePublisherResponseGenerator{
			GameRepository:            t.GetGameRepo(conf),
			PlatformRepository:        t.GetPlatformRepo(conf),
			GamePlatformRepository:    t.GetGamePlatformRepo(conf),
			PublisherRepository:       t.GetPublisherRepo(conf),
			PlatfromResponseGenerator: *t.GetPlatformResponseGenerator(conf),
		}

		t.gamePublisherResponseGenerator = &handler
	}

	return t.gamePublisherResponseGenerator
}

func (t Transport) GetPlatformResponseGenerator(conf *config.Config) *payload.PlatfromResponseGenerator {
	if t.platfromResponseGenerator == nil {
		handler := payload.PlatfromResponseGenerator{
			// RegionSales:         t.GetGameRepo(conf),
			PlatformRepository: t.GetPlatformRepo(conf),
			RegionRepository:   t.GetRegionRepo(conf),
			// PublisherRepository: t.GetPublisherRepo(conf),
		}
		t.platfromResponseGenerator = &handler
	}

	return t.platfromResponseGenerator
}

func (t Transport) GetCryptService(conf *config.Config) crypt.CryptService {
	if t.cryptService == nil {
		svc := crypt.NewCryptService(crypt.Params{
			Conf: conf,
		})

		t.cryptService = svc
	}

	return t.cryptService
}

func (t Transport) GetMiddleware(conf *config.Config) *Middleware {
	if t.middleware == nil {
		middlware := Middleware{
			CryptService: t.GetCryptService(conf),
		}

		t.middleware = &middlware
	}

	return t.middleware
}

func (t Transport) GetUserHandler(conf *config.Config) *UserHandler {
	if t.userHandler == nil {
		handler := UserHandler{
			UserRepository: t.GetUserRepository(conf),
			CryptService:   t.GetCryptService(conf),
		}

		t.userHandler = &handler
	}

	return t.userHandler
}

// /////////////
func shouldBind(c *gin.Context, body interface{}) error {
	err := c.ShouldBindWith(body, FQueryBinding{})
	if err != nil {
		return err
	}
	return nil
}
