package backend

type UseCase interface {
	UserUseCase
	UploadUseCase
	SessionManagement
	EnforcePolicy(i ...interface{}) bool
}

var impl UseCase

func SetUC(uc UseCase) {
	impl = uc
}

func EnforcePolicy(i ...interface{}) bool {
	return impl.EnforcePolicy(i)
}
