package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/config"
	"os"
	"os/exec"
)

func main() {
	direction := flag.String("d", "up", "Migration direction (up or down)")
	flag.Parse()
	if *direction != "up" && *direction != "down" {
		logrus.Fatal("Invalid direction")
	}

	cfg := config.NewConfig()
	url := cfg.DB.GetConnectionDSN()
	cmd := exec.Command(
		"migrate", "-path", "./migrations", "-database", url, *direction,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logrus.Fatal(err)
	}
}
