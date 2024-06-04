package give_xp

type UserExperience struct {
	ID     int
	UserID int
	XP     int
	Level  int
}
type SessionAndKeypoint struct {
	SessionId  string
	KeypointId int
}

type AllDetails struct {
	//UserXp UserExperience
	SessionId  string
	KeypointId int
}

type GiveXPCommandType int8

const (
	CompleteKeypoint GiveXPCommandType = iota
	RollbackKeypoint
	AddXP
	SubtractXP
	UnknownCommand
)

type GiveXpCommand struct {
	Details AllDetails
	Type    GiveXPCommandType
}

type GiveXPReplyType int8

const (
	KeypointCompleted GiveXPReplyType = iota
	KeypointNotCompleted
	KeypointRolledBack
	XPAdded
	XpSubtracted
	UnknownReply
)

type GiveXpReply struct {
	Details AllDetails
	Type    GiveXPReplyType
}
