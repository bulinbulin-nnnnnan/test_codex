package usersort

import (
	"fmt"
	"testing"
	"time"
)

func generateTestData(n int) []*User {
	users := make([]*User, n)
	for i := 0; i < n; i++ {
		users[i] = &User{
			ID:               n - i,
			UserID:           fmt.Sprintf("user_%c", 'a'+i%26),
			Age:              18 + i%50,
			Name:             fmt.Sprintf("%c", 'A'+i%26),
			Grade:            60 + float64(i%40),
			RegistrationTime: time.Date(2024, 1, 1, 0, 0, i, 0, time.UTC),
		}
	}
	return users
}

func BenchmarkSortFunc(b *testing.B) {
	users := generateTestData(10000)
	option := SortOption{Field: FieldID, Order: Ascending}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortFunc(option)(users[0], users[1])
	}
}

func BenchmarkSort(b *testing.B) {
	option := SortOption{Field: FieldID, Order: Ascending}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := generateTestData(10000)
		Sort(users, option)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	users := generateTestData(10000)
	option := SortOption{Field: FieldID, Order: Ascending}
	Sort(users, option)
	target := &User{ID: 5000}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(users, target, option)
	}
}

func BenchmarkFindInsertPosition(b *testing.B) {
	users := generateTestData(10000)
	option := SortOption{Field: FieldID, Order: Ascending}
	Sort(users, option)
	target := &User{ID: 5000}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindInsertPosition(users, target, option)
	}
}

func BenchmarkInsert(b *testing.B) {
	option := SortOption{Field: FieldID, Order: Ascending}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := generateTestData(10000)
		Sort(users, option)
		target := &User{ID: 5000}
		Insert(users, target, option)
	}
}

func BenchmarkFindByField(b *testing.B) {
	users := generateTestData(10000)
	target := &User{UserID: "user_m"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindByField(users, target, FieldUserID)
	}
}

func BenchmarkDelete(b *testing.B) {
	option := SortOption{Field: FieldUserID, Order: Ascending}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := generateTestData(10000)
		target := &User{UserID: "user_m"}
		Delete(users, target, option)
	}
}

func TestSortFunc_NilUsers(t *testing.T) {
	option := SortOption{Field: FieldID, Order: Ascending}
	f := SortFunc(option)
	if f(nil, nil) {
		t.Error("SortFunc should return false when both users are nil")
	}
}

func TestSort_EmptySlice(t *testing.T) {
	var users []*User
	option := SortOption{Field: FieldID, Order: Ascending}
	Sort(users, option)
	if len(users) != 0 {
		t.Error("Sort should handle empty slice")
	}
}

func TestSort_SingleElement(t *testing.T) {
	users := []*User{{ID: 1}}
	option := SortOption{Field: FieldID, Order: Ascending}
	Sort(users, option)
	if len(users) != 1 || users[0].ID != 1 {
		t.Error("Sort should handle single element")
	}
}

func TestBinarySearch_EmptySlice(t *testing.T) {
	var users []*User
	target := &User{ID: 1}
	option := SortOption{Field: FieldID, Order: Ascending}
	pos := BinarySearch(users, target, option)
	if pos != -1 {
		t.Error("BinarySearch should return -1 for empty slice")
	}
}

func TestBinarySearch_NotFound(t *testing.T) {
	users := []*User{{ID: 1}, {ID: 3}, {ID: 5}}
	target := &User{ID: 2}
	option := SortOption{Field: FieldID, Order: Ascending}
	pos := BinarySearch(users, target, option)
	if pos != -1 {
		t.Error("BinarySearch should return -1 for not found")
	}
}

func TestBinarySearch_BoundaryValues(t *testing.T) {
	users := []*User{{ID: 1}, {ID: 2}, {ID: 3}}
	option := SortOption{Field: FieldID, Order: Ascending}

	pos := BinarySearch(users, &User{ID: 1}, option)
	if pos != 0 {
		t.Error("BinarySearch should find first element")
	}

	pos = BinarySearch(users, &User{ID: 3}, option)
	if pos != 2 {
		t.Error("BinarySearch should find last element")
	}
}

func TestFindInsertPosition_EmptySlice(t *testing.T) {
	var users []*User
	target := &User{ID: 1}
	option := SortOption{Field: FieldID, Order: Ascending}
	pos := FindInsertPosition(users, target, option)
	if pos != 0 {
		t.Error("FindInsertPosition should return 0 for empty slice")
	}
}

func TestInsert_NilTarget(t *testing.T) {
	users := []*User{{ID: 1}}
	option := SortOption{Field: FieldID, Order: Ascending}
	result := Insert(users, nil, option)
	if len(result) != 1 {
		t.Error("Insert should handle nil target")
	}
}

func TestDelete_NilTarget(t *testing.T) {
	users := []*User{{ID: 1}}
	option := SortOption{Field: FieldID, Order: Ascending}
	result := Delete(users, nil, option)
	if len(result) != 1 {
		t.Error("Delete should handle nil target")
	}
}

func TestDelete_NotFound(t *testing.T) {
	users := []*User{{ID: 1}, {ID: 2}}
	option := SortOption{Field: FieldID, Order: Ascending}
	result := Delete(users, &User{ID: 3}, option)
	if len(result) != 2 {
		t.Error("Delete should return original slice when not found")
	}
}

func TestDelete_EmptySlice(t *testing.T) {
	var users []*User
	option := SortOption{Field: FieldID, Order: Ascending}
	result := Delete(users, &User{ID: 1}, option)
	if len(result) != 0 {
		t.Error("Delete should handle empty slice")
	}
}

func TestFindByField_EmptySlice(t *testing.T) {
	var users []*User
	pos := FindByField(users, &User{ID: 1}, FieldID)
	if pos != -1 {
		t.Error("FindByField should return -1 for empty slice")
	}
}

func TestFindByField_NilTarget(t *testing.T) {
	users := []*User{{ID: 1}}
	pos := FindByField(users, nil, FieldID)
	if pos != -1 {
		t.Error("FindByField should return -1 for nil target")
	}
}

func TestSort_FieldEdgeCases(t *testing.T) {
	users := []*User{
		{ID: 0, UserID: "", Age: 0, Name: "", Grade: 0, RegistrationTime: time.Time{}},
		{ID: -1, UserID: "a", Age: -1, Name: "z", Grade: -1, RegistrationTime: time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)},
	}

	Sort(users, SortOption{Field: FieldID, Order: Ascending})
	if users[0].ID != -1 {
		t.Error("Sort should handle negative ID")
	}

	Sort(users, SortOption{Field: FieldAge, Order: Ascending})
	if users[0].Age != -1 {
		t.Error("Sort should handle negative Age")
	}

	Sort(users, SortOption{Field: FieldGrade, Order: Ascending})
	if users[0].Grade != -1 {
		t.Error("Sort should handle negative Grade")
	}

	Sort(users, SortOption{Field: FieldUserID, Order: Ascending})
	if users[0].UserID != "" {
		t.Error("Sort should handle empty UserID")
	}

	Sort(users, SortOption{Field: FieldName, Order: Ascending})
	if users[0].Name != "" {
		t.Error("Sort should handle empty Name")
	}

	Sort(users, SortOption{Field: FieldRegistrationTime, Order: Ascending})
	if !users[0].RegistrationTime.IsZero() {
		t.Error("Sort should handle zero time")
	}
}
