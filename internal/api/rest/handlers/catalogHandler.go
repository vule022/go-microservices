package handlers

import (
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/dto"
	"go-ecommerce/internal/repository"
	"go-ecommerce/internal/service"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := CatalogHandler{
		svc: svc,
	}

	//Add versioning
	ver := app.Group("/v1")

	//Public endpoints
	ver.Get("/products", handler.ListProducts)
	ver.Get("/products/:id", handler.GetProduct)
	ver.Get("/categories", handler.ListCategories)
	ver.Get("/categories/:id", handler.GetCategoryById)

	// Private endpoints
	selRoutes := ver.Group("/seller", rh.Auth.AuthorizeSeller)

	selRoutes.Get("/categories", handler.ListCategories)
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.UpdateCategory)
	selRoutes.Put("/categories/:id", handler.ReplaceCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	selRoutes.Get("/products", handler.ListProducts)
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Patch("/products/:id", handler.UpdateProduct)
	selRoutes.Put("/products/:id", handler.ReplaceProduct)
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (h CatalogHandler) ListCategories(ctx *fiber.Ctx) error {
	cats, err := h.svc.ListCategories()

	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessMessage(ctx, "all categories", cats)
}

func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	cat, err := h.svc.GetCategory(id)

	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessMessage(ctx, "category", cat)
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)

	if err != nil {
		log.Print(err)
		return rest.BadRequest(ctx, "bad req")
	}

	err = h.svc.CreateCategory(req)

	if err != nil {
		return rest.InternalError(ctx, 500, err)
	}

	return rest.SuccessMessage(ctx, "category created", nil)
}

func (h CatalogHandler) UpdateCategory(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := ctx.BodyParser(&req)

	if err != nil {
		log.Print(err)
		return rest.BadRequest(ctx, "bad req")
	}

	cat, err := h.svc.EditCategory(id, req)

	if err != nil {
		return rest.InternalError(ctx, 500, err)
	}

	return rest.SuccessMessage(ctx, "category created", cat)
}

func (h CatalogHandler) ReplaceCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "category replaced", fiber.Map{"id": id})
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := h.svc.DeleteCategory(id)

	if err != nil {
		return rest.InternalError(ctx, 500, err)
	}

	return rest.SuccessMessage(ctx, "category deleted", fiber.Map{"id": id})
}

func (h CatalogHandler) ListProducts(ctx *fiber.Ctx) error {
	return rest.SuccessMessage(ctx, "products fetched", fiber.Map{
		"items": []string{},
	})
}

func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return rest.SuccessMessage(ctx, "product created", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "product fetched", fiber.Map{"id": id})
}

func (h CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "product updated", fiber.Map{"id": id})
}

func (h CatalogHandler) ReplaceProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "product replaced", fiber.Map{"id": id})
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "product deleted", fiber.Map{"id": id})
}
