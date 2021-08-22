package tbRedis

import "testing"

func TestStandaloneRedis_testRedis(t *testing.T) {
	type fields struct {
		Address []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "测试v8",
			fields: fields{
				Address: []string{"localhost:6379"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := StandaloneRedis{
				Address: tt.fields.Address,
			}
			if got := sr.testRedis(); got != tt.want {
				t.Errorf("StandaloneRedis.testRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
