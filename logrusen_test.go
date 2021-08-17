package logrusen

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
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
				dsn: "",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: true,
		},
		{
			name: "valid env name (env:dev)",
			args: args{
				dsn: "",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: false,
		},
		{
			name: "valid env name but dsn is invalid (env:prod, dsn:123456789)",
			args: args{
				dsn: "123456789",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: true,
		},
		{
			name: "invalid env name (env:test)",
			args: args{
				dsn: "123456789",
			},
			want:    &standardLogger{logrus.New()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			err := l.SetupWithSentry(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("standardLogger.Setup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
