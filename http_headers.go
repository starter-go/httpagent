package httpagent

import "strings"

// Headers ...
type Headers struct {
	table map[string]string
}

// Table ...
func (inst *Headers) Table() map[string]string {
	return inst.table
}

// Set ...
func (inst *Headers) Set(name, value string) {
	t := inst.table
	if t == nil {
		t = make(map[string]string)
		inst.table = t
	}
	name = inst.normalizeName(name)
	t[name] = value
}

// Get ...
func (inst *Headers) Get(name string) string {
	t := inst.table
	if t == nil {
		return ""
	}
	name = inst.normalizeName(name)
	return t[name]
}

// Remove ...
func (inst *Headers) Remove(name string) {
	name = inst.normalizeName(name)
	inst.Set(name, "")
}

func (inst *Headers) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}
