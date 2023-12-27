package dike

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

const (
	Text = iota
	JSON
	Pretty
)

type Dike struct {
	Logger *slog.Logger
}

type Options struct {
	Out     io.Writer
	IsDebug bool
	Format  int
}

func New(options *Options) *Dike {
	if options == nil {
		*options = Options{
			Out:     os.Stderr,
			IsDebug: false,
			Format:  Pretty,
		}
	}

	logLevel := slog.LevelInfo
	if options.IsDebug {
		logLevel = slog.LevelDebug
	}

	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}

	var handler slog.Handler

	switch options.Format {
	case Text:
		handler = slog.NewTextHandler(options.Out, &opts)
	case JSON:
		handler = slog.NewJSONHandler(options.Out, &opts)
	case Pretty:
		handler = &PrettyHandler{
			Handler: slog.NewJSONHandler(options.Out, &opts),
			l:       log.New(os.Stderr, "", 0),
		}
	default:
		handler = slog.NewTextHandler(options.Out, &opts)
	}

	sl := slog.New(handler)

	return &Dike{
		Logger: sl,
	}
}

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	//b, err := json.MarshalIndent(fields, "", "  ")
	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[2006-01-02T15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}
