package expiring_store

import (
	"testing"
	"time"
)

var (
	es = New(int64(time.Second))
)

func TestKeyGeneration(t *testing.T) {

	key1 := generateKey() // es.Set("1")
	key2 := generateKey()

	if key1 == key2 || key2 == "" {
		t.Errorf("Generated keys are the same! %s == %s",
			key1,
			key2)
	}
}

func TestRW(t *testing.T) {

	key := es.Set("woop")

	if key == "" {
		t.Errorf("Key is invalid! %s", key)
	}

	val, _ := es.Get(key)

	if val != "woop" {
		t.Errorf("Couldn't retreive value! %s", val)
	}
}

func TestExpiring(t *testing.T) {
	key := es.Set("woop")

	time.Sleep(500 * time.Millisecond)

	_, stat := es.Get(key)
	if !stat {
		t.Errorf("key shouldn't have expired!")
	}

	time.Sleep(2 * time.Second)

	val, stat := es.Get(key)

	if val != "" || stat {
		t.Errorf("Value should have expired! %s %s", val, stat)
	}
}
