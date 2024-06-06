package service

import (
	events "github.com/banja001/SOA/saga/give_xp"

	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type GiveXPOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCreateOrderOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*GiveXPOrchestrator, error) {
	o := &GiveXPOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *GiveXPOrchestrator) Start(keypointId int, sessionId string) error {
	details := &events.SessionAndKeypoint{
		KeypointId: keypointId,
		SessionId:  sessionId,
	}
	event := &events.GiveXpCommand{
		Type: events.CompleteKeypoint,
		Details: events.AllDetails{
			//UserXp: *userXp,
			//SAndK: *details,
			SessionId:  details.SessionId,
			KeypointId: details.KeypointId,
		},
	}
	return o.commandPublisher.Publish(event)
}

func (o *GiveXPOrchestrator) handle(reply *events.GiveXpReply) {
	command := events.GiveXpCommand{Details: reply.Details}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *GiveXPOrchestrator) nextCommandType(reply events.GiveXPReplyType) events.GiveXPCommandType {
	switch reply {
	case events.KeypointCompleted:
		return events.AddXP
	case events.KeypointNotCompleted:
		return events.RollbackKeypoint
	default:
		return events.UnknownCommand
	}
}
