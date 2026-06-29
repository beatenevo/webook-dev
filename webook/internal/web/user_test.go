package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordEncrypt(t *testing.T) {
	password := []byte("123456#hello")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	println(string(encrypted))
	err = bcrypt.CompareHashAndPassword(encrypted, []byte("123456#hello"))
	assert.NoError(t, err)
}

//func TestNil(t *testing.T) {
//	c := nil
//	claims := c.(*UserClaims)
//	println(claims.Uid)
//}
//func testTypeAssert(c any) {
//	claims := c.(*UserClaims)
//	println(claims.Uid)
//}
