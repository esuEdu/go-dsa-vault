package linkedlist

import "testing"

func TestAppendToEmptyList(t *testing.T) {
	l := New()
	l.Append(1)

	if l.Head == nil {
		t.Fatal("expected head to be set, got nil")
	}
	if l.Head.Data != 1 {
		t.Errorf("expected head data 1, got %d", l.Head.Data)
	}
	if l.Length != 1 {
		t.Errorf("expected length 1, got %d", l.Length)
	}
}

func TestAppendMultiple(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if l.Length != 3 {
		t.Errorf("expected length 3, got %d", l.Length)
	}

	got := l.String()
	want := "1 -> 2 -> 3 -> nil"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestPrepend(t *testing.T) {
	l := New()
	l.Append(2)
	l.Prepend(1)

	if l.Head.Data != 1 {
		t.Errorf("expected head data 1, got %d", l.Head.Data)
	}
	if l.Length != 2 {
		t.Errorf("expected length 2, got %d", l.Length)
	}

	got := l.String()
	want := "1 -> 2 -> nil"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestDeleteHead(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	ok := l.Delete(1)

	if !ok {
		t.Fatal("expected Delete to return true")
	}
	if l.Head.Data != 2 {
		t.Errorf("expected new head data 2, got %d", l.Head.Data)
	}
	if l.Length != 2 {
		t.Errorf("expected length 2, got %d", l.Length)
	}
}

func TestDeleteMiddle(t *testing.T) {
	l := New()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	ok := l.Delete(2)

	if !ok {
		t.Fatal("expected Delete to return true")
	}

	got := l.String()
	want := "0 -> 1 -> 3 -> nil"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestDeleteTail(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	ok := l.Delete(3)

	if !ok {
		t.Fatal("expected Delete to return true")
	}

	got := l.String()
	want := "1 -> 2 -> nil"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestDeleteNotFound(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)

	ok := l.Delete(99)

	if ok {
		t.Error("expected Delete to return false for missing value")
	}
	if l.Length != 2 {
		t.Errorf("expected length unchanged at 2, got %d", l.Length)
	}
}

func TestDeleteEmptyList(t *testing.T) {
	l := New()

	ok := l.Delete(1)

	if ok {
		t.Error("expected Delete to return false on empty list")
	}
}

func TestSearchFound(t *testing.T) {
	l := New()
	l.Append(10)
	l.Append(20)
	l.Append(30)

	node := l.Search(20)

	if node == nil {
		t.Fatal("expected to find node, got nil")
	}
	if node.Data != 20 {
		t.Errorf("expected node data 20, got %d", node.Data)
	}
}

func TestSearchNotFound(t *testing.T) {
	l := New()
	l.Append(10)
	l.Append(20)

	node := l.Search(99)

	if node != nil {
		t.Errorf("expected nil, got node with data %d", node.Data)
	}
}

func TestStringEmptyList(t *testing.T) {
	l := New()

	got := l.String()
	want := "nil"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
