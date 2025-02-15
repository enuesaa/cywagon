package main

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
)

func TestMain(m *testing.M) {
	infra.SetupMock()
	m.Run()
}
