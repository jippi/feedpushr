/*
	A simple feed aggregator daemon with sugar on top.

	Copyright (C) 2018 Nicolas Carlier

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/config"
	configflag "github.com/ncarlier/feedpushr/v3/pkg/config/flag"
	"github.com/ncarlier/feedpushr/v3/pkg/logging"
	"github.com/ncarlier/feedpushr/v3/pkg/metric"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/opml"
	"github.com/ncarlier/feedpushr/v3/pkg/server"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/ncarlier/feedpushr/v3/pkg/version"
	"github.com/rs/zerolog/log"
)

func main() {
	// Get global configuration
	conf := config.Config{}
	configflag.Bind(&conf, "FP")

	// Parse command line (and environment variables)
	flag.Parse()

	// Shutdown channels
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Show version if asked
	if *version.ShowVersion {
		version.Print()
		os.Exit(0)
	}

	// Log configuration
	if err := logging.Configure(conf.LogOutput, conf.LogLevel, conf.LogPretty, conf.SentryDSN); err != nil {
		log.Fatal().Err(err).Msg("unable to configure logger")
	}

	// Metric configuration
	metric.Configure(conf)

	// Init the data store
	db, err := store.NewDB(conf.DB, model.Quota{
		MaxNbFeeds:   conf.MaxNbFeeds,
		MaxNbOutputs: conf.MaxNbOutputs,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init data store")
	}

	// Init search index
	if err := db.BuildInitialIndex(); err != nil {
		log.Fatal().Err(err).Msg("unable to init search index")
	}

	// Import OPML file if asked
	if conf.ImportFilename != "" {
		log.Debug().Str("filename", conf.ImportFilename).Msg("importing OPML file...")
		importer := opml.NewOPMLImporter(db)
		job, err := importer.ImportOPMLFile(conf.ImportFilename)
		if err != nil {
			db.Close()
			log.Fatal().Err(err).Str("filename", conf.ImportFilename).Msg("unable to import OPML file")
		}
		if !job.Wait(time.Minute) {
			log.Warn().Str("filename", conf.ImportFilename).Msg("OPML is still importing in background...")
		} else {
			log.Info().Str("filename", conf.ImportFilename).Msg("OPML file imported")
		}
	}

	// Init global service
	srv, err := server.NewServer(db, conf)
	if err != nil {
		db.Close()
		log.Fatal().Err(err).Msg("unable to init server")
	}

	// Graceful shutdown handler
	go func() {
		<-quit
		log.Debug().Msg("shutting down server...")
		// Shutdown the server...
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("could not gracefully shutdown the server")
		}
		// Shutdown the database...
		db.Close()
		close(done)
	}()

	// Start service
	log.Info().Str("listen", conf.ListenAddr).Msg("starting HTTP server...")
	if err := srv.ListenAndServe(conf.ListenAddr); err != nil {
		log.Fatal().Err(err).Msg("unable to start server")
	}

	<-done
	log.Debug().Msg("server stopped")
}
