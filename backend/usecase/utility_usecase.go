package usecase

func (uc *usecase) HealthCheck() error {
	return uc.repo.Ping()
}
