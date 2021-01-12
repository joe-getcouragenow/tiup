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

package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/pingcap/tiup/pkg/cluster/embed"
)

// VictoriaMetricsConfig represent the data to generate VictoriaMetrics config
type VictoriaMetricsConfig struct {
	ClusterName               string
	TLSEnabled                bool
	KafkaAddrs                []string
	NodeExporterAddrs         []string
	TiDBStatusAddrs           []string
	TiKVStatusAddrs           []string
	PDAddrs                   []string
	TiFlashStatusAddrs        []string
	TiFlashLearnerStatusAddrs []string
	PumpAddrs                 []string
	DrainerAddrs              []string
	CDCAddrs                  []string
	ZookeeperAddrs            []string
	BlackboxExporterAddrs     []string
	LightningAddrs            []string
	MonitoredServers          []string
	AlertmanagerAddrs         []string
	PushgatewayAddr           string
	BlackboxAddr              string
	KafkaExporterAddr         string
	GrafanaAddr               string

	DMMasterAddrs []string
	DMWorkerAddrs []string
}

// NewVictoriaMetricsConfig returns a VictoriaMetricsConfig
func NewVictoriaMetricsConfig(cluster string, enableTLS bool) *VictoriaMetricsConfig {
	return &VictoriaMetricsConfig{
		ClusterName: cluster,
		TLSEnabled:  enableTLS,
	}
}

// AddKafka add a kafka address
func (c *VictoriaMetricsConfig) AddKafka(ip string, port uint64) *VictoriaMetricsConfig {
	c.KafkaAddrs = append(c.KafkaAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddNodeExpoertor add a node expoter address
func (c *VictoriaMetricsConfig) AddNodeExpoertor(ip string, port uint64) *VictoriaMetricsConfig {
	c.NodeExporterAddrs = append(c.NodeExporterAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddTiDB add a TiDB address
func (c *VictoriaMetricsConfig) AddTiDB(ip string, port uint64) *VictoriaMetricsConfig {
	c.TiDBStatusAddrs = append(c.TiDBStatusAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddTiKV add a TiKV address
func (c *VictoriaMetricsConfig) AddTiKV(ip string, port uint64) *VictoriaMetricsConfig {
	c.TiKVStatusAddrs = append(c.TiKVStatusAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddPD add a PD address
func (c *VictoriaMetricsConfig) AddPD(ip string, port uint64) *VictoriaMetricsConfig {
	c.PDAddrs = append(c.PDAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddTiFlashLearner add a TiFlash learner address
func (c *VictoriaMetricsConfig) AddTiFlashLearner(ip string, port uint64) *VictoriaMetricsConfig {
	c.TiFlashLearnerStatusAddrs = append(c.TiFlashLearnerStatusAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddTiFlash add a TiFlash address
func (c *VictoriaMetricsConfig) AddTiFlash(ip string, port uint64) *VictoriaMetricsConfig {
	c.TiFlashStatusAddrs = append(c.TiFlashStatusAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddPump add a pump address
func (c *VictoriaMetricsConfig) AddPump(ip string, port uint64) *VictoriaMetricsConfig {
	c.PumpAddrs = append(c.PumpAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddDrainer add a drainer address
func (c *VictoriaMetricsConfig) AddDrainer(ip string, port uint64) *VictoriaMetricsConfig {
	c.DrainerAddrs = append(c.DrainerAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddCDC add a cdc address
func (c *VictoriaMetricsConfig) AddCDC(ip string, port uint64) *VictoriaMetricsConfig {
	c.CDCAddrs = append(c.CDCAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddZooKeeper add a zookeeper address
func (c *VictoriaMetricsConfig) AddZooKeeper(ip string, port uint64) *VictoriaMetricsConfig {
	c.ZookeeperAddrs = append(c.ZookeeperAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddBlackboxExporter add a BlackboxExporter address
func (c *VictoriaMetricsConfig) AddBlackboxExporter(ip string, port uint64) *VictoriaMetricsConfig {
	c.BlackboxExporterAddrs = append(c.BlackboxExporterAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddLightning add a lightning address
func (c *VictoriaMetricsConfig) AddLightning(ip string, port uint64) *VictoriaMetricsConfig {
	c.LightningAddrs = append(c.LightningAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddMonitoredServer add a MonitoredServer address
func (c *VictoriaMetricsConfig) AddMonitoredServer(ip string) *VictoriaMetricsConfig {
	c.MonitoredServers = append(c.MonitoredServers, ip)
	return c
}

// AddAlertmanager add an alertmanager address
func (c *VictoriaMetricsConfig) AddAlertmanager(ip string, port uint64) *VictoriaMetricsConfig {
	c.AlertmanagerAddrs = append(c.AlertmanagerAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddPushgateway add an pushgateway address
func (c *VictoriaMetricsConfig) AddPushgateway(ip string, port uint64) *VictoriaMetricsConfig {
	c.PushgatewayAddr = fmt.Sprintf("%s:%d", ip, port)
	return c
}

// AddBlackbox add an blackbox address
func (c *VictoriaMetricsConfig) AddBlackbox(ip string, port uint64) *VictoriaMetricsConfig {
	c.BlackboxAddr = fmt.Sprintf("%s:%d", ip, port)
	return c
}

// AddKafkaExporter add an kafka exporter address
func (c *VictoriaMetricsConfig) AddKafkaExporter(ip string, port uint64) *VictoriaMetricsConfig {
	c.KafkaExporterAddr = fmt.Sprintf("%s:%d", ip, port)
	return c
}

// AddGrafana add an kafka exporter address
func (c *VictoriaMetricsConfig) AddGrafana(ip string, port uint64) *VictoriaMetricsConfig {
	c.GrafanaAddr = fmt.Sprintf("%s:%d", ip, port)
	return c
}

// AddDMMaster add an dm-master address
func (c *VictoriaMetricsConfig) AddDMMaster(ip string, port uint64) *VictoriaMetricsConfig {
	c.DMMasterAddrs = append(c.DMMasterAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// AddDMWorker add an dm-worker address
func (c *VictoriaMetricsConfig) AddDMWorker(ip string, port uint64) *VictoriaMetricsConfig {
	c.DMWorkerAddrs = append(c.DMWorkerAddrs, fmt.Sprintf("%s:%d", ip, port))
	return c
}

// Config generate the config file data.
func (c *VictoriaMetricsConfig) Config() ([]byte, error) {
	fp := path.Join("/templates", "config", "VictoriaMetrics.yml.tpl")
	tpl, err := embed.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return c.ConfigWithTemplate(string(tpl))
}

// ConfigWithTemplate generate the VictoriaMetrics config content by tpl
func (c *VictoriaMetricsConfig) ConfigWithTemplate(tpl string) ([]byte, error) {
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
func (c *VictoriaMetricsConfig) ConfigToFile(file string) error {
	config, err := c.Config()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, config, 0755)
}
