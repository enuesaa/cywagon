package main

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
)

func TestMain(m *testing.M) {
	infra.Default = infra.Container{}

	m.Run()
}
