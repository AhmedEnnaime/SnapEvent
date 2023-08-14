package serializer

import (
	"testing"

	"github.com/AhmedEnnaime/SnapEvent/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/event.bin"
	jsonFile := "../tmp/event.json"

	event1 := utils.NewEvent()
	err := WriteProtoBufToBinaryFile(event1, binaryFile)
	require.NoError(t, err)

	event2 := utils.NewEvent()
	err = ReadProtoBufFromBinaryFile(binaryFile, event2)
	require.NoError(t, err)
	require.True(t, proto.Equal(event1, event2))

	err = WriteProtoBufToJSONFile(event1, jsonFile)
	require.NoError(t, err)

}
