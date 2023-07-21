package main

import (
	"log"

	. "github.com/adamgoose/screeps2go/runtime"
)

func main() {
	Run(func(g *Game) any {
		var spawn *StructureSpawn
		for _, s := range g.Spawns {
			spawn = s
			break
		}

		if len(g.Creeps) == 0 {
			log.Printf("No creeps. Spawning first worker.\n")
			spawn.SpawnCreep([]BodyPart{WORK, CARRY, MOVE}, "worker")
		}

		for _, creep := range g.Creeps {
			log.Printf("Working Creep %s\n", creep.Name)
			if creep.Store().GetFreeCapacity() > 0 {
				log.Printf("Has capacity. Harvesting.\n")

				sources := creep.Room.Find(int(FIND_SOURCES))
				if sources.Length() == 0 {
					log.Printf("Unable to find source. Skipping creep.\n")
					continue
				}
				soruce := sources.Index(0)

				if err := creep.Harvest(soruce); err == ERR_NOT_IN_RANGE {
					log.Printf("Not in range. Moving to source.\n")
					creep.MoveTo(soruce)
				}
			} else {
				log.Printf("No capacity. Transferring.\n")
				if err := creep.Transfer(spawn.Value(), string(RESOURCE_ENERGY)); err == ERR_NOT_IN_RANGE {
					log.Printf("Not in range. Moving to spawn.\n")
					creep.MoveTo(spawn.Value())
				}
			}
		}

		return nil
	})
}
