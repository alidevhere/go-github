package github

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestActionsService_ListCaches(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/caches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"page": "2"})
		fmt.Fprint(w,
			`{
				"total_count":1,
				"actions_caches":[{"id":1}]
			}`,
		)
	})

	opts := &ActionsCacheListOptions{Page: 2}
	ctx := context.Background()
	cacheList, _, err := client.Actions.ListCaches(ctx, "o", "r", opts)
	if err != nil {
		t.Errorf("Actions.ListCaches returned error: %v", err)
	}

	want := &ActionsCacheList{TotalCount: Int64(1), ActionsCaches: []*ActionsCache{{ID: Int64(1)}}}
	if !cmp.Equal(cacheList, want) {
		t.Errorf("Actions.ListCaches returned %+v, want %+v", cacheList, want)
	}

	const methodName = "ListCaches"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.ListCaches(ctx, "\n", "\n", opts)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.ListCaches(ctx, "o", "r", opts)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}
