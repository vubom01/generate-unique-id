package generate

import "testing"

func TestNewUniqueIDGenerator(t *testing.T) {
	_, err := NewUniqueIDGenerator(0)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	_, err = NewUniqueIDGenerator(5000)
	if err == nil {
		t.Fatalf("no error creating NewNode, %s", err)
	}
}

func TestGenerateDuplicateID(t *testing.T) {
	node, _ := NewUniqueIDGenerator(1)

	var x, y ID
	for i := 0; i < 1000000; i++ {
		y = node.GenerateID()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}

func TestRace(t *testing.T) {
	node, _ := NewUniqueIDGenerator(1)
	go func() {
		for i := 0; i < 1000000000; i++ {

			_, _ = NewUniqueIDGenerator(1)
		}
	}()

	for i := 0; i < 4000; i++ {

		node.GenerateID()
	}
}

func TestInt64(t *testing.T) {
	node, err := NewUniqueIDGenerator(0)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	oID := node.GenerateID()
	i := oID.Int64()

	pID := ParseInt64(i)
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	mi := int64(1116766490855473152)
	pID = ParseInt64(mi)
	if pID.Int64() != mi {
		t.Fatalf("pID %v != mi %v", pID.Int64(), mi)
	}
}

func TestString(t *testing.T) {
	node, err := NewUniqueIDGenerator(0)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	oID := node.GenerateID()
	si := oID.String()

	pID, err := ParseString(si)
	if err != nil {
		t.Fatalf("error parsing, %s", err)
	}

	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := `1116766490855473152`
	_, err = ParseString(ms)
	if err != nil {
		t.Fatalf("error parsing, %s", err)
	}

	ms = `1112316766490855473152`
	_, err = ParseString(ms)
	if err == nil {
		t.Fatalf("no error parsing %s", ms)
	}
}

func TestBase2(t *testing.T) {
	node, err := NewUniqueIDGenerator(0)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	oID := node.GenerateID()
	i := oID.Base2()

	pID, err := ParseBase2(i)
	if err != nil {
		t.Fatalf("error parsing, %s", err)
	}
	if pID != oID {
		t.Fatalf("pID %v != oID %v", pID, oID)
	}

	ms := `111101111111101110110101100101001000000000000000000000000000`
	_, err = ParseBase2(ms)
	if err != nil {
		t.Fatalf("error parsing, %s", err)
	}

	ms = `1112316766490855473152`
	_, err = ParseBase2(ms)
	if err == nil {
		t.Fatalf("no error parsing %s", ms)
	}
}
