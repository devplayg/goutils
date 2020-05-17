package network

import (
	"net"
	"reflect"
	"testing"
)

func TestCidrHostCount(t *testing.T) {
	type args struct {
		cidr int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CidrHostCount(tt.args.cidr); got != tt.want {
				t.Errorf("CidrHostCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToIp(t *testing.T) {
	type args struct {
		nn uint32
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{name: "Int to IP", args: struct{ nn uint32 }{nn: 3232235522}, want: net.IP{192, 168, 0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToIp(tt.args.nn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpToInt(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "IP to Int", args: struct{ ip net.IP }{ip: net.IP{192, 168, 0, 2}}, want: 3232235522},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IpToInt(tt.args.ip); got != tt.want {
				t.Errorf("IpToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
