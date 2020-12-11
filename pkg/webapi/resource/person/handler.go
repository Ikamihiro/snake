package person

import (
	"github.com/gofiber/fiber"
	entities "snake/internal/entities/person"
	"snake/pkg/repositories/person"
	"strconv"
)

type Handler struct {
	Repository *person.Repository
}

func NewHandler(repository *person.Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

func (handler *Handler) GetAll(c *fiber.Ctx) {
	people, err := handler.Repository.GetAll()

	if err != nil {
		c.Status(500)
		_ = c.JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	c.Status(200)
	_ = c.JSON(fiber.Map{
		"ok":   true,
		"data": people,
	})
	return
}

func (handler *Handler) Find(c *fiber.Ctx){
	personId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	personSingle, err := handler.Repository.GetById(uint(personId))
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	_ = c.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": personSingle,
	})
	return
}

func (handler *Handler) Store(c *fiber.Ctx) {
	p := new(entities.Person)

	if err := c.BodyParser(p); err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	err := handler.Repository.Store(p)
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	_ = c.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": "Person saved successfully",
	})
	return
}

func (handler *Handler) Update(c *fiber.Ctx) {
	personId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	personSingle, err := handler.Repository.GetById(uint(personId))

	if err := c.BodyParser(personSingle); err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	err = handler.Repository.Update(personSingle)
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	_ = c.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": "Person updated successfully",
	})
	return
}

func (handler *Handler) Remove(c *fiber.Ctx) {
	personId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	err = handler.Repository.Delete(uint(personId))
	if err != nil {
		_ = c.Status(500).JSON(fiber.Map{
			"ok": false,
			"error": err.Error(),
		})
		return
	}

	_ = c.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": "Person removed successfully",
	})
	return
}
