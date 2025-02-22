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
	// Since the default DI container is overwritten by TestMain(), this is expected to be nil.
	assert.Equal(t, infra.Default.Cmd, nil)
}
