package fptai_go

import "testing"

func TestClient(t *testing.T) {
	want := "token"
	if got := NewClient("token").BotToken; got != want {
		t.Errorf("BotToken = %q, want %q", got, want)
	}
}
