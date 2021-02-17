package application

import (
	"github.com/lukluk/link-proxy/config"
	"github.com/rs/zerolog/log"
	"net/http"
)

func HTTPServer(cfg config.Config) {
	log.Info().Msg("Starting server at port 8080")
	if err := http.ListenAndServe(":" + cfg.Port, nil); err != nil {
		log.Fatal().Msgf("failed start http server, errors: %v", err)
	}
}
