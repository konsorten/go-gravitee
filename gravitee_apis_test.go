package gravitee

import (
	"fmt"
	"testing"
)

func TestGetAllAPIs(t *testing.T) {
	session := createTestSession(t)

	apis, err := session.GetAllAPIs()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", apis)
}
