package lua

import (
	"os"
)

var (
	CompatVarArg     = true
	FieldsPerFlush   = 50
	RegistrySize     = 256 * 20
	RegistryGrowStep = 32
	CallStackSize    = 256
	MaxTableGetLoop  = 100
	MaxArrayIndex    = 67108864
)

type LNumber float64

const (
	LNumberBit        = 64
	LNumberScanFormat = "%f"
	LuaVersion        = "Lua 5.1"
)

var (
	LuaPath        = "LUA_PATH"
	LuaLDir        string
	LuaPathDefault string
	LuaOS          string
)

func init() {
	if os.PathSeparator == '/' { // unix-like
		LuaOS = "unix"
		LuaLDir = "/usr/local/share/lua/5.1"
		LuaPathDefault = "./?.lua;" + LuaLDir + "/?.lua;" + LuaLDir + "/?/init.lua"
	} else { // windows
		LuaOS = "windows"
		LuaLDir = "!\\lua"
		LuaPathDefault = ".\\?.lua;" + LuaLDir + "\\?.lua;" + LuaLDir + "\\?\\init.lua"
	}
}
