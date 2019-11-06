package tools

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func createFile() {
	_, err := os.Create("output_response.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestWriteFile(t *testing.T) {
	createFile()
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"TestWriteFile", args{"testing"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		{"TestReadFile", []string{"testing"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"TestTruncateFile", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TruncateFile(); (err != nil) != tt.wantErr {
				t.Errorf("TruncateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
