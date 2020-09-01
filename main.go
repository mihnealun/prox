// prox API.
package main

import (
	"github.com/mihnealun/prox/infrastructure"

	"github.com/mihnealun/prox/infrastructure/container"
)

func main() {
	containerInstance, err := container.GetInstance()
	if err != nil {
		panic(err.Error())
	}

	err = infrastructure.Start(containerInstance)
	if err != nil {
		panic(err.Error())
	}
}
