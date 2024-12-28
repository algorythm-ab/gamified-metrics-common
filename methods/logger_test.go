package methods

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestCreateLogger(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  *zap.Logger
	}{
		{
			name:  "test case 1",
			level: "debug",
			want:  zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel)),
		},
		{
			name:  "test case 2",
			level: "warn",
			want:  zaptest.NewLogger(t, zaptest.Level(zap.WarnLevel)),
		},
		{
			name:  "test case 3",
			level: "error",
			want:  zaptest.NewLogger(t, zaptest.Level(zap.ErrorLevel)),
		},
		{
			name:  "test case 4",
			level: "panic",
			want:  zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel)),
		},
		{
			name:  "test case 5",
			level: "fatal",
			want:  zaptest.NewLogger(t, zaptest.Level(zap.FatalLevel)),
		},
		{
			name:  "test case 6",
			level: "",
			want:  zaptest.NewLogger(t, zaptest.Level(zap.InfoLevel)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateLogger(tt.level)
			if got.Level() != tt.want.Level() {
				t.Errorf("Zap logger = %v, want %v", got, tt.want)
			}
		})
	}
}
