package handler

import (
	"game-machine/branch"
	"game-machine/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type branchHandler struct {
	service branch.Service
}

func NewBranchHandler(branchService branch.Service) *branchHandler {
	return &branchHandler{branchService}
}

func (h *branchHandler) CreateBranch(c *gin.Context) {
	var input branch.CreateBranchInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unprocessable Entity", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBranch, err := h.service.CreateBranch(input)

	if err != nil {
		response := helper.APIResponse("Failed to add new branch", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := branch.FormatBranch(newBranch)

	response := helper.APIResponse("Successfully add new branch", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *branchHandler) GetBranch(c *gin.Context) {
	var input branch.GetBranchDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get data branch id required", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	branchDetail, err := h.service.GetBranchByID(input)

	if err != nil {
		response := helper.APIResponse("Failed to get data branch, cannot find data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Branch detail", http.StatusOK, "success", branch.FormatBranch(branchDetail))
	c.JSON(http.StatusOK, response)
}

func (h *branchHandler) UpdateBranch(c *gin.Context) {
	var inputUri branch.GetBranchDetailInput

	err := c.ShouldBindUri(&inputUri)
	if err != nil {
		response := helper.APIResponse("Failed to update branch, id required", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var updateData branch.CreateBranchInput
	err = c.ShouldBindJSON(&updateData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update branch", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedBranch, err := h.service.UpdateBranch(inputUri, updateData)
	if err != nil {
		response := helper.APIResponse("Failed to update branch", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update branch successfully", http.StatusOK, "success", branch.FormatBranch(updatedBranch))
	c.JSON(http.StatusOK, response)
}
