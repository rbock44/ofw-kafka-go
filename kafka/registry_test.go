package kafka

import (
	"bytes"
	"io"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -source registry.go -destination mock_registry_test.go -package kafka SchemaResolver,Encoder,Decoder

func TestWriteHeader(t *testing.T) {
	buf := &bytes.Buffer{}
	testSchema.WriteHeader(buf)
	assert.Equal(t, []byte{0x0, 0x0, 0x0, 0x0, byte(testSchema.ID)}, buf.Bytes())
}

func BenchmarkWriteHeader(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		testSchema.WriteHeader(buf)
		buf.Reset()
	}
}

func Test_Registry_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var testSchemaFile = "test.avsc"

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewSchemaResolver().
		Return(setupTestResolver(ctrl), nil)
	SetFrameworkFactory(f)

	registry, _ := NewSchemaRegistry()

	schema, err := registry.Register(testSchema.Subject, testSchema.Version, testSchemaFile, testDecoder, testEncoder)
	if assert.Nil(t, err) {
		if assert.NotNil(t, schema) {
			assert.Equal(t, testSchema, schema)
		}
	}
}

func Test_Registry_Lookup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewSchemaResolver().
		Return(setupTestResolver(ctrl), nil)
	SetFrameworkFactory(f)

	registry, _ := NewSchemaRegistry()
	schema, err := registry.Lookup(testSchema.Subject, testSchema.Version)
	if assert.Nil(t, err) {
		if assert.NotNil(t, schema) {
			assert.Equal(t, testSchema.ID, schema.ID)
		}
	}
}

func Test_Registry_GetSchemaByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewSchemaResolver().
		Return(setupTestResolver(ctrl), nil)
	SetFrameworkFactory(f)

	registry, _ := NewSchemaRegistry()
	var testSchemaFile = "test.avsc"

	schema, err := registry.Register(testSchema.Subject, testSchema.Version, testSchemaFile, testDecoder, testEncoder)
	if assert.Nil(t, err) {
		if assert.NotNil(t, schema) {
			assert.Equal(t, testSchema, schema)
		}
	}
	schema, err = registry.GetSchemaByID(int(testSchema.ID))
	if assert.Nil(t, err) {
		if assert.NotNil(t, schema) {
			assert.Equal(t, testSchema, schema)
			assert.NotNil(t, schema.GetDecoder())
			assert.NotNil(t, schema.GetEncoder())
		}
	}
}

func Test_Registry_GetSchemaByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := NewMockProvider(ctrl)
	f.EXPECT().
		NewSchemaResolver().
		Return(setupTestResolver(ctrl), nil)
	SetFrameworkFactory(f)

	registry, _ := NewSchemaRegistry()

	var testSchemaFile = "test.avsc"

	schema, err := registry.Register(testSchema.Subject, testSchema.Version, testSchemaFile, testDecoder, testEncoder)
	if assert.Nil(t, err) {
		if assert.NotNil(t, schema) {
			assert.Equal(t, testSchema, schema)
		}
	}
	schema, err = registry.GetSchemaByName(testSchema.Subject)
	if assert.Nil(t, err) {
		if assert.NotNil(t, schema) {
			assert.Equal(t, testSchema, schema)
		}
	}
}

var testSchema = &AvroSchema{
	Subject: "testSchema",
	ID:      5,
	Version: 1,
	Decoder: testDecoder,
	Encoder: testEncoder,
}

type testDecoderType struct{}

var testDecoder = &testDecoderType{}

func (d *testDecoderType) Decode(reader io.Reader) (interface{}, error) {
	return "test", nil
}

type testEncoderType struct {
	NumCalls int
}

func (d *testEncoderType) Encode(data interface{}, writer io.Writer) {
	d.NumCalls++
}

var testEncoder = &testEncoderType{}

type testResolver struct {
	TestSubjects map[string]int
}

func (r *testResolver) GetSchemaBySubject(subject string, version int) (schemaID int, err error) {
	return r.TestSubjects[subject], nil
}

func setupTestResolver(ctrl *gomock.Controller) SchemaResolver {
	resolver := NewMockSchemaResolver(ctrl)
	resolver.EXPECT().
		RegisterNewSchema(gomock.Eq(testSchema.Subject), gomock.Any()).
		Return(int(testSchema.ID), nil).
		AnyTimes()
	resolver.EXPECT().
		GetSchemaBySubject(gomock.Eq(testSchema.Subject), gomock.Eq(testSchema.Version)).
		Return(int(testSchema.ID), nil).
		AnyTimes()
	return resolver
}

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

func setupRegistryMock(t *testing.T, ctrl *gomock.Controller, mockErr error, encoder Encoder, decoder Decoder) *MockRegistry {
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
