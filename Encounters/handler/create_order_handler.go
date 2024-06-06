package handler

import (
	"encgo/service"

	events "github.com/banja001/SOA/saga/give_xp"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type AddxpCommandHandler struct {
	userXPService     *service.UserExperienceService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewAddxpCommandHandler(userXPService *service.UserExperienceService, publisher saga.Publisher, subscriber saga.Subscriber) (*AddxpCommandHandler, error) {
	o := &AddxpCommandHandler{
		userXPService:     userXPService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *AddxpCommandHandler) handle(command *events.GiveXpCommand) {
	/*id, err := primitive.ObjectIDFromHex(command.Details.SessionId)
	if err != nil {
		return
	}
	num, err := strconv.Atoi(id.String())
	session := model.Session{ID: num}*/

	reply := events.GiveXpReply{Details: command.Details}

	switch command.Type {
	case events.AddXP:
		handler.userXPService.AddXP(1, 20)
		reply.Type = events.XPAdded
	case events.SubtractXP:
		reply.Type = events.XpSubtracted
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
