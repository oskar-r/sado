package backend

type UseCase interface {
	UserUseCase
	UploadUseCase
	SessionManagement
	EnforcePolicy(i ...interface{}) bool
	HealthCheck() error
}

var impl UseCase

func SetUC(uc UseCase) {
	impl = uc
}

func EnforcePolicy(i ...interface{}) bool {
	return impl.EnforcePolicy(i)
}

func HealthCheck() error {
	return impl.HealthCheck()
}
