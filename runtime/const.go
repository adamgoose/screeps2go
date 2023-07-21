package runtime

type Error int

const (
	OK                        Error = 0
	ERR_NOT_OWNER             Error = -1
	ERR_NO_PATH               Error = -2
	ERR_NAME_EXISTS           Error = -3
	ERR_BUSY                  Error = -4
	ERR_NOT_FOUND             Error = -5
	ERR_NOT_ENOUGH_RESOURCES  Error = -6
	ERR_NOT_ENOUGH_ENERGY     Error = -6
	ERR_INVALID_TARGET        Error = -7
	ERR_FULL                  Error = -8
	ERR_NOT_IN_RANGE          Error = -9
	ERR_INVALID_ARGS          Error = -10
	ERR_TIRED                 Error = -11
	ERR_NO_BODYPART           Error = -12
	ERR_NOT_ENOUGH_EXTENSIONS Error = -6
	ERR_RCL_NOT_ENOUGH        Error = -14
	ERR_GCL_NOT_ENOUGH        Error = -15
)

type BodyPart string

const (
	MOVE          BodyPart = "move"
	WORK          BodyPart = "work"
	CARRY         BodyPart = "carry"
	ATTACK        BodyPart = "attack"
	RANGED_ATTACK BodyPart = "ranged_attack"
	TOUGH         BodyPart = "tough"
	HEAL          BodyPart = "heal"
	CLAIM         BodyPart = "claim"
)

type (
	FindExitConstant       int
	FindRoomObjectConstant int
)

const (
	FIND_EXIT_TOP                   FindExitConstant       = 1
	FIND_EXIT_RIGHT                 FindExitConstant       = 3
	FIND_EXIT_BOTTOM                FindExitConstant       = 5
	FIND_EXIT_LEFT                  FindExitConstant       = 7
	FIND_EXIT                       FindExitConstant       = 10
	FIND_CREEPS                     FindRoomObjectConstant = 101
	FIND_MY_CREEPS                  FindRoomObjectConstant = 102
	FIND_HOSTILE_CREEPS             FindRoomObjectConstant = 103
	FIND_SOURCES_ACTIVE             FindRoomObjectConstant = 104
	FIND_SOURCES                    FindRoomObjectConstant = 105
	FIND_DROPPED_RESOURCES          FindRoomObjectConstant = 106
	FIND_STRUCTURES                 FindRoomObjectConstant = 107
	FIND_MY_STRUCTURES              FindRoomObjectConstant = 108
	FIND_HOSTILE_STRUCTURES         FindRoomObjectConstant = 109
	FIND_FLAGS                      FindRoomObjectConstant = 110
	FIND_CONSTRUCTION_SITES         FindRoomObjectConstant = 111
	FIND_MY_SPAWNS                  FindRoomObjectConstant = 112
	FIND_HOSTILE_SPAWNS             FindRoomObjectConstant = 113
	FIND_MY_CONSTRUCTION_SITES      FindRoomObjectConstant = 114
	FIND_HOSTILE_CONSTRUCTION_SITES FindRoomObjectConstant = 115
	FIND_MINERALS                   FindRoomObjectConstant = 116
	FIND_NUKES                      FindRoomObjectConstant = 117
	FIND_TOMBSTONES                 FindRoomObjectConstant = 118
	FIND_POWER_CREEPS               FindRoomObjectConstant = 119
	FIND_MY_POWER_CREEPS            FindRoomObjectConstant = 120
	FIND_HOSTILE_POWER_CREEPS       FindRoomObjectConstant = 121
	FIND_DEPOSITS                   FindRoomObjectConstant = 122
	FIND_RUINS                      FindRoomObjectConstant = 123
)

type ResourceConstant string

const (
	RESOURCE_ENERGY ResourceConstant = "energy"
	RESOURCE_POWER  ResourceConstant = "power"

	RESOURCE_HYDROGEN  ResourceConstant = "H"
	RESOURCE_OXYGEN    ResourceConstant = "O"
	RESOURCE_UTRIUM    ResourceConstant = "U"
	RESOURCE_LEMERGIUM ResourceConstant = "L"
	RESOURCE_KEANIUM   ResourceConstant = "K"
	RESOURCE_ZYNTHIUM  ResourceConstant = "Z"
	RESOURCE_CATALYST  ResourceConstant = "X"
	RESOURCE_GHODIUM   ResourceConstant = "G"

	RESOURCE_SILICON ResourceConstant = "silicon"
	RESOURCE_METAL   ResourceConstant = "metal"
	RESOURCE_BIOMASS ResourceConstant = "biomass"
	RESOURCE_MIST    ResourceConstant = "mist"

	RESOURCE_HYDROXIDE        ResourceConstant = "OH"
	RESOURCE_ZYNTHIUM_KEANITE ResourceConstant = "ZK"
	RESOURCE_UTRIUM_LEMERGITE ResourceConstant = "UL"

	RESOURCE_UTRIUM_HYDRIDE    ResourceConstant = "UH"
	RESOURCE_UTRIUM_OXIDE      ResourceConstant = "UO"
	RESOURCE_KEANIUM_HYDRIDE   ResourceConstant = "KH"
	RESOURCE_KEANIUM_OXIDE     ResourceConstant = "KO"
	RESOURCE_LEMERGIUM_HYDRIDE ResourceConstant = "LH"
	RESOURCE_LEMERGIUM_OXIDE   ResourceConstant = "LO"
	RESOURCE_ZYNTHIUM_HYDRIDE  ResourceConstant = "ZH"
	RESOURCE_ZYNTHIUM_OXIDE    ResourceConstant = "ZO"
	RESOURCE_GHODIUM_HYDRIDE   ResourceConstant = "GH"
	RESOURCE_GHODIUM_OXIDE     ResourceConstant = "GO"

	RESOURCE_UTRIUM_ACID        ResourceConstant = "UH2O"
	RESOURCE_UTRIUM_ALKALIDE    ResourceConstant = "UHO2"
	RESOURCE_KEANIUM_ACID       ResourceConstant = "KH2O"
	RESOURCE_KEANIUM_ALKALIDE   ResourceConstant = "KHO2"
	RESOURCE_LEMERGIUM_ACID     ResourceConstant = "LH2O"
	RESOURCE_LEMERGIUM_ALKALIDE ResourceConstant = "LHO2"
	RESOURCE_ZYNTHIUM_ACID      ResourceConstant = "ZH2O"
	RESOURCE_ZYNTHIUM_ALKALIDE  ResourceConstant = "ZHO2"
	RESOURCE_GHODIUM_ACID       ResourceConstant = "GH2O"
	RESOURCE_GHODIUM_ALKALIDE   ResourceConstant = "GHO2"

	RESOURCE_CATALYZED_UTRIUM_ACID        ResourceConstant = "XUH2O"
	RESOURCE_CATALYZED_UTRIUM_ALKALIDE    ResourceConstant = "XUHO2"
	RESOURCE_CATALYZED_KEANIUM_ACID       ResourceConstant = "XKH2O"
	RESOURCE_CATALYZED_KEANIUM_ALKALIDE   ResourceConstant = "XKHO2"
	RESOURCE_CATALYZED_LEMERGIUM_ACID     ResourceConstant = "XLH2O"
	RESOURCE_CATALYZED_LEMERGIUM_ALKALIDE ResourceConstant = "XLHO2"
	RESOURCE_CATALYZED_ZYNTHIUM_ACID      ResourceConstant = "XZH2O"
	RESOURCE_CATALYZED_ZYNTHIUM_ALKALIDE  ResourceConstant = "XZHO2"
	RESOURCE_CATALYZED_GHODIUM_ACID       ResourceConstant = "XGH2O"
	RESOURCE_CATALYZED_GHODIUM_ALKALIDE   ResourceConstant = "XGHO2"

	RESOURCE_OPS ResourceConstant = "ops"

	RESOURCE_UTRIUM_BAR    ResourceConstant = "utrium_bar"
	RESOURCE_LEMERGIUM_BAR ResourceConstant = "lemergium_bar"
	RESOURCE_ZYNTHIUM_BAR  ResourceConstant = "zynthium_bar"
	RESOURCE_KEANIUM_BAR   ResourceConstant = "keanium_bar"
	RESOURCE_GHODIUM_MELT  ResourceConstant = "ghodium_melt"
	RESOURCE_OXIDANT       ResourceConstant = "oxidant"
	RESOURCE_REDUCTANT     ResourceConstant = "reductant"
	RESOURCE_PURIFIER      ResourceConstant = "purifier"
	RESOURCE_BATTERY       ResourceConstant = "battery"

	RESOURCE_COMPOSITE ResourceConstant = "composite"
	RESOURCE_CRYSTAL   ResourceConstant = "crystal"
	RESOURCE_LIQUID    ResourceConstant = "liquid"

	RESOURCE_WIRE       ResourceConstant = "wire"
	RESOURCE_SWITCH     ResourceConstant = "switch"
	RESOURCE_TRANSISTOR ResourceConstant = "transistor"
	RESOURCE_MICROCHIP  ResourceConstant = "microchip"
	RESOURCE_CIRCUIT    ResourceConstant = "circuit"
	RESOURCE_DEVICE     ResourceConstant = "device"

	RESOURCE_CELL     ResourceConstant = "cell"
	RESOURCE_PHLEGM   ResourceConstant = "phlegm"
	RESOURCE_TISSUE   ResourceConstant = "tissue"
	RESOURCE_MUSCLE   ResourceConstant = "muscle"
	RESOURCE_ORGANOID ResourceConstant = "organoid"
	RESOURCE_ORGANISM ResourceConstant = "organism"

	RESOURCE_ALLOY      ResourceConstant = "alloy"
	RESOURCE_TUBE       ResourceConstant = "tube"
	RESOURCE_FIXTURES   ResourceConstant = "fixtures"
	RESOURCE_FRAME      ResourceConstant = "frame"
	RESOURCE_HYDRAULICS ResourceConstant = "hydraulics"
	RESOURCE_MACHINE    ResourceConstant = "machine"

	RESOURCE_CONDENSATE  ResourceConstant = "condensate"
	RESOURCE_CONCENTRATE ResourceConstant = "concentrate"
	RESOURCE_EXTRACT     ResourceConstant = "extract"
	RESOURCE_SPIRIT      ResourceConstant = "spirit"
	RESOURCE_EMANATION   ResourceConstant = "emanation"
	RESOURCE_ESSENCE     ResourceConstant = "essence"
)
