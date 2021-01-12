// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package scripts

import (
	"bytes"
	"io/ioutil"
	"path"
	"regexp"
	"text/template"

	"github.com/pingcap/tiup/pkg/cluster/embed"
)

// VictoriaMetricsScript represent the data to generate VictoriaMetrics config
type VictoriaMetricsScript struct {
	IP        string
	Port      int
	DeployDir string
	DataDir   string
	LogDir    string
	NumaNode  string
	Retention string
	tplFile   string
}

// NewVictoriaMetricsScript returns a VictoriaMetricsScript with given arguments
func NewVictoriaMetricsScript(ip, deployDir, dataDir, logDir string) *VictoriaMetricsScript {
	return &VictoriaMetricsScript{
		IP:        ip,
		Port:      9090,
		DeployDir: deployDir,
		DataDir:   dataDir,
		LogDir:    logDir,
	}
}

// WithPort set Port field of VictoriaMetricsScript
func (c *VictoriaMetricsScript) WithPort(port int) *VictoriaMetricsScript {
	c.Port = port
	return c
}

// WithNumaNode set NumaNode field of VictoriaMetricsScript
func (c *VictoriaMetricsScript) WithNumaNode(numa string) *VictoriaMetricsScript {
	c.NumaNode = numa
	return c
}

// WithRetention set Retention field of VictoriaMetricsScript
func (c *VictoriaMetricsScript) WithRetention(retention string) *VictoriaMetricsScript {
	valid, _ := regexp.MatchString("^[1-9]\\d*d$", retention)
	if retention == "" || !valid {
		c.Retention = "30d"
	} else {
		c.Retention = retention
	}
	return c
}

// WithTPLFile set the template file.
func (c *VictoriaMetricsScript) WithTPLFile(fname string) *VictoriaMetricsScript {
	c.tplFile = fname
	return c
}

// Config generate the config file data.
func (c *VictoriaMetricsScript) Config() ([]byte, error) {
	fp := c.tplFile
	if fp == "" {
		fp = path.Join("/templates", "scripts", "run_prometheus.sh.tpl")
	}

	tpl, err := embed.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return c.ConfigWithTemplate(string(tpl))
}

// ConfigWithTemplate generate the VictoriaMetrics config content by tpl
func (c *VictoriaMetricsScript) ConfigWithTemplate(tpl string) ([]byte, error) {
	tmpl, err := template.New("VictoriaMetrics").Parse(tpl)
	if err != nil {
		return nil, err
	}

	content := bytes.NewBufferString("")
	if err := tmpl.Execute(content, c); err != nil {
		return nil, err
	}

	return content.Bytes(), nil
}

// ConfigToFile write config content to specific path
func (c *VictoriaMetricsScript) ConfigToFile(file string) error {
	config, err := c.Config()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, config, 0755)
}

