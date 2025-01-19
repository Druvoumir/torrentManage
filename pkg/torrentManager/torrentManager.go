package torrentManager

import (
	"context"
	"fmt"
	"math"
	"time"

	"patu.re/torrentManager/pkg/config"

	"github.com/hekmon/transmissionrpc/v2"
)

func ListRemovableTorrents(conf *config.Config) {
	tbt, err := transmissionrpc.New(
		conf.Hostname,
		conf.Username,
		conf.Password,
		&transmissionrpc.AdvancedConfig{
			Port:  conf.Port,
		},
	)
	if err != nil {
		panic(err)
	}
	torrents, err := tbt.TorrentGetAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, t := range torrents {
		if isTorrentRemovable(t, conf) {
			fmt.Println(*t.Name, *t.SecondsSeeding)
		}
	}
}

func isTorrentRemovable(t transmissionrpc.Torrent, conf *config.Config) bool {
	// Can't delete if not seeded enough time
	if t.SecondsSeeding == nil || *t.SecondsSeeding <= time.Duration(conf.MinimumSeeding) * time.Hour {
		return false
	}

	// Forbid to delete if there are less (or equal) than 2 seeders
	// Forbid to delete if there are leechers
	var maxSeeders int64 = 0
	for _, stats := range t.TrackerStats {
		if stats.SeederCount <= conf.MinimumSeeders || stats.LeecherCount >= conf.MinimumLeechers {
			return false
		}

		if stats.SeederCount > maxSeeders {
			maxSeeders = stats.SeederCount
		}
	}

	// Forbid to delete if last activity was less than 2 weeks ago
	if t.ActivityDate != nil && time.Since(*t.ActivityDate) <= time.Duration(conf.LastActivity) * time.Hour {
		return false
	}

	// Calculate Removability score based on parameters
	// x: time, s: num seeders -> f(x, s) = A*e^(-((B*s+C)*x)^2)
	// param: tolerence, max ratio (A), B C
	bonusSeedDays := float64(*t.SecondsSeeding - (time.Duration(conf.MinimumSeeding) * time.Hour)) / float64(24 * time.Hour)
	var ratio float64 = 0
	if t.UploadRatio != nil {
		ratio = *t.UploadRatio - conf.RatioTolerance
	}
	threshold := conf.RatioCoefA * math.Exp(-1 * math.Pow((conf.RatioCoefB*float64(maxSeeders) + conf.RatioCoefC)*bonusSeedDays, 2))
	if ratio <= threshold {
		return false
	}

	return true
}
