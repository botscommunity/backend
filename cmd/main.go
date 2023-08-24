package main

import (
  "log"
  "github.com/botscommunity/backend/internal/pkg/app"
)

func main() {
  server := app.NewGRPCServer()
  log.Fatalln(server.Run("3200"))
}
