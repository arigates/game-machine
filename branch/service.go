package branch

type Service interface {
	CreateBranch(input CreateBranchInput) (Branch, error)
	GetBranchByID(input GetBranchDetailInput) (Branch, error)
	UpdateBranch(inputUri GetBranchDetailInput, inputUpdate CreateBranchInput) (Branch, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateBranch(input CreateBranchInput) (Branch, error) {
	branch := Branch{}
	branch.Code = input.Code
	branch.Name = input.Name
	branch.Address = input.Address

	newBranch, err := s.repository.Create(branch)
	if err != nil {
		return newBranch, err
	}

	return newBranch, nil
}

func (s *service) GetBranchByID(input GetBranchDetailInput) (Branch, error) {
	branch, err := s.repository.FindByID(input.ID)

	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (s *service) UpdateBranch(inputUri GetBranchDetailInput, inputUpdate CreateBranchInput) (Branch, error) {
	branch, err := s.repository.FindByID(inputUri.ID)
	if err != nil {
		return branch, err
	}

	branch.Code = inputUpdate.Code
	branch.Name = inputUpdate.Name
	branch.Address = inputUpdate.Address

	updatedBranch, err := s.repository.Update(branch)
	if err != nil {
		return updatedBranch, err
	}

	return updatedBranch, nil
}
