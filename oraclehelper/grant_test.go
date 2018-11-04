package oraclehelper

import (
	"log"
	"testing"
)

func TestGrantServiceObjectGrants(t *testing.T) {
	objGrant := ResourceGrantObjectPrivilege{
		Grantee:    "GRANTTST01",
		Privilege:  []string{"SELECT"},
		Owner:      "SYSTEM",
		ObjectName: "TEST",
	}
	objGrant2 := ResourceGrantObjectPrivilege{
		Grantee:    "GRANTTST01",
		Privilege:  []string{"SELECT", "UPDATE"},
		Owner:      "SYSTEM",
		ObjectName: "TEST",
	}
	c.UserService.CreateUser(ResourceUser{Username: "GRANTTST01"})
	c.GrantService.GrantObjectPrivilege(objGrant)
	grants, err := c.GrantService.ReadGrantObjectPrivilege(objGrant)
	if err != nil {
		log.Fatalf("failed to read role, errormsg: %v\n", err)
	}

	if grants.Privileges[0] != "SELECT" {
		t.Errorf("Wanted: %s gott: %s", "SELECT", grants.Privileges[0])
	}
	c.GrantService.RevokeObjectPrivilege(objGrant)
	grants, err = c.GrantService.ReadGrantObjectPrivilege(objGrant)
	if err != nil {
		log.Fatalf("failed to read role, errormsg: %v\n", err)
	}
	if len(grants.Privileges) != 0 {
		t.Errorf("Wanted: %d gott: %d", 0, len(grants.Privileges))
	}

	//Test grant multiple privs
	c.GrantService.GrantObjectPrivilege(objGrant2)
	grants, err = c.GrantService.ReadGrantObjectPrivilege(objGrant2)
	if err != nil {
		log.Fatalf("failed to read role, errormsg: %v\n", err)
	}

	if len(grants.Privileges) != 2 {
		t.Errorf("Wanted: %d gott: %d", 2, len(grants.Privileges))
	}

	// Revoke insert from
	c.GrantService.RevokeObjectPrivilege(objGrant)
	grants, err = c.GrantService.ReadGrantObjectPrivilege(objGrant)
	if err != nil {
		log.Fatalf("failed to read role, errormsg: %v\n", err)
	}
	if len(grants.Privileges) != 1 {
		t.Errorf("Wanted: %d gott: %d", 1, len(grants.Privileges))
	}
	//Clean up
	c.UserService.DropUser(ResourceUser{Username: "GRANTTST01"})
}

func TestGrantServiceSysPrivsGrants(t *testing.T) {
	username := "GRANTSYSPRIVTEST"
	sysPrivs := ResourceGrantSystemPrivilege{
		Grantee:   username,
		Privilege: "CREATE SESSION",
	}
	c.UserService.CreateUser(ResourceUser{Username: username})

	err := c.GrantService.GrantSysPriv(sysPrivs)
	if err != nil {
		t.Errorf("GrantSysPriv FAiled")
	}
	userSysPrivs, err := c.GrantService.ReadGrantSysPrivs(sysPrivs)
	if err != nil {
		t.Errorf("ReadGrantSysPrivs Failed")
	}
	if value, ok := userSysPrivs["CREATE SESSION"]; !ok {
		t.Errorf("CREATE SESSION not exists value: %v\n", value)
	}
	//Clean up
	c.UserService.DropUser(ResourceUser{Username: username})

}

func TestGrantServiceRolePrivs(t *testing.T) {
	username := "GRANTROLEPRIVTEST"
	resourceUsername := ResourceUser{Username: username}
	dbRole := ResourceRole{
		Role: "TESTROLE",
	}
	rolePrivs := ResourceGrantRolePrivilege{
		Grantee: username,
		Role:    dbRole.Role,
	}
	c.UserService.CreateUser(resourceUsername)
	c.RoleService.CreateRole(dbRole)

	err := c.GrantService.GrantRolePriv(rolePrivs)
	if err != nil {
		t.Errorf("GrantRolePriv Failed")
	}

	roles, err := c.GrantService.ReadGrantRolePrivs(rolePrivs)
	if err != nil {
		t.Errorf("ReadGrantRolePrivs Failed")
	}
	if value, ok := roles[dbRole.Role]; !ok {
		t.Errorf("%s not exists value: %v\n", dbRole.Role, value)
	}

	c.RoleService.DropRole(dbRole)
	c.UserService.DropUser(resourceUsername)
}
