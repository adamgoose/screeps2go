package main

type (
	Error            int
	Value            int
	BodyPart         string
	ResourceConstant string
)

type Creep struct {
	v    ICreep `js:"Game.creeps.%s" key:"Name"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Room Room   `json:"room"`
}

type ICreep interface {
	Suicide() Error
	MoveTo(target Value) Error
	Say(message string, public bool) Error
	Harvest(target Value) Error
	Transfer(target Value, resource string) Error
}

type StructureSpawn struct {
	v    IStructureSpawn `js:"Game.spawns.%s" key:"Name"`
	ID   string          `json:"id"`
	Name string          `json:"name"`
}

type IStructureSpawn interface {
	IsActive() bool
	SpawnCreep(body []BodyPart, name string) Error
}

type Room struct {
	v    IRoom  `js:"Game.rooms.%s" key:"Name"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IRoom interface {
	Find(t int) Value
}

type Store struct {
	v         IStore `js:"Game.creeps.%s.store" key:"CreepName"`
	CreepName string
}

type IStore interface {
	GetFreeCapacity() int
}
