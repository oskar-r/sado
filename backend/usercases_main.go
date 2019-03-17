package backend

type UseCase interface {
	UserUseCase
	UploadUseCase
}

var impl UseCase

func SetUC(uc UseCase) {
	impl = uc
}
