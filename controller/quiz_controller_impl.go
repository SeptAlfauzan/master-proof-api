package controller

import (
	"errors"
	"fmt"
	"master-proof-api/dto"
	"master-proof-api/service"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
)

type QuizControllerImpl struct {
	QuizService service.QuizService
}

func NewQuizController(quizService service.QuizService) QuizController {
	return &QuizControllerImpl{
		QuizService: quizService,
	}
}

func (controller *QuizControllerImpl) FindQuizWithCorrectAnswer(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	result, err := controller.QuizService.FindQuizWithCorrectAnswer(name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

func (controller *QuizControllerImpl) FindQuizWithoutCorrectAnswer(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	result, err := controller.QuizService.FindQuizWithoutCorrectAnswer(name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}
func (controller *QuizControllerImpl) CreateUserDiagnosticReport(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	quizId := ctx.Params("name")
	var result dto.RequestBodyResult
	err := ctx.BodyParser(&result)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request := dto.DiagnosticReportRequest{
		UserId:             userId,
		QuizId:             quizId,
		DiagnosticReportId: result.Result,
	}
	err = controller.QuizService.CreateUserDiagnosticReport(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})

}

func (controller *QuizControllerImpl) FindUserDiagnosticReport(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	quizId := ctx.Params("name")
	fmt.Println(quizId, userId)
	if quizId != "learning-modalities-test" && quizId != "prior-knowledge-test" && quizId != "proof-format-preference-test" {
		return ctx.Status(400).JSON(fiber.Map{
			"errors": "Invalid quizId. Accepted values are: 'learning-modalities-test', 'prior-knowledge-test', 'proof-format-preference-test'.",
		})
	}
	request := dto.RequestGetDiagnosticResult{
		UserId:   userId,
		QuizName: quizId,
	}

	result, err := controller.QuizService.FindUserDiagnosticReport(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

func (controller *QuizControllerImpl) CreateUserCompetenceReport(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	quizId := ctx.Params("name")
	var score dto.RequestBodyScore
	err := ctx.BodyParser(&score)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	request := dto.CompetenceReportRequest{
		UserId: userId,
		QuizId: quizId,
		Score:  score.Score,
	}

	err = controller.QuizService.CreateUserCompetenceReport(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

func (controller *QuizControllerImpl) FindUserCompetenceReport(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	quizId := ctx.Params("name")
	request := dto.RequestGetCompetenceResult{
		UserId:   userId,
		QuizName: quizId,
	}

	result, err := controller.QuizService.FindUserCompetenceReport(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}
func (controller *QuizControllerImpl) FindUserDiagnosticReportForTeacher(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	quizId := ctx.Params("name")

	if quizId != "learning-modalities-test" && quizId != "prior-knowledge-test" && quizId != "proof-format-preference-test" {
		return ctx.Status(400).JSON(fiber.Map{
			"errors": "Invalid quizId. Accepted values are: 'learning-modalities-test', 'prior-knowledge-test', 'proof-format-preference-test'.",
		})
	}
	request := dto.RequestGetDiagnosticResult{
		UserId:   userId,
		QuizName: quizId,
	}

	result, err := controller.QuizService.FindUserDiagnosticReport(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}
func (controller *QuizControllerImpl) FindUserCompetenceReportForTeacher(ctx *fiber.Ctx) error {

	userId := ctx.Params("userId")
	quizId := ctx.Params("name")
	request := dto.RequestGetCompetenceResult{
		UserId:   userId,
		QuizName: quizId,
	}

	result, err := controller.QuizService.FindUserCompetenceReport(request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

// AvailableQuizCategories implements QuizController.
func (controller *QuizControllerImpl) AvailableDiagnosticQuizCategories(ctx *fiber.Ctx) error {
	result, err := controller.QuizService.GetAllDiagnosticQuizzesCategories()
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

func (controller *QuizControllerImpl) AvailableCompetenceQuizCategories(ctx *fiber.Ctx) error {
	result, err := controller.QuizService.GetAllCompetenceQuizzesCategories()
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

// CalculateCompetenceQuizResult implements QuizController.
func (controller *QuizControllerImpl) CalculateCompetenceQuizResult(ctx *fiber.Ctx) error {
	subCategory := ctx.Params("id")
	var request dto.RequestCalculateQuizResult
	errRequest := ctx.BodyParser(&request)
	request.QuizSubCategory = subCategory
	if errRequest != nil {
		return fiber.NewError(fiber.StatusInternalServerError, errRequest.Error())
	}

	result, errService := controller.QuizService.CalculateCompentenceQuizResult(request)
	if errService != nil {
		var fiberErr *fiber.Error
		if errors.As(errService, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": errService.Error(),
			})
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

// CalculateDiagnosticQuizResult implements QuizController.
func (controller *QuizControllerImpl) CalculateDiagnosticQuizResult(ctx *fiber.Ctx) error {
	subCategory := ctx.Params("id")
	var request dto.RequestCalculateQuizResult
	errRequest := ctx.BodyParser(&request)
	request.QuizSubCategory = subCategory
	if errRequest != nil {
		return fiber.NewError(fiber.StatusInternalServerError, errRequest.Error())
	}
	result, errService := controller.QuizService.CalculateDiagnosticQuizResult(request)
	if errService != nil {
		var fiberErr *fiber.Error
		if errors.As(errService, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": errService.Error(),
			})
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}
