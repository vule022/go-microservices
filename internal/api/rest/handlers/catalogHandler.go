package handlers

import (
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/repository"
	"go-ecommerce/internal/service"
	"log"

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
	ver.Get("/products")
	ver.Get("/products/:id")
	ver.Get("/categories")
	ver.Get("/categories/:id")

	// Private endpoints
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)

	selRoutes.Get("/categories", handler.GetCategories)
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

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	log.Printf("current user %v", user.ID)

	return rest.SuccessMessage(ctx, "category created", nil)
}

func (h CatalogHandler) UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "category updated", fiber.Map{"id": id})
}

func (h CatalogHandler) ReplaceCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "category replaced", fiber.Map{"id": id})
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return rest.SuccessMessage(ctx, "category deleted", fiber.Map{"id": id})
}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	return rest.SuccessMessage(ctx, "categories fetched", fiber.Map{
		"items": []string{},
	})
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
