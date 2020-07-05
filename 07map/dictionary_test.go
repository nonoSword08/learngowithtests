package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("execpet to get an error")
		}

		assertError(t, err, want)
	})

	got, _ := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)

}

func TestAdd(t *testing.T) {
	dictionnary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	
	dictionnary.Add(word, definition)

	assertDefinition(t, dictionnary, word, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if definition != got {
        t.Errorf("got %q want %q", got, definition)
    }
}