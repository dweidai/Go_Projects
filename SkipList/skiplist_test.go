package skiplist

import "testing"

func TestRandomLevel(t *testing.T) {
	s := newSkiplist()
	for i := 0; i <= 20; i++ {
		ranLevel := s.randomLevel()
		if ranLevel > 20 || ranLevel < 0 {
			t.Errorf(" error found in random: %d", ranLevel)
		}
	}
}

func TestInsert(t *testing.T){
	s := newSkiplist()
	s.insert(30, 30)
	s.insert(1, 1)
	s.insert(2, 2)
	s.insert(3, 3)
	s.insert(31, 31)
	s.insert(11, 11)
	s.insert(33, 33)
	printlist(s)
}

func TestSearch(t *testing.T){
	s := newSkiplist()
	s.insert(20, 20)
	s.insert(1, 1)
	s.insert(2, 2)
	s.insert(3, 3)
	s.insert(31, 31)
	s.insert(11, 11)
	s.insert(33, 33)
	if val, _ := s.search(20); val != 20 {
		t.Errorf("Expect 20, got %d \n", val)
	}
	if val, _ := s.search(2); val != 2 {
		t.Errorf("Expect 2, got %d \n", val)
	}
	if val, _ := s.search(31); val != 31 {
		t.Errorf("Expect 31, got %d \n", val)
	}
	if val, _ := s.search(11); val != 11 {
		t.Errorf("Expect 11, got %d \n", val)
	}
	if val, _ := s.search(3); val != 3 {
		t.Errorf("Expect 3, got %d \n", val)
	}
	if val, _ := s.search(5); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}
	if val, _ := s.search(30); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}
}



func TestDelete(t *testing.T){
	s := newSkiplist()
	s.insert(20, 20)
	s.insert(1, 1)
	s.insert(2, 2)
	s.insert(3, 3)
	s.insert(31, 31)
	s.insert(11, 11)
	s.insert(33, 33)
	if val := s.delete(20); val != nil {
		t.Errorf("Deletion error")
	}
	if val,_ := s.search(20); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}

	if val := s.delete(1); val != nil {
		t.Errorf("Deletion error")
	}
	if val,_ := s.search(1); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}

	if val := s.delete(2); val != nil {
		t.Errorf("Deletion error")
	}
	if val,_ := s.search(2); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}

	if val := s.delete(11); val != nil {
		t.Errorf("Deletion error")
	}
	if val,_ := s.search(11); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}

	if val := s.delete(3); val != nil {
		t.Errorf("Deletion error")
	}
	if val,_ := s.search(3); val != nil {
		t.Errorf("Expect nil, got %d \n", val)
	}

	
}