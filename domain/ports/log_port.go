package ports

type LogPortParams struct {
	Module string
	Msg    string
}

type LogPortErrorParams struct {
	Module string
	Error  error
	Msg    string
}

type LogPort interface {
	Info(params LogPortParams)
	Error(params LogPortErrorParams)
}
