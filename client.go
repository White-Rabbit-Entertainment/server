package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Vec2d struct {
  X float64 `json:"x"`
  Y float64 `json:"y"`
}

type Player struct {
  Id int `json:"id"`
  Position Vec2d `json:"position"`
}

func (p *Player) move(distanceX, distanceY float64) {
  // Request to server to move the player
  translation := Vec2d{distanceX, distanceY}
  jsonTranslation, _ := json.Marshal(translation)
  
  id := strconv.Itoa(p.Id)
  req, _ := http.NewRequest("POST", "http://localhost:8080/player/" + id + "/move", bytes.NewBuffer(jsonTranslation))

  client := http.Client{
	Timeout: time.Second * 2, // Timeout after 2 seconds
  }

  res, _ := client.Do(req)

  body, _ := ioutil.ReadAll(res.Body)
  fmt.Printf("%s\n", body)

  position := Vec2d{}
  json.Unmarshal(body, &position)

  p.Position.X = position.X
  p.Position.Y = position.Y
}


func (p *Player) getPlayerPosition() {
  // Make request to server to get the player location
  id := strconv.Itoa(p.Id)
  req, _ := http.NewRequest("GET", "http://localhost:8080/player/" + id + "/position", nil)
  client := http.Client{
	Timeout: time.Second * 2, // Timeout after 2 seconds
  }
  res, _ := client.Do(req)

  body, _ := ioutil.ReadAll(res.Body)
  fmt.Printf("%s\n", body)

  position := Vec2d{}
  json.Unmarshal(body, &position)

  p.Position.X = position.X
  p.Position.Y = position.Y
}

func (p *Player) init() {
  jsonPlayer, _ := json.Marshal(p)
  req, _ := http.NewRequest("POST", "http://localhost:8080/create", bytes.NewBuffer(jsonPlayer))
  client := http.Client{
	Timeout: time.Second * 2, // Timeout after 2 seconds
  }

  res, _ := client.Do(req)
  body, _ := ioutil.ReadAll(res.Body)
  fmt.Printf("%s\n", body)
  player := Player{}
  json.Unmarshal(body, &player)
 
  p.Id = player.Id
  fmt.Println(player.Id)
}

func main() {
  fmt.Println("Hello")
  player := Player{Position: Vec2d{10, 10}}
  player.init()
  player.getPlayerPosition()
  player.move(1, 1)
  player.getPlayerPosition()
}
