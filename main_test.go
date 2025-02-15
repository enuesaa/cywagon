package main

import (
	"testing"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	infra.Default = infra.Container{}

	m.Run()
}

func TestDoNotUseDefaultContainer(t *testing.T) {
	assert.Equal(t, infra.Default.Cmd, nil)
}
