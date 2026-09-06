package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		// Queue starting at 0 should match number of items in the queue
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %d, want %d", len(q.items), i)
		}
		// Should be append item to queue
		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Failed: Should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		// There should be items in the queue
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to get item from queue")
		}
		// Item returned should match item index as it was the same order in,
		//  thus satisfying definition of a queue
		if item != i {
			t.Errorf("Got item in wrong order: %d, wand %d", item, i)
		}
	}
	// Queue is empty
	item, ok := q.Next()
	if ok {
		t.Errorf("Should not be any more items in queue, got: %d", item)
	}
}
