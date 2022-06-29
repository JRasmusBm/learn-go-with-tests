package main

import (
	"os"
	"time"

	"github.com/JRasmusBm/learn-go-with-tests/pkg/clockface"
)

func main() {
	origin := 150.0
	scale := 100.0

	clockface.New(origin, scale).WriteSVG(os.Stdout, time.Now())
}
