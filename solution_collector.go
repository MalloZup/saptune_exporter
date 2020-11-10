package main

import (
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const subsystem = "solution"

// SolutionCollector is the saptune solution collector
type SolutionCollector struct {
	DefaultCollector
}

// NewSolutionCollector creates a new solution saptune collector
func NewSolutionCollector() (*SolutionCollector, error) {
	c := &SolutionCollector{
		NewDefaultCollector(subsystem),
	}
	c.SetDescriptor("name", "This metrics show with 1 all the enabled notes on the system", []string{"solutionID"})
	return c, nil
}

// Collect various metrics for saptune solution
func (c *SolutionCollector) Collect(ch chan<- prometheus.Metric) {
	log.Debugln("Collecting saptune solution metrics...")
	solutionName, err := exec.Command("saptune", "solution", "enabled").CombinedOutput()
	if err != nil {
		log.Warnf("%v - Failed to run saptune solution enabled command n %s ", err, string(solutionName))
		return
	}

	ch <- c.MakeGaugeMetric("enabled", float64(1), string(solutionName))

}
