package tasks

import (
	"inventoryApp/app/models"

	"github.com/gobuffalo/pop"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {
	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		tx := c.Value("tx").(*pop.Connection)

		lines := models.Lines{}
		if err := tx.Eager().All(&lines); err != nil {
			return err
		}

		if err := tx.Eager().Create(lines); err != nil {
			return err
		}

		return nil

	})

})
