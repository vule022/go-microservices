package handlers

import (
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/dto"
	"go-ecommerce/internal/repository"
	"go-ecommerce/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := UserHandler{
		svc: svc,
	}

	//Add versioning
	ver := app.Group("/v1")

	//Public routes
	pubRoutes := ver.Group("/users")

	//Public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	//Private routes
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	//Private endpoints
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)

	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserRegister{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "provide valid input",
		})
	}

	token, err := h.svc.Register(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "provide valid input",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "error on login",
			"reason":  err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	//create and db update verification code
	err := h.svc.GetVerificationCode(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "invalid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create user",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"user":    user,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart items",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create an order",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order by id",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}
	err := ctx.BodyParser(&req)

	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "req params not valid",
		})
	}

	token, err := h.svc.BecomeSeller(user.ID, req)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "failed to become seller",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "became seller",
		"token":   token,
	})
}
