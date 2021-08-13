package logrusen

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *standardLogger
	}{
		{
			name: "create standard logger",
			want: &standardLogger{logrus.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			if reflect.TypeOf(got).String() != reflect.TypeOf(tt.want).String() {
				t.Errorf("New() = %T, want %T", got, tt.want)
			}
		})
	}
}

func Test_standardLogger_Setup(t *testing.T) {
	type args struct {
		env string
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    *standardLogger
		wantErr bool
	}{
		{
			name: "invalid env name (env:development)",
			args: args{
				env: "development",
				dsn: "",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: true,
		},
		{
			name: "valid env name (env:dev)",
			args: args{
				env: "dev",
				dsn: "",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: false,
		},
		{
			name: "valid env name but dsn is invalid (env:prod, dsn:123456789)",
			args: args{
				env: "prod",
				dsn: "123456789",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: true,
		},
		{
			name: "invalid env name (env:test)",
			args: args{
				env: "test",
				dsn: "123456789",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			_, err := l.Setup(tt.args.env, tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("standardLogger.Setup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_stdFields(t *testing.T) {
	type args struct {
		event string
		topic string
	}
	tests := []struct {
		name string
		args args
		want *log.Fields
	}{
		{
			name: "foo and topic1",
			args: args{
				event: "foo",
				topic: "topic1",
			},
			want: &log.Fields{
				"event": "foo",
				"topic": "topic1",
			},
		},
		{
			name: "fib and topic2",
			args: args{
				event: "fib",
				topic: "topic2",
			},
			want: &log.Fields{
				"event": "fib",
				"topic": "topic2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stdFields(tt.args.event, tt.args.topic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stdFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
