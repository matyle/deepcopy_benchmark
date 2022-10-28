package deepcopy

import "testing"

func BenchmarkGob(b *testing.B) {
	books := make([]*Book, 1)
	country := 1156
	author := AuthorInfo{"David", 38, &country}
	price := make(map[string]string)
	price["Europe"] = "$56"
	books[0] = &Book{"Tutorial", &author, 2020, []string{"math", "art"}, price}
	for i := 0; i < b.N; i++ {
		booksCpy := make([]Book, 0, 1)
		_ = DeepCopyByGob(&booksCpy, books)
		// b.Log(booksCpy)
		// check(err)
	}
}

func BenchmarkJson(b *testing.B) {
	books := make([]*Book, 1)
	country := 1156
	author := AuthorInfo{"David", 38, &country}
	price := make(map[string]string)
	price["Europe"] = "$56"
	books[0] = &Book{"Tutorial", &author, 2020, []string{"math", "art"}, price}
	for i := 0; i < b.N; i++ {
		// booksCpy := make([]Book, 0, 1)
		_, _ = DeepCopyByJson(books)
		// b.Log(boo)
		// check(err)
	}
}

func BenchmarkCustom(b *testing.B) {
	books := make([]*Book, 1)
	country := 1156
	author := AuthorInfo{"David", 38, &country}
	price := make(map[string]string)
	price["Europe"] = "$56"
	books[0] = &Book{"Tutorial", &author, 2020, []string{"math", "art"}, price}
	for i := 0; i < b.N; i++ {
		// booksCpy := make([]Book, 0, 1)
		DeepCopyByCustom(books)
		// b.Log(boo)
		// check(err)
	}
}
