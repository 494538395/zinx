package net

import "testing"

func TestServer_Start(t *testing.T) {
	type fields struct {
		ServerName string
		IPVersion  string
		IP         string
		Port       int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			"start",
			fields{
				"test",
				"tcp",
				"127.0.0.1",
				9999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ServerName: tt.fields.ServerName,
				IPVersion:  tt.fields.IPVersion,
				IP:         tt.fields.IP,
				Port:       tt.fields.Port,
			}
			s.Start()
		})
	}
}
