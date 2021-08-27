package tbRedis

import (
	"reflect"
	"testing"
)

// 单节点
func TestStandaloneRedis_SetAndGet(t *testing.T) {
	type fields struct {
		Address  []string
		Password string
		DB       int
	}
	type args struct {
		opt Operator
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "测试单节点set",
			fields: fields{
				Address: []string{
					"localhost:6379",
				},
				Password: "",
				DB:       0,
			},
			args: args{
				opt: Operator{
					Opt: "set",
					K:   "auth:test:result",
					V:   "this is ok",
				},
			},
			wantResult: "OK",
			wantErr:    false,
		},
		{
			name: "测试单节点get",
			fields: fields{
				Address: []string{
					"localhost:6379",
				},
				Password: "",
				DB:       0,
			},
			args: args{
				opt: Operator{
					Opt: "get",
					K:   "auth:test:result",
					V:   "",
				},
			},
			wantResult: "this is ok",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := StandaloneRedis{
				Address:  tt.fields.Address,
				Password: tt.fields.Password,
				DB:       tt.fields.DB,
			}
			gotResult, err := sr.SetAndGet(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("StandaloneRedis.SetAndGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("StandaloneRedis.SetAndGet() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// 哨兵模式
func TestSentinelRedis_SetAndGet(t *testing.T) {
	type fields struct {
		MasterName string
		Address    []string
		UserName   string
		PassWord   string
		DB         int
	}
	type args struct {
		opt Operator
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "测试哨兵连接",
			fields: fields{
				MasterName: "mymaster",
				Address: []string{
					"172.16.123.137:26379",
					"172.16.123.138:26379",
					"172.16.123.139:26379",
				},
				UserName: "redis_cncp",
				PassWord: "wXGwskVXi2vCBSld",
				DB:       0,
			},
			args: args{
				opt: Operator{
					Opt: "set",
					K:   "auth:test:result",
					V:   "this is ok",
				},
			},
			wantResult: "OK",
			wantErr:    false,
		},
		{
			name: "测试哨兵连接",
			fields: fields{
				MasterName: "mymaster",
				Address: []string{
					"172.16.123.137:26379",
				},
				UserName: "redis_cncp",
				PassWord: "wXGwskVXi2vCBSld",
				DB:       0,
			},
			args: args{
				opt: Operator{
					Opt: "get",
					K:   "auth:test:result",
					V:   "",
				},
			},
			wantResult: "this is ok",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := SentinelRedis{
				MasterName: tt.fields.MasterName,
				Address:    tt.fields.Address,
				UserName:   tt.fields.UserName,
				PassWord:   tt.fields.PassWord,
				DB:         tt.fields.DB,
			}
			gotResult, err := sr.SetAndGet(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentinelRedis.SetAndGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SentinelRedis.SetAndGet() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// 集群模式
func TestClusterRedis_SetAndGet(t *testing.T) {
	type fields struct {
		Address  []string
		Password string
		Username string
	}
	type args struct {
		opt Operator
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "测试集群set",
			fields: fields{
				Address: []string{
					"172.16.123.131:50101",
					"172.16.123.132:50101",
					"172.16.123.133:50101",
					"172.16.123.134:50101",
					"172.16.123.135:50101",
					"172.16.123.136:50101",
				},
				Password: "RMKIOPZAdF9e2s7G",
				Username: "redis_cncp",
			},
			args: args{
				opt: Operator{
					Opt: "set",
					K:   "auth:test:result",
					V:   "this is ok",
				},
			},
			wantResult: "OK",
			wantErr:    false,
		},
		{
			name: "测试集群get",
			fields: fields{
				Address: []string{
					"172.16.123.131:50101",
				},
				Password: "RMKIOPZAdF9e2s7G",
				Username: "redis_cncp",
			},
			args: args{
				opt: Operator{
					Opt: "get",
					K:   "auth:test:result",
					V:   "",
				},
			},
			wantResult: "this is ok",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := ClusterRedis{
				Address:  tt.fields.Address,
				Password: tt.fields.Password,
				Username: tt.fields.Username,
			}
			gotResult, err := sr.SetAndGet(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterRedis.SetAndGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ClusterRedis.SetAndGet() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
