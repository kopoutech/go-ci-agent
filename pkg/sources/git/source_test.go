package git

import "testing"

func TestGitSource_Download(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       GitSource
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.Download(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitSource.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GitSource.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}
