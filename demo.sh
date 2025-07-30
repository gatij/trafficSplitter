#!/bin/bash

echo "Making 10 requests to the user service through Envoy..."
echo "Expected: ~7 requests should go to v1 and ~3 requests to v2"
echo ""

for i in {1..10}; do
  echo "Request $i:"
  curl -s http://localhost:8080/users/1 | jq
  echo ""
  sleep 0.5
done

echo "Summary of traffic distribution from Envoy admin console:"
curl -s http://localhost:9901/stats | grep "cluster.service_v" | grep "upstream_rq_completed"
