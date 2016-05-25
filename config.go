package main

import (
        "log"
        "io/ioutil"
        "gopkg.in/yaml.v2"
)

type DataSourceConfig struct {
    Host string
    Port int
    User string
    Password string
    Database string
}

type GraphiteConfig struct {
    Host string
    Port int
}

type PipeConfig struct {
    Source DataSourceConfig
    Query string
    Dest string
    Graphite GraphiteConfig
}

func LoadConfig(path string) []PipeConfig {
    data, err := ioutil.ReadFile(path)

    if err != nil {
        log.Fatalf("error: %v", err)
    }

    config := []PipeConfig{}

    err = yaml.Unmarshal([]byte(data), &config)

    if err != nil {
        log.Fatalf("error: %v", err)
    }

    return config
}