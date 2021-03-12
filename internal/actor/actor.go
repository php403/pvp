package actor

type IActor interface {
	OnStart()
	OnMessage()
	OnBroadcast()
	OnExit()
}
