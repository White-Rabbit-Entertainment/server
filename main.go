package main

import (
    models "server/models"
)

type Env struct {
    players []*models.Player
}

func main() {
  env := &Env{}
  r := SetupRouter(env)
  r.Run(":8080")
}
