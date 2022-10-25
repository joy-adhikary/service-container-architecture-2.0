package main

import (
	"github.com/joy-adhikary/SC2.0/Provider"
	"github.com/joy-adhikary/SC2.0/container"
)

func main() {
	containers := &container.Container{
		Services: make(map[string]interface{}),
	}
	AuthorProvider := &Provider.AuthProvider{}
	AuthorProvider.Register(containers)

	BossProvider := &Provider.BossProvider{}
	BossProvider.Register(containers)

	CourseProvider := &Provider.CourseProvider{}
	CourseProvider.Register(containers)
}
