package weatherdata

import "testing"

func TestGetSymbolByNumberExistsCheckSlackEmoji(t *testing.T) {
	symbol := GetSymbolByNumber(1)

	if symbol.SlackEmoji != ":sunny:" {
		t.Fail()
	}
}

func TestGetSymbolByNumberExistsCheckName(t *testing.T) {
	symbol := GetSymbolByNumber(1)

	if symbol.Name != "Sunny" {
		t.Fail()
	}
}

func TestGetSymbolByNumberUnknown(t *testing.T) {
	symbol := GetSymbolByNumber(-1)

	if symbol.Name != "Unknown" {
		t.Fail()
	}
}
