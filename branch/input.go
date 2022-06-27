package branch

type GetBranchDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type CreateBranchInput struct {
	Code    string `json:"code" form:"code" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}
