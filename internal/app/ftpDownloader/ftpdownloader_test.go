package ftpdownloader

import (
	"path/filepath"
	"testing"

	"github.com/Sterks/XmlReader/internal/app/configuration"
	"github.com/secsy/goftp"
	"github.com/sirupsen/logrus"
)

func TestFtpDownloader_GetFiles(t *testing.T) {
	type fields struct {
		config *configuration.Configuration
		logger *logrus.Logger
		ftp    *goftp.Client
	}
	type args struct {
		client *goftp.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FtpDownloader{
				config: tt.fields.config,
				logger: tt.fields.logger,
				ftp:    tt.fields.ftp,
			}
			f.GetFiles(tt.args.client)
		})
	}
}

func TestWalk(t *testing.T) {
	type args struct {
		client *goftp.Client
		root   string
		walkFn filepath.WalkFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Walk(tt.args.client, tt.args.root, tt.args.walkFn); (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
