package protobuf

import (
	"log"
	"github.com/golang/protobuf/proto"
)

//Encode : Returns byte array after encoding the proto message pb.
func Encode(pb proto.Message) []byte {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Coudn't encode the message : ",err)
		return nil
	}
	return out
}

func Decode(byteArr []byte, pb proto.Message) proto.Message {
	err := proto.Unmarshal(byteArr, pb)
	if err != nil {
		log.Fatalln("Coudn't decode the message : ",err)
	}

	return pb
}