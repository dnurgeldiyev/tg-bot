package telegram

import (
	"context"
	"dovran/tg-bot/internal/client/telegram"
	"dovran/tg-bot/internal/controller/events"
	"dovran/tg-bot/pkg/e"
	"errors"
)

type Meta struct {
	ChatId   int
	Username string
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

type EventHandler struct {
	tg     *telegram.Client
	offset int
}

func NewEventHandler(tg *telegram.Client) *EventHandler {
	return &EventHandler{
		tg:     tg,
		offset: 0,
	}
}

func (ev *EventHandler) Fetch(ctx context.Context, limit int) ([]events.Event, error) {
	updates, err := ev.tg.Updates(ctx, ev.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get updates", err)
	}

	if len(updates) > 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	ev.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (ev *EventHandler) Process(ctx context.Context, event events.Event) error {
	switch event.Type {
	case events.Message:
		return ev.processMessage(ctx, event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}
func (ev *EventHandler) processMessage(ctx context.Context, event events.Event) error {
	_, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	//if err := p.doCmd(ctx, event.Text, meta.ChatId, meta.Username); err != nil {
	//	return e.Wrap("can't process message", err)
	//}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}

	return res, nil
}
func event(u telegram.Update) events.Event {
	uType := getTypeUpdate(u)

	res := events.Event{
		Type: uType,
		Text: getTextUpdate(u),
	}

	if uType == events.Message {
		res.Meta = Meta{
			ChatId:   u.Message.Chat.ID,
			Username: u.Message.From.Username,
		}
	}

	return res
}
func getTypeUpdate(u telegram.Update) string {
	if u.Message != nil {
		return events.Message
	}
	return events.Unknown
}
func getTextUpdate(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}
	return u.Message.Text
}
