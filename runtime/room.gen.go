// This file is generated by screeps2go. DO NOT EDIT!
package runtime

import "syscall/js"

type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (r Room) Value() js.Value {
	return js.Global().Get("Game").Get("rooms").Get(r.Name)
}

func (r Room) Find(p0 int) js.Value {
	return r.Value().Call("find", p0)
}
