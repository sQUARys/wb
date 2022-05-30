package main

import "testing"

func Test_startBinaryFile(t *testing.T) {
	type args struct {
		binPath string
		data    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test first", args{binPath: "cdBin/cd", data: ".."}, "Succesful changing. Now you in /Users/roman/Desktop/Work/WorkRepo/L2/develop."},
		{"Test second", args{binPath: "pwdBin/pwd", data: ""}, "Now you in /Users/roman/Desktop/Work/WorkRepo/L2/develop/dev08."},
		{"Test third", args{binPath: "echoBin/echo", data: "Hello"}, "Echo: Hello\n"},
		{"Test fourth", args{binPath: "killBin/kill", data: ""}, "Process finished successfully(Killed)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startBinaryFile(tt.args.binPath, tt.args.data); got != tt.want {
				t.Errorf("startBinaryFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
