package config

import (
	"encoding/json"
	"os"
)

const (
	defaultHostname        = "localhost"
	defaultPort            = 9091
	defaultMinimumSeeding  = 1128
	defaultMinimumSeeders  = 2
	defaultMinimumLeechers = 1
	defaultLastActivity    = 336
	defaultRatioTolerance  = 0.05
	defaultRatioCoefA      = 5
	defaultRatioCoefB      = 0.006
	defaultRatioCoefC      = 0.01
)

type Config struct {
	Hostname        string `json:"hostname"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Port            uint16 `json:"port"`
	MinimumSeeding  int64  `json:"minimum_seeding"`
	MinimumSeeders  int64  `json:"minimum_seeders"`
	MinimumLeechers int64  `json:"minimum_leechers"`
	LastActivity    int64  `json:"last_activity"`
	RatioTolerance  float64 `json:"ratio_tolerance"`
	RatioCoefA      float64 `json:"ratio_coef_a"`
	RatioCoefB      float64 `json:"ratio_coef_b"`
	RatioCoefC      float64 `json:"ratio_coef_c"`
}

func Load(filename string) (*Config, error) {
	byteConfig, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = json.Unmarshal(byteConfig, &conf)
	if err != nil {
		return nil, err
	}

	if conf.Hostname == "" {
		conf.Hostname = defaultHostname
	}

	if conf.Port == 0 {
		conf.Port = defaultPort
	}

	if conf.MinimumSeeding == 0 {
		conf.MinimumSeeding = defaultMinimumSeeding
	}

	if conf.MinimumSeeders == 0 {
		conf.MinimumSeeders = defaultMinimumSeeders
	}

	if conf.MinimumLeechers == 0 {
		conf.MinimumLeechers = defaultMinimumLeechers
	}

	if conf.LastActivity == 0 {
		conf.LastActivity = defaultLastActivity
	}

	if conf.RatioTolerance == 0 {
		conf.RatioTolerance = defaultRatioTolerance
	}

	if conf.RatioCoefA == 0 {
		conf.RatioCoefA = defaultRatioCoefA
	}

	if conf.RatioCoefB == 0 {
		conf.RatioCoefB = defaultRatioCoefB
	}

	if conf.RatioCoefC == 0 {
		conf.RatioCoefC = defaultRatioCoefC
	}

	return &conf, nil
}
