package handler

import (
	"database-example/service"

	events "github.com/banja001/SOA/saga/give_xp"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type AddXPCommandHandler struct {
	sessionService    *service.SessionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewAddXPCommandHandler(sessionService *service.SessionService, publisher saga.Publisher, subscriber saga.Subscriber) (*AddXPCommandHandler, error) {
	o := &AddXPCommandHandler{
		sessionService:    sessionService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *AddXPCommandHandler) handle(command *events.GiveXpCommand) {
	/*id, err := primitive.ObjectIDFromHex(command.Details.SessionId)
	if err != nil {
		return
	}
	num, err := strconv.Atoi(id.String())
	session := model.Session{ID: num}*/

	reply := events.GiveXpReply{Details: command.Details}

	switch command.Type {
	case events.CompleteKeypoint:
		handler.sessionService.CompleteKeypoint(command.Details.SessionId, command.Details.KeypointId)
		reply.Type = events.KeypointCompleted
	case events.RollbackKeypoint:
		reply.Type = events.KeypointRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
