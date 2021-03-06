package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const (
	vmAddr        = "vm-addr"
	vmUser        = "vm-user"
	vmPassword    = "vm-password"
	vmAccountID   = "vm-account-id"
	vmConcurrency = "vm-concurrency"
	vmCompress    = "vm-compress"
	vmBatchSize   = "vm-batch-size"
)

var (
	vmFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  vmAddr,
			Value: "http://localhost:8428",
			Usage: "VictoriaMetrics address to perform import requests. " +
				"Should be the same as --httpListenAddr value for single-node version or VMSelect component.",
		},
		&cli.StringFlag{
			Name:    vmUser,
			Usage:   "VictoriaMetrics username for basic auth",
			EnvVars: []string{"VM_USERNAME"},
		},
		&cli.StringFlag{
			Name:    vmPassword,
			Usage:   "VictoriaMetrics password for basic auth",
			EnvVars: []string{"VM_PASSWORD"},
		},
		&cli.IntFlag{
			Name:  vmAccountID,
			Value: -1,
			Usage: "Account(tenant) ID - is required for cluster VM.",
		},
		&cli.UintFlag{
			Name:  vmConcurrency,
			Usage: "Number of workers concurrently performing import requests to VM",
			Value: 2,
		},
		&cli.BoolFlag{
			Name:  vmCompress,
			Value: true,
			Usage: "Whether to apply gzip compression to import requests",
		},
		&cli.IntFlag{
			Name:  vmBatchSize,
			Value: 200e3,
			Usage: "How many datapoints importer collects before sending the import request to VM",
		},
	}
)

const (
	influxAddr                      = "influx-addr"
	influxUser                      = "influx-user"
	influxPassword                  = "influx-password"
	influxDB                        = "influx-database"
	influxRetention                 = "influx-retention-policy"
	influxChunkSize                 = "influx-chunk-size"
	influxConcurrency               = "influx-concurrency"
	influxFilterSeries              = "influx-filter-series"
	influxFilterTimeStart           = "influx-filter-time-start"
	influxFilterTimeEnd             = "influx-filter-time-end"
	influxMeasurementFieldSeparator = "influx-measurement-field-separator"
)

var (
	influxFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  influxAddr,
			Value: "http://localhost:8086",
			Usage: "Influx server addr",
		},
		&cli.StringFlag{
			Name:    influxUser,
			Usage:   "Influx user",
			EnvVars: []string{"INFLUX_USERNAME"},
		},
		&cli.StringFlag{
			Name:    influxPassword,
			Usage:   "Influx user password",
			EnvVars: []string{"INFLUX_PASSWORD"},
		},
		&cli.StringFlag{
			Name:     influxDB,
			Usage:    "Influx database",
			Required: true,
		},
		&cli.StringFlag{
			Name:  influxRetention,
			Usage: "Influx retention policy",
			Value: "autogen",
		},
		&cli.IntFlag{
			Name:  influxChunkSize,
			Usage: "The chunkSize defines max amount of series to be returned in one chunk",
			Value: 10e3,
		},
		&cli.IntFlag{
			Name:  influxConcurrency,
			Usage: "Number of concurrently running fetch queries to InfluxDB",
			Value: 1,
		},
		&cli.StringFlag{
			Name: influxFilterSeries,
			Usage: "Influx filter expression to select series. E.g. \"from cpu where arch='x86' AND hostname='host_2753'\".\n" +
				"See for details https://docs.influxdata.com/influxdb/v1.7/query_language/schema_exploration#show-series",
		},
		&cli.StringFlag{
			Name:  influxFilterTimeStart,
			Usage: "The time filter to select timeseries with timestamp equal or higher than provided value. E.g. '2020-01-01T20:07:00Z'",
		},
		&cli.StringFlag{
			Name:  influxFilterTimeEnd,
			Usage: "The time filter to select timeseries with timestamp equal or lower than provided value. E.g. '2020-01-01T20:07:00Z'",
		},
		&cli.StringFlag{
			Name:  influxMeasurementFieldSeparator,
			Usage: "The {separator} symbol used to concatenate {measurement} and {field} names into series name {measurement}{separator}{field}.",
			Value: "_",
		},
	}
)

const (
	promSnapshot         = "prom-snapshot"
	promConcurrency      = "prom-concurrency"
	promFilterTimeStart  = "prom-filter-time-start"
	promFilterTimeEnd    = "prom-filter-time-end"
	promFilterLabel      = "prom-filter-label"
	promFilterLabelValue = "prom-filter-label-value"
)

var (
	promFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     promSnapshot,
			Usage:    "Path to Prometheus snapshot. Pls see for details https://www.robustperception.io/taking-snapshots-of-prometheus-data",
			Required: true,
		},
		&cli.IntFlag{
			Name:  promConcurrency,
			Usage: "Number of concurrently running snapshot readers",
			Value: 1,
		},
		&cli.StringFlag{
			Name:  promFilterTimeStart,
			Usage: "The time filter to select timeseries with timestamp equal or higher than provided value. E.g. '2020-01-01T20:07:00Z'",
		},
		&cli.StringFlag{
			Name:  promFilterTimeEnd,
			Usage: "The time filter to select timeseries with timestamp equal or lower than provided value. E.g. '2020-01-01T20:07:00Z'",
		},
		&cli.StringFlag{
			Name:  promFilterLabel,
			Usage: "Prometheus label name to filter timeseries by. E.g. '__name__' will filter timeseries by name.",
		},
		&cli.StringFlag{
			Name:  promFilterLabelValue,
			Usage: fmt.Sprintf("Prometheus regular expression to filter label from %q flag.", promFilterLabel),
			Value: ".*",
		},
	}
)
