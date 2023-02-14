package lines

import (
	"inventoryApp/app/models"
	"inventoryApp/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func List(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	lines := models.Lines{}

	err := tx.Eager().All(&lines)

	if err != nil {
		return err
	}

	c.Set("lines", lines)
	return c.Render(http.StatusOK, r.HTML("line/line.plush.html"))
}

func Show(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	line, err := findLine(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding line error"))
	}

	c.Set("line", line)
	return c.Render(http.StatusOK, r.HTML("line/details.plush.html"))
}

func Create(c buffalo.Context) error {
	line := models.Line{}

	if err := c.Bind(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "add line bind error"))
	}
	tx := c.Value("tx").(*pop.Connection)

	if verrs := line.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("line", line)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/line/create.plush.html"))
	}

	if err := tx.Create(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "create line error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

func New(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	lines := models.Line{}
	users := []models.User{}

	err := tx.Eager().All(&users)

	if err != nil {
		return err
	}

	c.Set("users", users)
	c.Set("line", lines)
	return c.Render(http.StatusOK, r.HTML("line/create.plush.html"))

}

func Edit(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	line, err := findLine(c, id)

	if err != nil {

		return errors.WithStack(errors.Wrap(err, "finding line error"))
	}

	c.Set("line", line)
	return c.Render(http.StatusUnprocessableEntity, r.HTML("line/edit.plush.html"))
}

func findLine(c buffalo.Context, id uuid.UUID) (models.Line, error) {
	tx := c.Value("tx").(*pop.Connection)

	line := models.Line{}

	if err := tx.Eager().Find(&line, id); err != nil {
		return line, err
	}

	return line, nil
}

func Update(c buffalo.Context) error {

	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	line, err := findLine(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding line error"))
	}

	if err := c.Bind(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "add line bind error"))
	}

	if verrs := line.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("line", line)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/line/edit.plush.html"))
	}

	if err := tx.Update(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "create line error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

func ChangeStatus(c buffalo.Context) error {

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	line, err := findLine(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding line error"))
	}

	if err := c.Bind(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "add line bind error"))
	}

	if err := tx.Update(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "create line error"))
	}

	return c.Redirect(http.StatusSeeOther, "rootPath()")
}

func Delete(c buffalo.Context) error {

	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	line, err := findLine(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding line error"))
	}

	if err := tx.Destroy(&line); err != nil {
		return errors.WithStack(errors.Wrap(err, "destroy line error"))
	}
	return c.Redirect(http.StatusFound, "rootPath()")
}

func Test(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("line/test.plush.html"))

}
