package serializer

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoBufToJSON(message proto.Message) (string, error) {
	jsonBytes, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
		Indent:          "",
		UseProtoNames:   true,
	}.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	return string(jsonBytes), nil
}
