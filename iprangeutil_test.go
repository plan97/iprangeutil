package iprangeutil

import (
	"testing"
)

func TestExpandIPv4Range(t *testing.T) {
	printIPv4 := func(ip0, ip1, ip2, ip3 *uint8) (err error) {
		t.Logf("%d.%d.%d.%d\n", *ip0, *ip1, *ip2, *ip3)
		return
	}
	type args struct {
		start  string
		end    string
		ipFunc IPv4Func
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check for IP range 0.0.0.0 to 0.0.1.255",
			args: args{
				start:  "0.0.0.0",
				end:    "0.0.1.255",
				ipFunc: printIPv4,
			},
			wantErr: false,
		},
		{
			name: "check for IP range 0.0.255.255 to 0.1.0.0",
			args: args{
				start:  "0.0.255.255",
				end:    "0.1.0.0",
				ipFunc: printIPv4,
			},
			wantErr: false,
		},
		{
			name: "check for IP range 0.255.255.255 to 1.0.0.255",
			args: args{
				start:  "0.255.255.255",
				end:    "1.0.0.255",
				ipFunc: printIPv4,
			},
			wantErr: false,
		},
		{
			name: "check for errors in starting IP",
			args: args{
				start:  "0.0..0",
				end:    "0.0.1.0",
				ipFunc: printIPv4,
			},
			wantErr: true,
		},
		{
			name: "check for errors in ending IP",
			args: args{
				start:  "0.0.0.0",
				end:    "0.0.1111.0",
				ipFunc: printIPv4,
			},
			wantErr: true,
		},
		{
			name: "check for errors in ipFunc",
			args: args{
				start: "0.0.0.0",
				end:   "0.0.1.0",
				ipFunc: func(_, _, _, _ *uint8) error {
					return ErrIPFunc
				},
			},
			wantErr: true,
		},
		{
			name: "check for nil ipFunc",
			args: args{
				start:  "0.0.0.0",
				end:    "0.0.1.0",
				ipFunc: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandIPv4(tt.args.start, tt.args.end, tt.args.ipFunc); (err != nil) != tt.wantErr {
				t.Errorf("ExpandIPv4() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
