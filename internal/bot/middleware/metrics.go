package middleware

import (
	"app/internal/bot/metrics" // путь подставь свой
	"github.com/prometheus/client_golang/prometheus"
	tele "gopkg.in/telebot.v4"
)

func InstrumentHandler(name string, h func(c tele.Context) error) func(c tele.Context) error {
	return func(c tele.Context) error {
		updateType := detectUpdateType(c)
		chatType := detectChatType(c)

		metrics.InFlight.Inc()
		timer := prometheusNewTimer(metrics.HandlerDuration.WithLabelValues(name))
		defer func() {
			timer.ObserveDuration()
			metrics.InFlight.Dec()
		}()

		err := h(c)

		metrics.MessagesTotal.WithLabelValues(name, updateType, chatType).Inc()
		if err != nil {
			metrics.HandlerErrors.WithLabelValues(name).Inc()
		}
		return err
	}
}

func prometheusNewTimer(obs prometheus.Observer) *prometheus.Timer {
	return prometheus.NewTimer(obs)
}

func detectUpdateType(c tele.Context) string {
	if c == nil {
		return "unknown"
	}
	if c.Callback() != nil {
		return "callback_query"
	}
	if m := c.Message(); m != nil {
		// можно детализировать: photo, document, sticker ...
		if m.Photo != nil {
			return "photo"
		}
		if m.Text != "" {
			return "message"
		}
		return "message"
	}
	return "unknown"
}

func detectChatType(c tele.Context) string {
	if c == nil || c.Chat() == nil {
		return "unknown"
	}
	return string(c.Chat().Type)
}
