package challenges

type challengeSpec struct {
	ID          string
	PublicTests []challengeTest
	HiddenTests []challengeTest
}

type challengeTest struct {
	Name string
	Code string
}

func lookupChallenge(id string) (challengeSpec, bool) {
	spec, ok := challengeCatalog[id]
	return spec, ok
}

var challengeCatalog = map[string]challengeSpec{
	"retry-context": {
		ID: "retry-context",
		PublicTests: []challengeTest{
			{
				Name: "TestFetchUser_Success",
				Code: `func TestFetchUser_Success(t *testing.T) {
	client := &fakeDoer{
		results: []doResult{
			{response: jsonResponse(http.StatusOK, ` + "`" + `{"id":42,"name":"Alex"}` + "`" + `)},
		},
	}
	setClient(client)

	user, err := FetchUser(context.Background(), 42)
	if err != nil {
		t.Fatalf("FetchUser returned error: %v", err)
	}
	if user.ID != 42 || user.Name != "Alex" {
		t.Fatalf("unexpected user: %+v", user)
	}
	if client.calls != 1 {
		t.Fatalf("expected 1 call, got %d", client.calls)
	}
}`,
			},
			{
				Name: "TestFetchUser_RetryOnTemporary",
				Code: `func TestFetchUser_RetryOnTemporary(t *testing.T) {
	client := &fakeDoer{
		results: []doResult{
			{err: temporaryError{}},
			{err: temporaryError{}},
			{response: jsonResponse(http.StatusOK, ` + "`" + `{"id":7,"name":"Retry"}` + "`" + `)},
		},
	}
	setClient(client)

	user, err := FetchUser(context.Background(), 7)
	if err != nil {
		t.Fatalf("FetchUser returned error: %v", err)
	}
	if user.ID != 7 || user.Name != "Retry" {
		t.Fatalf("unexpected user: %+v", user)
	}
	if client.calls != 3 {
		t.Fatalf("expected 3 calls, got %d", client.calls)
	}
}`,
			},
			{
				Name: "TestFetchUser_ContextCancel",
				Code: `func TestFetchUser_ContextCancel(t *testing.T) {
	client := &fakeDoer{
		results: []doResult{
			{err: temporaryError{}},
			{response: jsonResponse(http.StatusOK, ` + "`" + `{"id":1,"name":"late"}` + "`" + `)},
		},
	}
	setClient(client)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := FetchUser(ctx, 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if client.calls != 1 {
		t.Fatalf("expected 1 call after cancellation, got %d", client.calls)
	}
}`,
			},
		},
		HiddenTests: []challengeTest{
			{
				Name: "TestFetchUser_NonRetryable",
				Code: `func TestFetchUser_NonRetryable(t *testing.T) {
	client := &fakeDoer{
		results: []doResult{
			{err: errors.New("permanent error")},
			{response: jsonResponse(http.StatusOK, ` + "`" + `{"id":1,"name":"unexpected"}` + "`" + `)},
		},
	}
	setClient(client)

	_, err := FetchUser(context.Background(), 1)
	if err == nil {
		t.Fatal("expected non-retryable error")
	}
	if client.calls != 1 {
		t.Fatalf("expected 1 call for non-retryable error, got %d", client.calls)
	}
}`,
			},
			{
				Name: "TestFetchUser_MaxRetries",
				Code: `func TestFetchUser_MaxRetries(t *testing.T) {
	client := &fakeDoer{
		results: []doResult{
			{err: temporaryError{}},
			{err: temporaryError{}},
			{err: temporaryError{}},
			{response: jsonResponse(http.StatusOK, ` + "`" + `{"id":1,"name":"unexpected"}` + "`" + `)},
		},
	}
	setClient(client)

	_, err := FetchUser(context.Background(), 1)
	if err == nil {
		t.Fatal("expected error after max retries")
	}
	if client.calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", client.calls)
	}
}`,
			},
		},
	},
}
