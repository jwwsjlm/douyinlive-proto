package generated_test

import (
	"encoding/json"
	"testing"

	"github.com/jwwsjlm/douyinlive-proto/generated"
	"github.com/jwwsjlm/douyinlive-proto/generated/new_douyin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestRoomRankAudienceRanks(t *testing.T) {
	message, err := generated.GetMessageInstance("WebcastRoomRankMessage")
	if err != nil {
		t.Fatal(err)
	}
	defer generated.PutMessageInstance("WebcastRoomRankMessage", message)

	// Field 3 contains two Rank messages: (score=100, rank=1), (score=200, rank=2).
	payload := []byte{0x1a, 0x04, 0x10, 0x64, 0x18, 0x01, 0x1a, 0x05, 0x10, 0xc8, 0x01, 0x18, 0x02}
	if err := proto.Unmarshal(payload, message); err != nil {
		t.Fatal(err)
	}
	ranks := message.(*new_douyin.Webcast_Im_RoomRankMessage).GetAudienceRanks()
	if len(ranks) != 2 {
		t.Fatalf("audience ranks = %v, want two entries", ranks)
	}
	if ranks[0].GetScore() != 100 || ranks[0].GetRank() != 1 || ranks[1].GetScore() != 200 || ranks[1].GetRank() != 2 {
		t.Fatalf("audience ranks = %v, want distinct scores and ranks", ranks)
	}

	data, err := protojson.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		AudienceRanks []json.RawMessage `json:"audienceRanks"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}
	if len(result.AudienceRanks) != 2 {
		t.Fatalf("JSON = %s, want audienceRanks array with two entries", data)
	}
}
