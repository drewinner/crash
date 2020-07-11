package g

import "path/filepath"

var Modules map[string]bool
var BinOf map[string]string
var cfgOf map[string]string
var ModuleApps map[string]string
var logpathOf map[string]string
var PidOf map[string]string
var AllModulesInOrder []string

func init() {
	Modules = map[string]bool{
		"admin":       true,
		"api":         true,
		"transaction": true,
	}

	BinOf = map[string]string{
		"admin":       "./admin/bin/crash-admin",
		"api":         "./api/bin/crash-api",
		"transaction": "./transaction/bin/crash-transaction",
	}

	cfgOf = map[string]string{
		"admin":       "./admin/config/cfg.json",
		"api":         "./api/config/cfg.json",
		"transaction": "./transaction/config/cfg.json",
	}

	ModuleApps = map[string]string{
		"admin":       "crash-admin",
		"api":         "crash-api",
		"transaction": "crash-transaction",
	}

	logpathOf = map[string]string{
		"admin":       "./admin/logs/admin.log",
		"api":         "./api/logs/api.log",
		"transaction": "./transaction/logs/transaction.log",
	}

	PidOf = map[string]string{
		"admin":       "<NOT SET>",
		"api":         "<NOT SET>",
		"transaction": "<NOT SET>",
	}

	// Modules are deployed in this order
	AllModulesInOrder = []string{
		"admin",
		"api",
		"transaction",
	}
}

func Bin(name string) string {
	p, _ := filepath.Abs(BinOf[name])
	return p
}

func Cfg(name string) string {
	p, _ := filepath.Abs(cfgOf[name])
	return p
}

func LogPath(name string) string {
	p, _ := filepath.Abs(logpathOf[name])
	return p
}

func LogDir(name string) string {
	d, _ := filepath.Abs(filepath.Dir(logpathOf[name]))
	return d
}
