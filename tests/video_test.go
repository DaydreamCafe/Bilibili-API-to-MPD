package tests

import (
	"encoding/json"
	"os"
	"testing"

	bilibiliapitompd "github.com/DaydreamCafe/Bilibili-API-To-MPD"
)

type Secret struct {
	SESSDATA string `json:"SESSDATA"`
}

func TestConvertToMPD(t *testing.T) {
	file, err := os.Open("./secret.json")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	var secret Secret
	err = json.NewDecoder(file).Decode(&secret)
	if err != nil {
		t.Error(err)
	}

	v := bilibiliapitompd.NewVideoWithBvid("BV1Lv4y1b7RH", 997535329, secret.SESSDATA)
	res, err := v.ConvertToMPD()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestConvertToMPDFile(t *testing.T) {
	file, err := os.Open("./secret.json")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	var secret Secret
	err = json.NewDecoder(file).Decode(&secret)
	if err != nil {
		t.Error(err)
	}

	v := bilibiliapitompd.NewVideoWithBvid("BV1Lv4y1b7RH", 997535329, secret.SESSDATA)
	err = v.ConvertToMPDFile("./test.mpd")
	if err != nil {
		t.Error(err)
	}
}