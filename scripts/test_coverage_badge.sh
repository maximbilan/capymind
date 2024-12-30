#!/bin/bash

go test -coverprofile coverage.out -covermode count -coverpkg=./... -v ./...
go-ignore-cov --file coverage.out
go tool cover -func=coverage.out

COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

# Green color if coverage is 75% or more
# Yellow color if coverage is 50% or more (but white text should be visible)
# Otherwise, red color
if [ $(echo "$COVERAGE >= 75" | bc) -eq 1 ]; then
  COLOR="#4c1"
elif [ $(echo "$COVERAGE >= 50" | bc) -eq 1 ]; then
  COLOR="#dfb317"
else
  COLOR="#e05d44"
fi

echo "<svg xmlns='http://www.w3.org/2000/svg' width='150' height='20'>
  <rect width='150' height='20' fill='#555' rx='5' ry='5'/>
  <defs>
    <clipPath id='clip-rounded'>
      <rect width='150' height='20' rx='5' ry='5'/>
    </clipPath>
  </defs>
  <rect x='100' width='50' height='20' fill='${COLOR}' clip-path='url(#clip-rounded)'/>
  <text x='50' y='14' fill='#fff' font-family='Verdana' font-size='11' text-anchor='middle'>Test Coverage</text>
  <text x='125' y='14' fill='#fff' font-family='Verdana' font-size='11' text-anchor='middle'>${COVERAGE}%</text>
</svg>" > .badges/test_coverage.svg
