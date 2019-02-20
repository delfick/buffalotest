package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/lifx/test/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
