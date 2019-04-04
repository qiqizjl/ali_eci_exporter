package exporter

import (
	"ali_eci_exporter/eci"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

var namespace = "alicloud_eci"

type Exporter struct {
	eciClient       *eci.Client
	metricsMtx      sync.RWMutex
	cpu             *prometheus.Desc
	useMemory       *prometheus.Desc
	availableMemory *prometheus.Desc
	networkRx       *prometheus.Desc
	networkTx       *prometheus.Desc
}

func MewExporter(accesskey, accesssecert, region string) (*Exporter, error) {
	exporter := &Exporter{}
	var err error
	exporter.eciClient, err = eci.NewClientWithAccessKey(region, accesskey, accesssecert)
	if err != nil {
		return nil, err
	}
	exporter.init()
	return exporter, nil
}

func (e *Exporter) init() {
	e.cpu = prometheus.NewDesc(prometheus.BuildFQName(namespace, "cpu", "cores"), "Cpu Use", []string{"eci_id"}, nil)
	e.useMemory = prometheus.NewDesc(prometheus.BuildFQName(namespace, "memory", "use"), "Memory Use", []string{"eci_id"}, nil)
	e.availableMemory = prometheus.NewDesc(prometheus.BuildFQName(namespace, "memory", "available"), "Memory All", []string{"eci_id"}, nil)
	e.networkRx = prometheus.NewDesc(prometheus.BuildFQName(namespace, "network", "rx"), "NetWorkRX", []string{"eci_id"}, nil)
	e.networkTx = prometheus.NewDesc(prometheus.BuildFQName(namespace, "network", "tx"), "NetWorkTX", []string{"eci_id"}, nil)
}

// Describe outputs Redis metric descriptions.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.cpu
	ch <- e.useMemory
	ch <- e.availableMemory
	ch <- e.networkTx
	ch <- e.networkRx
}

// 采集数据
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	list, err := e.getEciMetricData()
	if err != nil {
		return
	}
	for _, v := range list {
		if len(v.Records) > 0 {
			record := v.Records[len(v.Records)-1]
			//if len(record.Network.Interfaces) <= 0 {
			//	continue
			//}
			if record.Cpu.UsageNanoCores > 0 {
				ch <- prometheus.MustNewConstMetric(
					e.cpu,
					prometheus.GaugeValue,
					float64(record.Cpu.UsageNanoCores)/1000000000,
					v.ContainerGroupId,
				)
			}
			if record.Memory.UsageBytes > 0 {
				ch <- prometheus.MustNewConstMetric(
					e.useMemory,
					prometheus.GaugeValue,
					float64(record.Memory.UsageBytes),
					v.ContainerGroupId,
				)
			}
			if record.Memory.AvailableBytes > 0 {
				ch <- prometheus.MustNewConstMetric(
					e.availableMemory,
					prometheus.GaugeValue,
					float64(record.Memory.AvailableBytes),
					v.ContainerGroupId,
				)
			}
			if len(record.Network.Interfaces) > 0 {
				ch <- prometheus.MustNewConstMetric(
					e.networkRx,
					prometheus.GaugeValue,
					float64(record.Network.Interfaces[0].RxBytes),
					v.ContainerGroupId,
				)
				ch <- prometheus.MustNewConstMetric(
					e.networkTx,
					prometheus.GaugeValue,
					float64(record.Network.Interfaces[0].TxBytes),
					v.ContainerGroupId,
				)
			}
		}
	}
}

func (e *Exporter) getEciMetricData() ([]eci.DescribeContainerGroupMetricResponse, error) {
	list := make([]eci.DescribeContainerGroupMetricResponse, 0)
	nextToken := ""
	for true {
		request := eci.CreateDescribeContainerGroupsRequest()
		request.NextToken = nextToken
		result, err := e.eciClient.DescribeContainerGroups(request)
		if err != nil {
			return nil, err
		}
		//tmpEci := make([]string, 0)
		for _, v := range result.ContainerGroups {
			monitorRequest := eci.CreateDescribeContainerGroupMetricRequest()
			monitorRequest.ContainerGroupId = v.ContainerGroupId
			monitorResult, err := e.eciClient.DescribeContainerGroupMetric(monitorRequest)
			if err != nil {
				return nil, err
			}
			list = append(list, *monitorResult)
		}
		////去拿数据
		nextToken = result.NextToken
		if result.NextToken == "" {
			break
		}
	}
	return list, nil
}
