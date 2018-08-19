#!/usr/bin/env bash
set -e
./test-tdd.sh

echo "Ball Clock Simulation - Mode 1: 30 Balls"
curl -X POST \
  http://localhost:3000/ \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 54238cb5-de9a-4745-9087-980dbd2b6c58' \
  -d '{
	"BallCount": 30
}' | jq '.'
echo ""
echo "Ball Clock Simulation - Mode 1: 45 Balls"
curl -X POST \
  http://localhost:3000/ \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 54238cb5-de9a-4745-9087-980dbd2b6c58' \
  -d '{
	"BallCount": 45
}' | jq '.'
echo ""
echo "Ball Clock Simulation - Mode 2: 30 Balls & 325 Iterations"
curl -X POST \
  http://localhost:3000/ \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: b1f62898-b03f-4ef1-ae0c-b153dbe2d629' \
  -d '{
	"BallCount": 30,
	"IterationCount": 325
}' | jq '.'
