package zerologgelfoutput

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_new(t *testing.T) {
	type args struct {
		addr    string
		appName string
		env     string
		version string
	}
	tests := []struct {
		name            string
		args            args
		want            io.WriteCloser
		wantPassthrough string
		wantErr         bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passthrough := &bytes.Buffer{}
			got, err := new(tt.args.addr, tt.args.appName, tt.args.env, tt.args.version, passthrough)
			if (err != nil) != tt.wantErr {
				t.Errorf("new() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("new() = %v, want %v", got, tt.want)
			}
			if gotPassthrough := passthrough.String(); gotPassthrough != tt.wantPassthrough {
				t.Errorf("new() = %v, want %v", gotPassthrough, tt.wantPassthrough)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		addr    string
		appName string
		env     string
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    io.WriteCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.addr, tt.args.appName, tt.args.env, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithPassthrough(t *testing.T) {
	type args struct {
		addr    string
		appName string
		env     string
		version string
	}
	tests := []struct {
		name            string
		args            args
		want            io.WriteCloser
		wantPassthrough string
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passthrough := &bytes.Buffer{}
			got, err := NewWithPassthrough(tt.args.addr, tt.args.appName, tt.args.env, tt.args.version, passthrough)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithPassthrough() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithPassthrough() = %v, want %v", got, tt.want)
			}
			if gotPassthrough := passthrough.String(); gotPassthrough != tt.wantPassthrough {
				t.Errorf("NewWithPassthrough() = %v, want %v", gotPassthrough, tt.wantPassthrough)
			}
		})
	}
}

func Test_zGelfOutput_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		w       *zGelfOutput
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.w.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("zGelfOutput.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("zGelfOutput.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_zGelfOutput_WriteZerologMessage(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		w       *zGelfOutput
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.w.WriteZerologMessage(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("zGelfOutput.WriteZerologMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("zGelfOutput.WriteZerologMessage() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_zGelfOutput_Close(t *testing.T) {
	tests := []struct {
		name    string
		w       *zGelfOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("zGelfOutput.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
