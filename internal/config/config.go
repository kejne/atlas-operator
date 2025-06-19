package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	devDBConfig     *DevDBConfig
	devDBConfigOnce sync.Once
	confLog         = ctrl.Log.WithName("devdb-config")
)

type DevDBConfig struct {
	Prewarm bool `yaml:"prewarm"`
	Images  struct {
		Postgres   string `yaml:"postgres"`
		SQLServer  string `yaml:"sqlServer"`
		MySQL      string `yaml:"mySql"`
		MariaDB    string `yaml:"mariaDb"`
		ClickHouse string `yaml:"clickHouse"`
	} `yaml:"images"`
}

func MustGetDevDBConfig() *DevDBConfig {

	devDBConfigOnce.Do(func() {
		file, err := os.Open("/etc/config/devdb.yaml")
		if err != nil {
			confLog.Error(err, "unable to read DevDb config")
			os.Exit(1)
		}
		defer file.Close()

		config := &DevDBConfig{}
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(config); err != nil {
			confLog.Error(err, "unable to read DevDb config")
			os.Exit(1)
		}

		devDBConfig = config
	})

	return devDBConfig
}
