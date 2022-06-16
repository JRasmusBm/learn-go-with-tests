package main

import (
	"os"
	"time"

	"github.com/JRasmusBm/learn-go-with-tests/pkg/clockface"
)

func main() {
	origin := 150.0
	scale := 60.0

	clockface.New(origin, scale).WriteSvg(os.Stdout, time.Now())
}
