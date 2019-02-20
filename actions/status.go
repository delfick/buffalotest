package actions

import "github.com/gobuffalo/buffalo"

// Status returns the status
func (t *stuff) Status(c buffalo.Context) error {
	if t.thing == nil {
		err := "The application hasn't started properly and I don't have access to the thing"
		return c.Render(200, r.JSON(map[string]string{"error": err}))
	}

	return c.Render(200, r.JSON(map[string]bool{"reachable": t.thing.Reachable()}))
}
