package cmd

type ExecutorOptsFunc func(*Executor)

func BuildExecutor(optFuncs ...ExecutorOptsFunc) *Executor {
	p := defaultExecutor()
	for _, fn := range optFuncs {
		fn(p)
	}

	return p
}

func WithCommandExecutor(e CommandExecutor) ExecutorOptsFunc {
	return func(p *Executor) {
		p.CommandExecutor = e
	}
}

func defaultExecutor() *Executor {
	return &Executor{
		CommandExecutorDefault{},
	}
}
