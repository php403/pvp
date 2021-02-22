package actor

type IActor interface {
	OnStart()
	OnMessage()
	OnExit()
}
