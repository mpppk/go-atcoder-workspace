package lib

import (
	"reflect"
	"testing"
)

func TestReverseType(t *testing.T) {
	type args struct {
		values []ZZZ
	}
	tests := []struct {
		name string
		args args
		want []ZZZ
	}{
		{
			name: "can reverse int slice",
			args: args{
				values: []ZZZ{1, 2, 3},
			},
			want: []ZZZ{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseZZZ(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
