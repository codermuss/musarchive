package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/mustafayilmazdev/musarchive/worker"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymetricKey:    util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	// Initialize the LocalizationManager singleton
	if err := localization.Initialize("../locales/assets/"); err != nil {

		log.Fatal().Msg("Can not load localization")
	}
	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)
	return server
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
