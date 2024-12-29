#!/bin/bash

go test -coverprofile coverage.out -covermode count -coverpkg=./... -v ./...
go-ignore-cov --file coverage.out
go tool cover -func=coverage.out

COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

if [ "$COVERAGE" -ge 80 ]; then
  COLOR="brightgreen" # High coverage
elif [ "$COVERAGE" -ge 50 ]; then
  COLOR="yellow" # Medium coverage
elif [ "$COVERAGE" -ge 25 ]; then
  COLOR="orange" # Low-medium coverage
else
  COLOR="red" # Very low coverage
fi

echo "<svg xmlns='http://www.w3.org/2000/svg' width='150' height='20'>
  <rect width='100' height='20' fill='#555'/>
  <rect x='100' width='50' height='20' fill='${COLOR}'/>
  <text x='50' y='14' fill='#fff' font-family='Verdana' font-size='11' text-anchor='middle'>Test Coverage</text>
  <text x='125' y='14' fill='#fff' font-family='Verdana' font-size='11' text-anchor='middle'>${COVERAGE}%</text>
</svg>" > .badges/test_coverage.svg