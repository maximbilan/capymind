#!/bin/bash

go test -coverprofile=coverage.out ./...
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

COLOR="red"
if (( $(echo "$COVERAGE >= 90" | bc -l) )); then
  COLOR="brightgreen"
elif (( $(echo "$COVERAGE >= 75" | bc -l) )); then
  COLOR="yellow"
fi

echo "<svg xmlns='http://www.w3.org/2000/svg' width='120' height='20'>
  <rect width='60' height='20' fill='#555'/>
  <rect x='60' width='60' height='20' fill='#${COLOR}'/>
  <text x='30' y='14' fill='#fff' font-family='Verdana' font-size='11' text-anchor='middle'>coverage</text>
  <text x='90' y='14' fill='#fff' font-family='Verdana' font-size='11' text-anchor='middle'>${COVERAGE}%</text>
</svg>" > .badges/test_coverage.svg