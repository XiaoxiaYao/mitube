package dbops

import (
	"testing"
)

func clearTables() {
	//TODO: Drop tables
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("testusername", "testpwd")
	if err != nil {
		t.Errorf("Error of Adduser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	//TODO
}