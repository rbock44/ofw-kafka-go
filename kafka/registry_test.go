package kafka

import (
	"bytes"
	"io"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReadHeader(t *testing.T) {
	buf := bytes.NewBuffer([]byte{0x0, 0x0, 0x0, 0x0, 0x1})
	schemaID, err := readSchemaID(buf)
	assert.Nil(t, err)
	assert.Equal(t, uint32(1), schemaID)
}

func BenchmarkReadSchemaID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer([]byte{0, 0, 0, 0, 1})
		_, err := readSchemaID(buf)
		if err != nil {
			b.Errorf("validate header should not fail [%#v]", err)
			return
		}
	}
}

var testSchemaID = 3
var testSchemaName = "testSchema"
var testSchemaFile = "../avsc/events-key.avsc"
var testSchemaVersion = 1

func setupDecoder(ctrl *gomock.Controller) Decoder {
	decoder := NewMockDecoder(ctrl)
	decoder.EXPECT().
		Decode(gomock.Any()).
		Return(nil, nil).
		AnyTimes()
	return decoder
}

func setupEncoder(ctrl *gomock.Controller, times int) Encoder {
	encoder := NewMockEncoder(ctrl)
	encoder.EXPECT().
		Encode(gomock.Any(), gomock.Any()).
		Return().
		Times(times)
	return encoder
}

func setupRegistryMock(t *testing.T, ctrl *gomock.Controller, mockErr error, encoder Encoder, decoder Decoder) Registry {
	ms := NewMockMessageSchema(ctrl)

	ms.EXPECT().
		GetID().
		Return(testSchemaID).
		AnyTimes()
	ms.EXPECT().
		GetDecoder().
		Return(decoder).
		AnyTimes()
	ms.EXPECT().
		GetEncoder().
		Return(encoder).
		AnyTimes()

	r := NewMockRegistry(ctrl)
	r.EXPECT().
		GetSchemaByID(gomock.Eq(testSchemaID)).
		Return(ms, nil).
		AnyTimes()

	return r
}

func Test_readSchemaID(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name:    "valid schema id",
			args:    args{reader: bytes.NewBuffer([]byte{0x0, 0x0, 0x0, 0x0, 0x5, 0x10})},
			want:    5,
			wantErr: false,
		},
		{
			name:    "invalid magic number",
			args:    args{reader: bytes.NewBuffer([]byte{0x5, 0x0, 0x0, 0x0, 0x5, 0x10})},
			want:    0,
			wantErr: true,
		},
		{
			name:    "header too short",
			args:    args{reader: bytes.NewBuffer([]byte{0x5, 0x0})},
			want:    0,
			wantErr: true,
		},
		{
			name:    "buffer empty",
			args:    args{reader: bytes.NewBuffer([]byte{})},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSchemaID(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSchemaID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadSchemaID() = %v, want %v", got, tt.want)
			}
		})
	}
}
