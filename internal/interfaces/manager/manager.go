package manager

type Manager interface {
	Repository() Repository
	Cache() Cache
	Service() Service
	Processor() Processor
	Controller() Controller
}
