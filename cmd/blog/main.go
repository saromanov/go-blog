// Package blog provides starting of the service from command line
package blog

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/saromanov/go-blog/internal/platform/db"
	"github.com/saromanov/go-blog/internal/platform/db/postgresql"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .config.yml: %v", err)
	}

	return c, nil
}

// setupService provides setup of the all parts of the service
func setupService(cfg *Config) error {
	log.WithFields(log.Fields{
		"method": "setupService",
	}).Info("Initialization of storage")
	storage, err := postgresql.Create(&db.Config{

	})
	if err != nil {
		return fmt.Errorf("unable to setup storage: %v", err)
	}

	log.WithFields(log.Fields{
		"method": "setupService",
	  }).Info("Initialization of server")

	api := http.Server{
		Addr:           cfg.Host,
		Handler:        handlers.Hanlde(shutdown, log, storage, authenticator),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.WithFields(log.Fields{
			"method": "setupService",
		  }).Infof("API listening")
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <- serverErrors:
		log.WithFields(log.Fields{
			"method": "setupService",
		}).Errorf("unable to setup server: %v", err)
	}

	return nil
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	app := cli.NewApp()
	app.Name = "go-blog"
	app.Usage = "example of blog service"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to .yml config",
			Action: func(c *cli.Context) error {
				configPath := c.Args().First()
				config, err := parseConfig(configPath)
				if err != nil {
					panic(err)
				}
				if err := setupService(config); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}