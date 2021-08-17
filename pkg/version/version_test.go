package version

import "testing"

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
			name: "test",
			fields: fields{
				Name:    "ToolBox",
				Version: "v1.0.0",
			},
			wantPkgName:    "ToolBox",
			wantPkgVersion: "v1.0.0",
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
