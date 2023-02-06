package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"pcbook/pb"
	"pcbook/sample"
	"testing"
)

func Test_WriteProtobufToBinaryFile_ReadProtobufFromBinaryFile(t *testing.T) {
	t.Parallel()

	binFile := "../tmp/laptop.bin"
	laptop := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop, binFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binFile, laptop2)

	require.NoError(t, err)
	require.True(t, proto.Equal(laptop, laptop2))
}

func Test_WriteProtobufToJSONFile(t *testing.T) {
	t.Parallel()

	jsonFile := "../tmp/laptop.json"
	laptop := sample.NewLaptop()

	err := WriteProtobufToJSONFile(laptop, jsonFile)

	require.NoError(t, err)
}
