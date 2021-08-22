package tbRedis

import "testing"

func TestStandaloneRedis_Add(t *testing.T) {
	type fields struct {
		Address []string
	}
	type args struct {
		opt string
		k   string
		v   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// 测试： 插入数据是否成功
		{
			name: "test",
			fields: fields{
				Address: []string{
					"localhost:6379",
				},
			},
			args: args{
				"set",
				"username",
				"this is test of autotest",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := StandaloneRedis{
				Address: tt.fields.Address,
			}
			if err := sr.Add(tt.args.opt, tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("StandaloneRedis.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
