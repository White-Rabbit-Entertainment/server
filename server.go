package main

import (
    "fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Vec2d struct {
  X float64 `json:"x"`
  Y float64 `json:"y"`
}

type ServerPlayer struct {
  Id int `json:"id"`
  Position Vec2d `json:"position"`
}

type Env struct {
    players []*ServerPlayer
}

func (e *Env) getPlayer(id int) *ServerPlayer {
  for _, player := range e.players {
    if player.Id == id {
      return player
    }
  }
  return nil
}

func (e *Env) nextId() int {
  maxId := 0
  for _, player := range e.players {
    if player.Id > maxId {
      maxId = player.Id
    }
  }
  return maxId + 1
}

func SetupRouter(env *Env) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
    r.GET("/player/:id/position", env.getPlayerPosition)
    r.POST("/player/:id/move", env.updatePlayerPosition)
    r.POST("/create", env.createPlayer)

    return r
}

func (e *Env) getPlayerPosition(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  c.JSON(200, gin.H{
    "x":e.getPlayer(id).Position.X,
    "y":e.getPlayer(id).Position.Y,
  })
}

func (e *Env) updatePlayerPosition(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var translation Vec2d 
  c.BindJSON(&translation)

  e.getPlayer(id).Position.X += translation.X
  e.getPlayer(id).Position.Y += translation.Y
  
  fmt.Println(e.getPlayer(id).Position.X)
  fmt.Println(e.getPlayer(id).Position.Y)
}


func (e *Env) createPlayer(c *gin.Context) {
  var newPlayer ServerPlayer
  c.BindJSON(&newPlayer)

  newPlayer.Id = e.nextId()

  e.players = append(e.players, &newPlayer)
  c.JSON(200, newPlayer)
}

func main() {
  fmt.Println("Here we go")
  env := &Env{}
  r := SetupRouter(env)
  r.Run(":8080")
}

