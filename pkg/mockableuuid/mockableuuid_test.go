package mockableuuid

import "testing"

func Test_mockableuuid(t *testing.T) {
	got := New().String()
	if got == "" {
		t.Fatalf("should not empty: got=%s", got)
	}
	t.Log(got)

	const fake = "fake"
	unMock := Mock(t, fake)
	if got := New().String(); got != fake {
		t.Fatalf("failed to mock: got=%s, fake=%s", got, fake)
	}

	unMock()
	if got := New().String(); got == fake {
		t.Fatalf("failed to unMock: got=%s, fake=%s", got, fake)
	}
}
