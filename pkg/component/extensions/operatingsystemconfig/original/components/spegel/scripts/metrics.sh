#!/bin/bash
set -o nounset
set -o pipefail

function scrape_spegel_metrics {
  while true; do
    curl --request GET -sL \
         --url 'http://localhost:9590/metrics'\
         --output "$output_file.tmp"
    mv "$output_file.tmp" "$output_file"
    sleep $SLEEP_SECONDS
  done
}

output_file="var/lib/node-exporter/textfile-collector/spegel.prom"
SLEEP_SECONDS=5
echo "Start scraping spegel metrics"
scrape_spegel_metrics
