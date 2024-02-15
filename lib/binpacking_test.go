package binpacking

import (
	"reflect"
	"repack/utils"
	"testing"
)

func TestPackOrder(t *testing.T) {
	// init logger
	utils.InitializeLogger()

	// setup tests
	tests := []struct {
		name        string
		packSizes   []int
		orderSize   int
		want        []int
		expectError bool
	}{
		{
			name:        "Valid input with exact match",
			packSizes:   []int{250, 500, 1000, 2000},
			orderSize:   501,
			want:        []int{1, 1, 0, 0},
			expectError: false,
		},
		{
			name:        "Valid input with no exact match",
			packSizes:   []int{3, 4, 5},
			orderSize:   10,
			want:        []int{0, 0, 2},
			expectError: false,
		},
		{
			name:        "Order size smaller than smallest pack",
			packSizes:   []int{5, 10, 20},
			orderSize:   2,
			want:        nil,
			expectError: true,
		},
		{
			name:        "Empty pack sizes",
			packSizes:   []int{},
			orderSize:   10,
			want:        nil,
			expectError: true,
		},
		{
			name:        "Negative order size",
			packSizes:   []int{1, 2, 3},
			orderSize:   -1,
			want:        nil,
			expectError: true,
		},
		{
			name:        "Zero order size",
			packSizes:   []int{1, 2, 3},
			orderSize:   0,
			want:        nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PackOrder(tt.packSizes, tt.orderSize)
			if (err != nil) != tt.expectError {
				t.Errorf("PackOrder() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
