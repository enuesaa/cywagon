package model

import "github.com/enuesaa/cywagon/internal/liblua"

type Conf struct {
	Host        string          `lua:"host"`
	Handler     liblua.Fn       `lua:"handler"`
}
