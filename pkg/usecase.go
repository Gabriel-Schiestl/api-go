package pkg

type UseCase interface {
	Execute() (any, error)
}