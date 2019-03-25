package goaxis

import (
    "strings"
)

type caller struct{
    modules []string
}

func newCaller(master string) ICaller {
    p := &caller{modules: make([]string, 0)}
    p.modules = append(p.modules, master)
    return p
}

func (p *caller) Last() string {
    return p.modules[len(p.modules) - 1]
}

func (p *caller) First() string {
    return p.modules[0]
}

func (p *caller) Push(names ...string) ICaller {
    p.modules = append(p.modules, names...)
    return p
}

func (p *caller) Format(split string) string {
    return strings.Join(p.modules, split)
}