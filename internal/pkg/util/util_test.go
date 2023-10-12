package util

import (
	"mime/multipart"
	"os"
	"testing"
)

func TestFileType(t *testing.T) {
	type args struct {
		file *os.File
	}
	tests := []struct {
		name         string
		args         args
		wantFileType string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileType, err := FileType(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFileType != tt.wantFileType {
				t.Errorf("FileType() gotFileType = %v, want %v", gotFileType, tt.wantFileType)
			}
		})
	}
}

func TestFormFileType(t *testing.T) {
	type args struct {
		fileHeader *multipart.FileHeader
	}
	tests := []struct {
		name         string
		args         args
		wantFileType string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileType, err := FormFileType(tt.args.fileHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormFileType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFileType != tt.wantFileType {
				t.Errorf("FormFileType() gotFileType = %v, want %v", gotFileType, tt.wantFileType)
			}
		})
	}
}

func TestGenUnixTimeNano(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenUnixTimeNano(); got != tt.want {
				t.Errorf("GenUnixTimeNano() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenderTransform(t *testing.T) {
	type args struct {
		gender string
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
			if got := GenderTransform(tt.args.gender); got != tt.want {
				t.Errorf("GenderTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFileImage(t *testing.T) {
	type args struct {
		file *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFileImage(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFileImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFormFileImage(t *testing.T) {
	type args struct {
		fileHeader *multipart.FileHeader
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFormFileImage(tt.args.fileHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFormFileImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFormFileImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandCode(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{6},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := RandSmsCode(tt.args.number)
			t.Log(got)
			if got != tt.want {
				t.Errorf("RandSmsCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {

	type args struct {
		password        string
		encryptPassword string
	}

	pass1 := "123456"
	pass2 := "asdfawefae"
	pass3 := "123423d2d1"
	pass4 := "23racase542"
	encryptedPass1, _ := Crypt(pass1)
	encryptedPass2, _ := Crypt(pass2)
	encryptedPass3, _ := Crypt(pass3)
	encryptedPass4, _ := Crypt(pass4)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				password:        pass1,
				encryptPassword: encryptedPass1,
			},
			true,
		},
		{
			"2",
			args{
				password:        pass2,
				encryptPassword: encryptedPass2,
			},
			true,
		},
		{
			"3",
			args{
				password:        pass3,
				encryptPassword: encryptedPass3,
			},
			true,
		},
		{
			"4",
			args{
				password:        pass4,
				encryptPassword: encryptedPass4,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := VerifyPassword(tt.args.password, tt.args.encryptPassword); got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
