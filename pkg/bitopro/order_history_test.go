package bitopro

import (
	"encoding/json"
	"testing"
)

func TestAuthAPI_GetOrderHistory(t *testing.T) {
	if json, err := json.MarshalIndent(getAuthClient().GetOrderHistory(), "", "  "); err != nil {
		t.Error(err)
	} else {
		t.Logf("\n%s", string(json))
	}
}
