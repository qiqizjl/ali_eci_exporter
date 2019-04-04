package main

import (
	"ali_eci_exporter/exporter"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func getEnvBool(key string) (envValBool bool) {
	if envVal, ok := os.LookupEnv(key); ok {
		envValBool, _ = strconv.ParseBool(envVal)
	}
	return
}

func main() {
	var (
		listenAddress        = flag.String("web.listen-address", getEnv("REDIS_EXPORTER_WEB_LISTEN_ADDRESS", ":8080"), "Address to listen on for web interface and telemetry.")
		metricPath           = flag.String("web.telemetry-path", getEnv("REDIS_EXPORTER_WEB_TELEMETRY_PATH", "/metrics"), "Path under which to expose metrics.")
		aliCloudAccessKey    = flag.String("alicloud.accesskey", getEnv("ALICLOUD_ACCESSKEY", ""), "Alibaba Cloud Access Key")
		aliCloudAccessSecert = flag.String("alicloud.accesssecret", getEnv("ALICLOUD_ACCESSSECRET", ""), "Alibaba Cloud Access Secret")
		aliCloudRegion       = flag.String("alicloud.region", getEnv("ALICLOUD_REGION", ""), "Alibaba Cloud Region")
	)
	flag.Parse()

	eciClient, _ := exporter.MewExporter(*aliCloudAccessKey, *aliCloudAccessSecert, *aliCloudRegion)
	registry := prometheus.NewRegistry()
	registry.MustRegister(eciClient)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.Handle(*metricPath, handler)
	//http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
