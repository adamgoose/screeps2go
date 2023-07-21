// This file is generated by screeps2go. DO NOT EDIT!
package runtime

import "syscall/js"

type Store struct {
	CreepName string `json:""`
}

func (s Store) Value() js.Value {
	return js.Global().Get("Game").Get("creeps").Get(s.CreepName).Get("store")
}

func (s Store) GetFreeCapacity() int {
	return s.Value().Call("getFreeCapacity").Int()
}
