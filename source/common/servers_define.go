package common

type ServerID uint8

const (
	RouteServerName          = "RouteServer"
	RouteServerID   ServerID = 0

	UserServerName          = "UserServer"
	UserServerID   ServerID = 1

	EntryServerName          = "EntryServer"
	EntryServerID   ServerID = 2

	GameServerName          = "GameServer"
	GameServerID   ServerID = 3

	PropServerName          = "PropServer"
	PropServerID   ServerID = 4

	TaskServerName          = "TaskServer"
	TaskServerID   ServerID = 5
)

//go:generate stringer -type=ServerID -output=servers_string.go
