package zerologgelfoutput

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/aphistic/golf.v0"
)

type zGelfOutput struct {
	passthrough                 io.Writer
	gl                          *golf.Logger
	gc                          *golf.Client
	addr, appName, env, version string
}

func new(addr, appName, env, version string, passthrough io.Writer) (io.WriteCloser, error) {
	gl, err := golf.NewClient()
	if err != nil {
		return nil, fmt.Errorf("error creating gelf client: %w", err)
	}
	err = gl.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("error dialing gelf server: %w", err)
	}

	l, _ := gl.NewLogger()
	l.SetAttr("application", appName)
	l.SetAttr("environment", env)
	l.SetAttr("version", version)

	return &zGelfOutput{
		passthrough: passthrough,
		addr:        addr,
		appName:     appName,
		env:         env,
		version:     version,
		gl:          l,
	}, nil
}

// New creates a new io.WriteCloser that sends logs to a GELF server at the given address.
func New(addr, appName, env, version string) (io.WriteCloser, error) {
	return new(addr, appName, env, version, os.Stdout)
}

// NewWithPassthrough creates a new io.WriteCloser that sends logs to a GELF server at the given address
// and also writes the logs to the given passthrough writer.
//
// This is useful for debugging, as it allows you to see the logs that are being sent to the GELF server.
// If you don't need this, use New instead.
//
// example:
//
//	w, err := NewWithPassthrough("localhost:12201", "myapp", "dev", "1.0.0", os.Stdout)
//	if err != nil {
//		...
//	}
//
// OR
//
//	w, err := NewWithPassthrough("localhost", "myapp", "dev", "1.0.0", zerolog.ConsoleWriter{Out: os.Stdout}))
//	if err != nil {
//		...
//	}
func NewWithPassthrough(addr, appName, env, version string, passthrough io.Writer) (io.WriteCloser, error) {
	return new(addr, appName, env, version, passthrough)
}

func (w *zGelfOutput) Write(p []byte) (n int, err error) {
	return w.WriteZerologMessage(p)
}

var zerologLevelsToGelfLevels = map[string]int{
	zerolog.DebugLevel.String(): golf.LEVEL_DBG,
	zerolog.InfoLevel.String():  golf.LEVEL_INFO,
	zerolog.WarnLevel.String():  golf.LEVEL_WARN,
	zerolog.ErrorLevel.String(): golf.LEVEL_ERR,
	zerolog.FatalLevel.String(): golf.LEVEL_CRIT,
	zerolog.PanicLevel.String(): golf.LEVEL_EMERG,
}

func (w *zGelfOutput) WriteZerologMessage(p []byte) (n int, err error) {
	n, err = w.passthrough.Write(p)
	if err != nil {
		return
	}

	var jsonMsg map[string]interface{}
	if err := json.Unmarshal(p, &jsonMsg); err != nil {
		return 0, err
	}

	message := getStringFromMap(jsonMsg, "message", "")
	delete(jsonMsg, "message")

	level := getStringFromMap(jsonMsg, "level", zerolog.InfoLevel.String())
	delete(jsonMsg, "level")

	delete(jsonMsg, "time")
	delete(jsonMsg, "timestamp")

	golfMsg := w.gl.NewMessage()
	if _, ok := zerologLevelsToGelfLevels[level]; ok {
		golfMsg.Level = zerologLevelsToGelfLevels[level]
	} else {
		golfMsg.Level = golf.LEVEL_INFO
	}
	golfMsg.ShortMessage = message
	golfMsg.Attrs = jsonMsg

	// w.gl.(golfMsg)
	err = w.gc.QueueMsg(golfMsg)

	return
}

func (w *zGelfOutput) Close() error {
	if closer, ok := w.passthrough.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}

var _ io.WriteCloser = &zGelfOutput{}
