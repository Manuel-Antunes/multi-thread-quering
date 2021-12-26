// concurrent.go
package main

import (
	"testing"
)

func Test_call(t *testing.T) {
	type args struct {
		url   string
		index int
		ch    chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "call default and recieve channel", args: args{url: "http://localhost:8000/test/1", index: 1, ch: make(chan string)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go call(tt.args.url, tt.args.index, tt.args.ch)
			<-tt.args.ch
		})
	}
}
