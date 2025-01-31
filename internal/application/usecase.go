package application

import "context"

type UseCase[Input, Output any] interface {
	Execute(ctx context.Context, input *Input) (*Output, error)
}

type OuterUseCase[Input any] interface {
	Execute(ctx context.Context, input *Input) error
}

type InnerUseCase[Output any] interface {
	Execute(ctx context.Context) (*Output, error)
}
