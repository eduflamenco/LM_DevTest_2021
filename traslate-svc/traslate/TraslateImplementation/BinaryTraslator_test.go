package TraslateImplementation

import (
	"reflect"
	"testing"
)

func TestNewBinaryTraslator(t *testing.T) {
	type args struct {

	}
	tests := []struct {
		name string
		args args
		want *translator
	}{
		{
			name: "TestNewBinaryTraslator",
			args: args{
			},
			want: NewTraslator(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTraslator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryTraslator() = %v, want %v", got, tt.want)
			}
		})
	}
}





