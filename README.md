<p align="center">
  <img alt="logo" src="https://github.com/idoqo/starlink_exporter/raw/main/.docs/assets/logo.jpg" height="100" />
  <h3 align="center">Starlink Prometheus Exporter</h3>
</p>

---
A [Starlink](https://www.starlink.com/) exporter for Prometheus. Not affiliated with or acting on behalf of Starlink(™)

[![goreleaser](https://github.com/idoqo/starlink_exporter/actions/workflows/release.yaml/badge.svg)](https://github.com/idoqo/starlink_exporter/actions/workflows/release.yaml)
[![build](https://github.com/idoqo/starlink_exporter/actions/workflows/build.yaml/badge.svg)](https://github.com/idoqo/starlink_exporter/actions/workflows/build.yaml)
[![License](https://img.shields.io/github/license/idoqo/starlink_exporter)](/LICENSE)
[![Release](https://img.shields.io/github/release/idoqo/starlink_exporter.svg)](https://github.com/idoqo/starlink_exporter/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/idoqo/starlink_exporter)
![os/arch](https://img.shields.io/badge/os%2Farch-amd64-yellow)
![os/arch](https://img.shields.io/badge/os%2Farch-arm64-yellow)
![os/arch](https://img.shields.io/badge/os%2Farch-armv7-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/idoqo/starlink_exporter)](https://goreportcard.com/report/github.com/idoqo/starlink_exporter)

If you would like to see this in action in your Starlink network, you can run the docker-compose file at https://github.com/idoqo/starlink_exporter/tree/main/demo/docker-compose.yml. 
It includes this exporter, Grafana and Prometheus in one Docker Compose file.

## Usage:

### Flags

`starlink_exporter` is configured through the use of optional command line flags

```bash
$ ./starlink_exporter --help
Usage of starlink_exporter
  -dish-address string
        IP address and port to reach dish (default "192.168.100.1:9200")
  -router-address string
        IP address and port to reach router (default "192.168.1.1:9000")
  -port string
        listening port to expose metrics on (default "9817")

```

### Binaries

For pre-built binaries please take a look at the [releases](https://github.com/danopstech/starlink_exporter/releases).

```bash
./starlink_exporter [flags]
```

### Docker

Docker Images can be found at [GitHub Container Registry](https://github.com/orgs/idoqo/packages/container/package/starlink_exporter).

Example:
```bash
docker pull ghcr.io/idoqo/starlink_exporter:latest

docker run \
  -p 9817:9817 \
  ghcr.io/idoqo/starlink_exporter:latest [flags]
```

### Setup Prometheus to scrape `starlink_exporter`

Configure [Prometheus](https://prometheus.io/) to scrape metrics from localhost:9817/metrics

```yaml
...
scrape_configs
    - job_name: starlink
      scrape_interval: 3s
      scrape_timeout:  3s
      static_configs:
        - targets: ['localhost:9817']
...
```

## Exported Metrics:

```text
# HELP starlink_dish_alert_install_pending Installation Pending
# TYPE starlink_dish_alert_install_pending gauge
# HELP starlink_dish_alert_is_heating Is Heating
# TYPE starlink_dish_alert_is_heating gauge
# HELP starlink_dish_alert_mast_not_near_vertical Status of mast position
# TYPE starlink_dish_alert_mast_not_near_vertical gauge
# HELP starlink_dish_alert_motors_stuck Status of motor stuck
# TYPE starlink_dish_alert_motors_stuck gauge
# HELP starlink_dish_alert_roaming Status of roaming
# TYPE starlink_dish_alert_roaming gauge
# HELP starlink_dish_alert_slow_eth_speeds Status of ethernet
# TYPE starlink_dish_alert_slow_eth_speeds gauge
# HELP starlink_dish_alert_thermal_shutdown Status of thermal shutdown
# TYPE starlink_dish_alert_thermal_shutdown gauge
# HELP starlink_dish_alert_thermal_throttle Status of thermal throttling
# TYPE starlink_dish_alert_thermal_throttle gauge
# HELP starlink_dish_alert_unexpected_location Status of location
# TYPE starlink_dish_alert_unexpected_location gauge
# HELP starlink_dish_anti_rollback_version Starlink Dish Anti Rollback Version.
# TYPE starlink_dish_anti_rollback_version counter
# HELP starlink_dish_boot_count Starlink Dish boot count.
# TYPE starlink_dish_boot_count counter
# HELP starlink_dish_bore_sight_azimuth_deg azimuth in degrees
# TYPE starlink_dish_bore_sight_azimuth_deg gauge
# HELP starlink_dish_bore_sight_elevation_deg elevation in degrees
# TYPE starlink_dish_bore_sight_elevation_deg gauge
# HELP starlink_dish_currently_obstructed Status of view of the sky
# TYPE starlink_dish_currently_obstructed gauge
# HELP starlink_dish_dish_stow_requested stow requested
# TYPE starlink_dish_dish_stow_requested gauge
# HELP starlink_dish_downlink_throughput_bytes Amount of bandwidth in bytes per second download
# TYPE starlink_dish_downlink_throughput_bytes gauge
# HELP starlink_dish_eth_speed ethernet speed
# TYPE starlink_dish_eth_speed untyped
# HELP starlink_dish_first_nonempty_slot_seconds Seconds to next non empty slot
# TYPE starlink_dish_first_nonempty_slot_seconds gauge
# HELP starlink_dish_fraction_obstruction_ratio Percentage of obstruction
# TYPE starlink_dish_fraction_obstruction_ratio gauge
# HELP starlink_dish_gps_sats Number of GPS Sats.
# TYPE starlink_dish_gps_sats gauge
# HELP starlink_dish_gps_valid GPS Status.
# TYPE starlink_dish_gps_valid gauge
# HELP starlink_dish_info Running software versions and IDs of hardware
# TYPE starlink_dish_info gauge
# HELP starlink_dish_info_debug Debug Dish Info
# TYPE starlink_dish_info_debug gauge
# HELP starlink_dish_is_dev Starlink Dish is Dev.
# TYPE starlink_dish_is_dev gauge
# HELP starlink_dish_is_hit Starlink Dish is Hit.
# TYPE starlink_dish_is_hit gauge
# HELP starlink_dish_outage_did_switch Starlink Dish Outage Information
# TYPE starlink_dish_outage_did_switch gauge
# HELP starlink_dish_outage_duration Starlink Dish Outage Information
# TYPE starlink_dish_outage_duration gauge
# HELP starlink_dish_pop_ping_drop_ratio Percent of pings dropped
# TYPE starlink_dish_pop_ping_drop_ratio gauge
# HELP starlink_dish_pop_ping_latency_seconds Latency of connection in seconds
# TYPE starlink_dish_pop_ping_latency_seconds gauge
# HELP starlink_dish_prolonged_obstruction_duration_seconds Average in seconds of prolonged obstructions
# TYPE starlink_dish_prolonged_obstruction_duration_seconds gauge
# HELP starlink_dish_prolonged_obstruction_interval_seconds Average prolonged obstruction interval in seconds
# TYPE starlink_dish_prolonged_obstruction_interval_seconds gauge
# HELP starlink_dish_prolonged_obstruction_valid Average prolonged obstruction is valid
# TYPE starlink_dish_prolonged_obstruction_valid gauge
# HELP starlink_dish_scrape_duration_seconds Time to scrape metrics from starlink dish
# TYPE starlink_dish_scrape_duration_seconds gauge
# HELP starlink_dish_software_partitions_equal Starlink Dish Software Partitions Equal.
# TYPE starlink_dish_software_partitions_equal gauge
# HELP starlink_dish_up Was the last query of Starlink dish successful.
# TYPE starlink_dish_up gauge
# HELP starlink_dish_uplink_throughput_bytes Amount of bandwidth in bytes per second upload
# TYPE starlink_dish_uplink_throughput_bytes gauge
# HELP starlink_dish_uptime_seconds Dish running time
# TYPE starlink_dish_uptime_seconds counter
# HELP starlink_dish_valid_seconds Unknown
# TYPE starlink_dish_valid_seconds counter
# HELP starlink_wifi_connected_clients_count Number of connected WiFi clients
# TYPE starlink_wifi_connected_clients_count gauge
# HELP starlink_wifi_connected_clients_info Connected Clients Info
# TYPE starlink_wifi_connected_clients_info gauge
```

## Example Grafana Dashboard:
You can import the current dashboard from ()

<p align="center">
	<img src="https://github.com/idoqo/starlink_exporter/raw/main/.docs/assets/dashboard-screenshot.png" width="95%">
</p>
