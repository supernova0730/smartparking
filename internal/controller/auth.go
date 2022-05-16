package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
	"smartparking/pkg/validate"
)

type AuthController struct {
	m manager.Manager
}

func NewAuthController(m manager.Manager) *AuthController {
	return &AuthController{m: m}
}

// Register godoc
// @Description client registration route
// @Tags auth
// @Accept json
// @Produce json
// @Param registerDTO body dtos.RegisterDTO true "RegisterDTO"
// @Success 200 {object} views.ClientView
// @Failure 400 {object} validate.ValidationError
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/register [post]
func (ctl *AuthController) Register(c *fiber.Ctx) error {
	var (
		registerDTO dtos.RegisterDTO
		model       models.Client
		result      views.ClientView
	)

	err := c.BodyParser(&registerDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	if err = validate.Struct(registerDTO); err != nil {
		return err
	}

	model.SetFromRegisterDTO(registerDTO)
	client, err := ctl.m.Service().Auth().SignUp(model)
	if err != nil {
		return err
	}

	result = client.ToView()
	return response.Success(c, result)
}

// Login godoc
// @Description client login route
// @Tags auth
// @Accept json
// @Produce json
// @Param loginDTO body dtos.LoginDTO true "LoginDTO"
// @Success 200 {object} views.TokensView
// @Failure 400 {object} validate.ValidationError
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/login [post]
func (ctl *AuthController) Login(c *fiber.Ctx) error {
	var (
		loginDTO dtos.LoginDTO
		model    models.Client
	)

	err := c.BodyParser(&loginDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	if err = validate.Struct(loginDTO); err != nil {
		return err
	}

	model.SetFromLoginDTO(loginDTO)
	client, token, err := ctl.m.Service().Auth().SignIn(model)
	if err != nil {
		return err
	}

	return response.Success(c, fiber.Map{
		"tokens": token.ToView(),
		"client": client.ToView(),
	})
}

// ForgetPassword godoc
// @Description forget password step 1
// @Tags auth
// @Accept json
// @Produce json
// @Param forgetPasswordDTO body dtos.ForgetPasswordDTO true "ForgetPasswordDTO"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} validate.ValidationError
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/forget-password [post]
func (ctl *AuthController) ForgetPassword(c *fiber.Ctx) error {
	var (
		forgetPasswordDTO dtos.ForgetPasswordDTO
	)

	err := c.BodyParser(&forgetPasswordDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = validate.Struct(forgetPasswordDTO)
	if err != nil {
		return err
	}

	err = ctl.m.Service().Auth().GenerateEmailVerificationAndSendToClient(forgetPasswordDTO.Email, forgetPasswordDTO.NewPassword)
	if err != nil {
		return err
	}

	return response.Success(c, response.CommonResponse{
		Message: "please check your email",
	})
}

// RefreshTokens godoc
// @Description refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshTokenDTO body dtos.RefreshTokenDTO true "RefreshTokenDTO"
// @Success 200 {object} views.TokensView
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/refresh [post]
func (ctl *AuthController) RefreshTokens(c *fiber.Ctx) error {
	var input dtos.RefreshTokenDTO
	err := c.BodyParser(&input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	tokens, err := ctl.m.Service().Auth().RefreshTokens(input.RefreshToken)
	if err != nil {
		return err
	}

	return response.Success(c, tokens.ToView())
}

// CheckEmailVerification godoc
// @Description forget password step 2 (verifies code from email)
// @Tags auth
// @Accept json
// @Produce json
// @Param checkDTO body dtos.CheckEmailVerificationDTO true "CheckEmailVerificationDTO"
// @Success 200 {object} views.TokensView
// @Failure 400 {object} validate.ValidationError
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/verify [post]
func (ctl *AuthController) CheckEmailVerification(c *fiber.Ctx) error {
	var (
		checkDTO dtos.CheckEmailVerificationDTO
	)

	err := c.BodyParser(&checkDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = validate.Struct(checkDTO)
	if err != nil {
		return err
	}

	tokens, err := ctl.m.Service().Auth().CheckEmailVerificationAndUpdatePassword(checkDTO.Email, checkDTO.Code)
	if err != nil {
		return err
	}

	return response.Success(c, tokens.ToView())
}
