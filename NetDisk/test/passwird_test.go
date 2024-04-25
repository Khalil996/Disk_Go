package test

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcryptPasswordHashing(t *testing.T) {
	// 这是我们想要哈希的密码
	password := "123456"

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("无法哈希密码: %v", err)
	}

	// 测试：验证哈希是否与密码匹配
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		t.Errorf("验证失败：密码和哈希不匹配")
	}
}
