package version

import (
	"testing"
)

func TestToolBox_Info(t *testing.T) {
	type fields struct {
		Name    string
		Version string
	}
	tests := []struct {
		name           string
		fields         fields
		wantPkgName    string
		wantPkgVersion string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				Name:    "ToolBox",
				Version: "v1.0.0",
			},
			wantPkgName:    "ToolBox",
			wantPkgVersion: "v1.0.0",
		},
		{
			name: "test02",
			fields: fields{
				Name:    "ToolBox",
				Version: "v1.0.1",
			},
			wantPkgName:    "ToolBox",
			wantPkgVersion: "v1.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := ToolBox{
				Name:    tt.fields.Name,
				Version: tt.fields.Version,
			}
			gotPkgName, gotPkgVersion := tb.Info()
			if gotPkgName != tt.wantPkgName {
				t.Errorf("ToolBox.Info() gotPkgName = %v, want %v", gotPkgName, tt.wantPkgName)
			}
			if gotPkgVersion != tt.wantPkgVersion {
				t.Errorf("ToolBox.Info() gotPkgVersion = %v, want %v", gotPkgVersion, tt.wantPkgVersion)
			}
		})
	}
}

func BenchmarkToolBox_Info(b *testing.B) {
	tb := ToolBox{
		Name:    "test",
		Version: "v1.0.0",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tb.Info()
	}
}
