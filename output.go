package zerologgelfoutput

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/rs/zerolog"
	"gopkg.in/aphistic/golf.v0"
)

type zGelfOutput struct {
	passthrough                 io.Writer
	gl                          *golf.Logger
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
	return new(addr, appName, env, version, io.Discard)
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
// // OR
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

func (w *zGelfOutput) WriteZerologMessage(p []byte) (n int, err error) {
	n, err = w.passthrough.Write(p)
	if err != nil {
		return
	}

	var jsonMsg map[string]interface{}
	if err := json.Unmarshal(p, &jsonMsg); err != nil {
		return 0, err
	}
	message := jsonMsg["message"].(string)
	delete(jsonMsg, "message")
	level := zerolog.NoLevel.String()
	if _, ok := jsonMsg["level"]; ok {
		level = jsonMsg["level"].(string)
	}
	delete(jsonMsg, "level")
	delete(jsonMsg, "time")
	delete(jsonMsg, "timestamp")

	switch level {
	case zerolog.DebugLevel.String():
		err = w.gl.Dbgm(jsonMsg, message)
	case zerolog.InfoLevel.String():
		err = w.gl.Infom(jsonMsg, message)
	case zerolog.WarnLevel.String():
		err = w.gl.Warnm(jsonMsg, message)
	case zerolog.ErrorLevel.String():
		err = w.gl.Errm(jsonMsg, message)
	case zerolog.FatalLevel.String():
		err = w.gl.Critm(jsonMsg, message)
	case zerolog.PanicLevel.String():
		err = w.gl.Emergm(jsonMsg, message)
	default:
		err = w.gl.Noticem(jsonMsg, message)
	}

	return
}

func (w *zGelfOutput) Close() error {
	if closer, ok := w.passthrough.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}

var _ io.WriteCloser = &zGelfOutput{}
