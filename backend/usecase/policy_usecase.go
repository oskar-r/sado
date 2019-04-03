package usecase

func (uc *usecase) EnforcePolicy(i ...interface{}) bool {
	return uc.enforcer.Enforce(i)
}
