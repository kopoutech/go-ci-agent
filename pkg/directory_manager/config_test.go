package directorymanager

import (
	"reflect"
	"testing"
)

func TestGetGlobalConfigInstance(t *testing.T) {
	tests := []struct {
		name string
		want GlobalConfigStruct
	}{
		{
			name: "TestGlobalConfig",
			want: initConfig(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGlobalConfigInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGlobalConfigInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
