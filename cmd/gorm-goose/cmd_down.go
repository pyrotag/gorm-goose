package main

import (
	"log"

	goose "github.com/pyrotag/gorm-goose/lib/gorm-goose"
)

var downCmd = &Command{
	Name:    "down",
	Usage:   "",
	Summary: "Roll back the version by 1",
	Help:    `down extended help here...`,
	Run:     downRun,
}

func downRun(cmd *Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	current, err := goose.GetDBVersion(conf)
	if err != nil {
		log.Fatal(err)
	}

	previous, err := goose.GetPreviousDBVersion(conf.MigrationsDir, current)
	if err != nil {
		log.Fatal(err)
	}

	if err = goose.RunMigrations(conf, conf.MigrationsDir, previous); err != nil {
		log.Fatal(err)
	}
}
