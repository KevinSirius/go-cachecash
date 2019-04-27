package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	_ "net/http/pprof"
	"os"

	"github.com/cachecashproject/go-cachecash/cache"
	"github.com/cachecashproject/go-cachecash/cache/migrations"
	"github.com/cachecashproject/go-cachecash/common"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

var (
	logLevelStr = flag.String("logLevel", "info", "Verbosity of log output")
	logCaller   = flag.Bool("logCaller", false, "Enable method name logging")
	logFile     = flag.String("logFile", "", "Path where file should be logged")
	configPath  = flag.String("config", "cache.config.json", "Path to configuration file")
)

func loadConfigFile(path string) (*cache.ConfigFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cf cache.ConfigFile
	if err := json.Unmarshal(data, &cf); err != nil {
		return nil, err
	}

	return &cf, nil
}

func main() {
	if err := mainC(); err != nil {
		if _, err := os.Stderr.WriteString(err.Error() + "\n"); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func mainC() error {
	flag.Parse()
	log.SetFlags(0)

	l := logrus.New()
	if err := common.ConfigureLogger(l, &common.LoggerConfig{
		LogLevelStr: *logLevelStr,
		LogCaller:   *logCaller,
		LogFile:     *logFile,
	}); err != nil {
		return errors.Wrap(err, "failed to configure logger")
	}

	cf, err := loadConfigFile(*configPath)
	if err != nil {
		return errors.Wrap(err, "failed to load configuration file")
	}

	db, err := sql.Open("sqlite3", cf.Database)
	if err != nil {
		return errors.Wrap(err, "failed to open database")
	}

	l.Info("applying migrations")
	n, err := migrate.Exec(db, "sqlite3", migrations.Migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	l.Infof("applied %d migrations", n)

	c, err := cache.NewCache(l, db, cf)
	if err != nil {
		return nil
	}
	defer c.Close()

	c.Escrows = cf.Escrows

	app, err := cache.NewApplication(l, c, cf.Config)
	if err != nil {
		return errors.Wrap(err, "failed to create cache application")
	}

	if err := common.RunStarterShutdowner(app); err != nil {
		return err
	}
	return nil
}
