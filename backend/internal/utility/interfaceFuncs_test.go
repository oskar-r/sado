package utility

import (
	"reflect"
	"testing"
)

func TestInterfaceToIntSlice(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		{
			"TEST_1",
			args{
				v: []int{1, 2, 3},
			},
			[]int{1, 2, 3},
			true,
		},
		{
			"TEST_2",
			args{
				v: []interface{}{"string", 2, 3},
			},
			[]int{},
			false,
		},
		{
			"TEST_3",
			args{
				v: []int{},
			},
			[]int{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := InterfaceToIntSlice(tt.args.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceToIntSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("InterfaceToIntSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
