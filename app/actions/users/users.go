package users

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
	users := models.Users{}
	err := tx.All(&users)

	if err != nil {
		return err
	}

	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("user/user.plush.html"))
}

func User(c buffalo.Context) error {
	user := models.User{}
	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("user/create.plush.html"))
}

func Create(c buffalo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return errors.WithStack(errors.Wrap(err, "add user bind error"))
	}
	tx := c.Value("tx").(*pop.Connection)

	if verrs := user.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/user/create.plush.html"))
	}

	if err := tx.Eager().Create(&user); err != nil {
		return errors.WithStack(errors.Wrap(err, "create user error"))
	}

	return c.Redirect(http.StatusSeeOther, "usersPath()")
}

func Show(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	user, err := findUser(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding user error"))
	}

	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("user/details.plush.html"))
}

func findUser(c buffalo.Context, id uuid.UUID) (models.User, error) {
	tx := c.Value("tx").(*pop.Connection)

	user := models.User{}

	if err := tx.Find(&user, id); err != nil {
		return user, err
	}

	return user, nil
}

func Edit(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	user, err := findUser(c, id)

	if err != nil {

		return errors.WithStack(errors.Wrap(err, "finding user error"))
	}

	c.Set("user", user)
	return c.Render(http.StatusUnprocessableEntity, r.HTML("user/edit.plush.html"))
}

func Update(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	user, err := findUser(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding user error"))
	}

	if err := c.Bind(&user); err != nil {
		return errors.WithStack(errors.Wrap(err, "add user bind error"))
	}

	if verrs := user.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/user/edit.plush.html"))
	}

	if err := tx.Update(&user); err != nil {
		return errors.WithStack(errors.Wrap(err, "create user error"))
	}

	return c.Redirect(http.StatusSeeOther, "usersPath()")
}

func Delete(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "loading id error"))
	}

	tx := c.Value("tx").(*pop.Connection)

	user, err := findUser(c, id)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, "finding user error"))
	}

	if err := tx.Destroy(&user); err != nil {
		return errors.WithStack(errors.Wrap(err, "destroy user error"))
	}
	return c.Redirect(http.StatusFound, "usersPath()")
}
