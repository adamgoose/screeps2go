package runtime

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall/js"
)

type Game struct {
	v      js.Value
	Spawns map[string]*StructureSpawn `json:"spawns"`
	Creeps map[string]*Creep          `json:"creeps"`
}

func (g Game) Value() js.Value {
	return g.v
}

func ParseGame(v js.Value) (*Game, error) {
	g := &Game{v: v}
	rawState := js.Global().Get("JSON").Call("stringify", v).String()
	if err := json.Unmarshal([]byte(rawState), &g); err != nil {
		return nil, err
	}

	return g, nil
}

func Run(f func(*Game) any) {
	c := make(chan struct{}, 0)
	js.Global().Set("_screeps2go", js.FuncOf(loop(f)))
	log.Println("wasm initialized")
	<-c
}

func loop(f func(*Game) any) func(js.Value, []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		g, err := ParseGame(js.Global().Get("Game"))
		if err != nil {
			log.Println(err)
			return nil
		}

		return f(g)
	}
}

func toInterfaceSlice[T any](s []T) []interface{} {
	ret := make([]interface{}, len(s))
	for i, v := range s {
		ret[i] = fmt.Sprintf("%v", v)
	}

	return ret
}
