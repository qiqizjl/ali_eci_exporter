package eci

type MetricRecord struct {
	Timestamp  string                   `json:"Timestamp",xml:"Timestamp"`
	Cpu        MetricRecordCpuStats     `json:"CPU",xml:"CPU"`
	Memory     MetricRecordMemoryStats  `json:"Memory",xml:"Memory"`
	Network    MetricRecordNetworkStats `json:"Network",xml:"Network"`
	Containers []MetricRecordContainer  `json:"Containers",xml:"Containers"`
}

type MetricRecordContainer struct {
	Name   string                  `json:"Name",xml:"Name"`
	CPU    MetricRecordCpuStats    `json:"CPU",xml:"CPU"`
	Memory MetricRecordMemoryStats `json:"Timestamp",xml:"Memory"`
}

type MetricRecordCpuStats struct {
	UsageNanoCores       uint64 `json:"UsageNanoCores",xml:"UsageNanoCores"`
	UsageCoreNanoSeconds uint64 `json:"UsageCoreNanoSeconds",xml:"UsageCoreNanoSeconds"`
}
type MetricRecordMemoryStats struct {
	AvailableBytes uint64 `json:"AvailableBytes",xml:"AvailableBytes"`
	UsageBytes     uint64 `json:"UsageBytes",xml:"UsageBytes"`
}

type MetricRecordNetworkStats struct {
	Interfaces []MetricRecordNetworkInterface `json:"Interfaces",xml:"Interfaces"`
}

type MetricRecordNetworkInterface struct {
	Name     string `json:"Name",xml:"Name"`
	RxBytes  uint64 `json:"RxBytes",xml:"RxBytes"`
	TxBytes  uint64 `json:"TxBytes",xml:"TxBytes"`
	RxErrors uint64 `json:"RxErrors",xml:"RxErrors"`
	TxErrors uint64 `json:"TxErrors",xml:"TxErrors"`
}
