// Package blog provides starting of the service from command line
package blog

import (
	"fmt"
	"github.com/saromanov/go-blog/internal/platform/db"
	"github.com/saromanov/go-blog/internal/platform/db/postgresql"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*structs.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *structs.Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .config.yml: %v", err)
	}

	return c, nil
}

// setupService provides setup of the all parts of the service
func setupService() error {
	storage, err := postgresql.Create(&db.Config{

	})
	if err != nil {
		return fmt.Errorf("unable to setup storage: %v", err)
	}

	api := http.Server{
		Addr:           cfg.Web.APIHost,
		Handler:        handlers.Hanlde(shutdown, log, storage, authenticator),
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main : API Listening %s", cfg.Web.APIHost)
		serverErrors <- api.ListenAndServe()
	}()

	return nil
}

func main() {
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
				if err := run(config); err != nil {
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