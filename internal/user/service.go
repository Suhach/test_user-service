package user

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, email, password string) (*User, error) {
	user := &User{Email: email, Pass: password}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Get(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) Update(id uint, email, password string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	user.Email = email
	user.Pass = password
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
