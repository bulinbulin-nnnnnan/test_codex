package main

import (
	"fmt"
	"test_codex/mathutil"
	"test_codex/usersort"
	"time"
)

func main() {
	fmt.Println("=== 数学工具包示例 ===")
	fmt.Println("Sum(1, 2, 3, 4, 5)        =", mathutil.Sum(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Average([10,20,30,40,50]) =", mathutil.Average(nums))

	if f, err := mathutil.Factorial(6); err == nil {
		fmt.Println("Factorial(6)               =", f)
	}
	if _, err := mathutil.Factorial(-3); err != nil {
		fmt.Println("Factorial(-3)              = 错误:", err)
	}

	fmt.Println("GCD(48, 18)                =", mathutil.GCD(48, 18))

	fmt.Println("Fibonacci(10)              =", mathutil.Fibonacci(10))

	fmt.Println("\n=== 用户排序工具包示例 ===")

	users := []*usersort.User{
		{ID: 3, UserID: "user_003", Age: 25, Name: "Charlie", Grade: 85.5, RegistrationTime: time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)},
		{ID: 1, UserID: "user_001", Age: 22, Name: "Alice", Grade: 92.0, RegistrationTime: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)},
		{ID: 4, UserID: "user_004", Age: 28, Name: "David", Grade: 78.0, RegistrationTime: time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)},
		{ID: 2, UserID: "user_002", Age: 24, Name: "Bob", Grade: 88.5, RegistrationTime: time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC)},
	}

	fmt.Println("\n原始用户列表:")
	printUsers(users)

	fmt.Println("\n按 ID 升序排序:")
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldID, Order: usersort.Ascending})
	printUsers(users)

	fmt.Println("\n按 Age 降序排序:")
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldAge, Order: usersort.Descending})
	printUsers(users)

	fmt.Println("\n按 Name 升序排序:")
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldName, Order: usersort.Ascending})
	printUsers(users)

	fmt.Println("\n按 Grade 降序排序:")
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldGrade, Order: usersort.Descending})
	printUsers(users)

	fmt.Println("\n按 RegistrationTime 升序排序:")
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldRegistrationTime, Order: usersort.Ascending})
	printUsers(users)

	fmt.Println("\n二分查找 user_002 (按 UserID 升序):")
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldUserID, Order: usersort.Ascending})
	target := &usersort.User{UserID: "user_002"}
	pos := usersort.BinarySearch(users, target, usersort.SortOption{Field: usersort.FieldUserID, Order: usersort.Ascending})
	if pos != -1 {
		fmt.Printf("找到用户 %s，位置: %d\n", users[pos].UserID, pos)
	} else {
		fmt.Println("未找到用户")
	}

	fmt.Println("\n插入新用户 (按 Age 升序):")
	newUser := &usersort.User{ID: 5, UserID: "user_005", Age: 23, Name: "Eve", Grade: 90.0, RegistrationTime: time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)}
	usersort.Sort(users, usersort.SortOption{Field: usersort.FieldAge, Order: usersort.Ascending})
	users = usersort.Insert(users, newUser, usersort.SortOption{Field: usersort.FieldAge, Order: usersort.Ascending})
	printUsers(users)

	fmt.Println("\n删除用户 user_003 (按 UserID 查找):")
	users = usersort.Delete(users, &usersort.User{UserID: "user_003"}, usersort.SortOption{Field: usersort.FieldUserID, Order: usersort.Ascending})
	printUsers(users)
}

func printUsers(users []*usersort.User) {
	for i, u := range users {
		fmt.Printf("[%d] ID:%d UserID:%s Age:%d Name:%s Grade:%.1f Time:%s\n",
			i, u.ID, u.UserID, u.Age, u.Name, u.Grade, u.RegistrationTime.Format("2006-01-02"))
	}
}
