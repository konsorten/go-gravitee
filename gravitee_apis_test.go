package gravitee

import (
	"fmt"
	"testing"
)

func TestGetAPI(t *testing.T) {
	session := createTestSession(t)

	apis, err := session.GetAllAPIs()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", apis)

	api, err := session.GetAPI(apis[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", api)

	meta, err := session.GetAPIMetadata(apis[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", meta)
}
