package agent

import (
	"fmt"
	"time"
)

type Foreman struct {
	Name string
}

func NewForeman() *Foreman {
	return &Foreman{
		Name: "Foreman",
	}
}

func (f *Foreman) Start() {
	fmt.Printf("[%s] Starting orchestration loop...\n", f.Name)
	for {
		// Basic orchestration logic
		time.Sleep(10 * time.Second)
	}
}

func StartForeman() {
	f := NewForeman()
	f.Start()
}
