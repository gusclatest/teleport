/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package srv

import (
	"errors"
	"fmt"
	"os/user"
	"slices"
	"strings"
	"testing"

	"github.com/gravitational/trace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/backend/memory"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/services/local"
)

type testHostUserBackend struct {
	// users: user -> []groups
	users map[string][]string
	// groups: group -> groupid
	groups map[string]string
	// sudoers: user -> entries
	sudoers map[string]string
	// userUID: user -> uid
	userUID map[string]string
	// userGID: user -> gid
	userGID map[string]string

	setUserGroupsCalls       int
	createHomeDirectoryCalls int
}

func newTestUserMgmt() *testHostUserBackend {
	return &testHostUserBackend{
		users:   map[string][]string{},
		groups:  map[string]string{},
		sudoers: map[string]string{},
		userUID: map[string]string{},
		userGID: map[string]string{},
	}
}

func (tm *testHostUserBackend) GetAllUsers() ([]string, error) {
	keys := make([]string, 0, len(tm.users))
	for key := range tm.users {
		keys = append(keys, key)
	}
	return keys, nil
}

func (tm *testHostUserBackend) Lookup(username string) (*user.User, error) {
	if _, ok := tm.users[username]; !ok {
		return nil, user.UnknownUserError(username)
	}
	return &user.User{
		Username: username,
		Uid:      tm.userUID[username],
		Gid:      tm.userGID[username],
	}, nil
}

func (tm *testHostUserBackend) LookupGroup(groupname string) (*user.Group, error) {
	gid, ok := tm.groups[groupname]
	if !ok {
		return nil, user.UnknownGroupError(groupname)
	}
	return &user.Group{
		Gid:  gid,
		Name: groupname,
	}, nil
}

func (tm *testHostUserBackend) LookupGroupByID(gid string) (*user.Group, error) {
	for groupName, groupGid := range tm.groups {
		if groupGid == gid {
			return &user.Group{
				Gid:  gid,
				Name: groupName,
			}, nil
		}
	}
	return nil, user.UnknownGroupIdError(gid)
}

func (tm *testHostUserBackend) SetUserGroups(name string, groups []string) error {
	tm.setUserGroupsCalls++
	if _, ok := tm.users[name]; !ok {
		return trace.NotFound("User %q doesn't exist", name)
	}
	tm.users[name] = groups
	return nil
}

func (tm *testHostUserBackend) UserGIDs(u *user.User) ([]string, error) {
	ids := make([]string, 0, len(tm.users[u.Username])+1)
	for _, id := range tm.users[u.Username] {
		ids = append(ids, tm.groups[id])
	}
	// Include primary group.
	ids = append(ids, u.Gid)
	return ids, nil
}

func (tm *testHostUserBackend) CreateGroup(group, gid string) error {
	_, ok := tm.groups[group]
	if ok {
		return trace.AlreadyExists("Group %q, already exists", group)
	}
	if gid == "" {
		gid = fmt.Sprint(len(tm.groups) + 1)
	}
	tm.groups[group] = gid
	return nil
}

func (tm *testHostUserBackend) CreateUser(user string, groups []string, home, uid, gid string) error {
	_, ok := tm.users[user]
	if ok {
		return trace.AlreadyExists("Group %q, already exists", user)
	}
	if uid == "" {
		uid = fmt.Sprint(len(tm.users) + 1)
	}
	if gid == "" {
		gid = fmt.Sprint(len(tm.groups) + 1)
	}
	// Ensure that the user has a primary group. It's OK if it already exists.
	_ = tm.CreateGroup(user, gid)
	tm.users[user] = groups
	tm.userUID[user] = uid
	tm.userGID[user] = gid
	return nil
}

func (tm *testHostUserBackend) DeleteUser(user string) error {
	delete(tm.users, user)
	return nil
}

func (tm *testHostUserBackend) CreateHomeDirectory(user, uid, gid string) error {
	tm.createHomeDirectoryCalls++
	return nil
}

func (tm *testHostUserBackend) GetDefaultHomeDirectory(user string) (string, error) {
	return "", nil
}

// RemoveSudoersFile implements HostUsersBackend
func (tm *testHostUserBackend) RemoveSudoersFile(user string) error {
	delete(tm.sudoers, user)
	return nil
}

// CheckSudoers implements HostUsersBackend
func (*testHostUserBackend) CheckSudoers(contents []byte) error {
	if strings.Contains(string(contents), "validsudoers") {
		return nil
	}
	return errors.New("invalid")
}

// WriteSudoersFile implements HostUsersBackend
func (tm *testHostUserBackend) WriteSudoersFile(user string, entries []byte) error {
	entry := strings.TrimSpace(string(entries))
	err := tm.CheckSudoers([]byte(entry))
	if err != nil {
		return trace.Wrap(err)
	}
	tm.sudoers[user] = entry
	return nil
}

func (*testHostUserBackend) RemoveAllSudoersFiles() error {
	return nil
}

var _ HostUsersBackend = &testHostUserBackend{}
var _ HostSudoersBackend = &testHostUserBackend{}

func TestUserMgmt_CreateTemporaryUser(t *testing.T) {
	t.Parallel()

	users, backend := initBackend(t, nil)

	userinfo := services.HostUsersInfo{
		Groups: []string{"hello", "sudo"},
		Mode:   services.HostUserModeDrop,
	}
	// create a user with some groups
	closer, err := users.UpsertUser("bob", userinfo)
	require.NoError(t, err)
	// NOTE (eriktate): assert.Nil and assert.NotNil will pass for nil interfaces where nilInterface != nil.
	// assert.Equal and assert.NotEqual perform the same comparisons we would in non-test code and are safer
	// for interface types.
	//
	// https://glucn.com/posts/2019-05-20-golang-an-interface-holding-a-nil-value-is-not-nil
	require.NotEqual(t, nil, closer, "user closer was nil")

	// temproary users must always include the teleport-service group
	require.Equal(t, []string{
		"hello", "sudo", types.TeleportDropGroup,
	}, backend.users["bob"])

	// try create the same user again
	secondCloser, err := users.UpsertUser("bob", userinfo)
	require.NoError(t, err)
	require.NotEqual(t, nil, secondCloser)

	// Close will remove the user if the user is in the teleport-system group
	require.NoError(t, closer.Close())
	require.NotContains(t, backend.users, "bob")

	backend.CreateGroup("testgroup", "")
	backend.CreateUser("simon", []string{}, "", "", "")

	// an existing, unmanaged user should not be changed
	closer, err = users.UpsertUser("simon", userinfo)
	require.ErrorIs(t, err, unmanagedUserErr)
	require.Equal(t, nil, closer)
}

func TestUserMgmtSudoers_CreateTemporaryUser(t *testing.T) {
	t.Parallel()

	backend := newTestUserMgmt()
	bk, err := memory.New(memory.Config{})
	require.NoError(t, err)
	pres := local.NewPresenceService(bk)
	users := HostUserManagement{
		backend: backend,
		storage: pres,
	}
	sudoers := HostSudoersManagement{
		backend: backend,
	}

	closer, err := users.UpsertUser("bob", services.HostUsersInfo{
		Groups: []string{"hello", "sudo"},
		Mode:   services.HostUserModeDrop,
	})
	require.NoError(t, err)
	require.NotEqual(t, nil, closer)

	require.Empty(t, backend.sudoers)
	sudoers.WriteSudoers("bob", []string{"validsudoers"})
	require.Equal(t, map[string]string{"bob": "bob validsudoers"}, backend.sudoers)
	sudoers.RemoveSudoers("bob")

	require.NoError(t, closer.Close())
	require.Empty(t, backend.sudoers)

	t.Run("no teleport-service group", func(t *testing.T) {
		backend := newTestUserMgmt()
		users := HostUserManagement{
			backend: backend,
			storage: pres,
		}
		// test user already exists but teleport-service group has not yet
		// been created
		backend.CreateUser("testuser", nil, "", "", "")
		_, err := users.UpsertUser("testuser", services.HostUsersInfo{
			Mode: services.HostUserModeDrop,
		})
		require.ErrorIs(t, err, unmanagedUserErr)
		backend.CreateGroup(types.TeleportDropGroup, "")
		_, err = users.UpsertUser("testuser", services.HostUsersInfo{
			Mode: services.HostUserModeDrop,
		})
		require.ErrorIs(t, err, unmanagedUserErr)
	})
}

func TestUserMgmt_DeleteAllTeleportSystemUsers(t *testing.T) {
	t.Parallel()

	type userAndGroups struct {
		user   string
		groups []string
	}

	usersDB := []userAndGroups{
		{"fgh", []string{types.TeleportDropGroup}},
		{"xyz", []string{types.TeleportDropGroup}},
		{"pqr", []string{"not-deleted"}},
		{"abc", []string{"not-deleted"}},
	}

	remainingUsers := []string{"pqr", "abc"}

	mgmt := newTestUserMgmt()
	bk, err := memory.New(memory.Config{})
	require.NoError(t, err)
	pres := local.NewPresenceService(bk)
	users := HostUserManagement{
		backend:   mgmt,
		storage:   pres,
		userGrace: 0,
	}

	for _, user := range usersDB {
		for _, group := range user.groups {
			mgmt.CreateGroup(group, "")
		}
		if slices.Contains(user.groups, types.TeleportDropGroup) {
			users.UpsertUser(user.user, services.HostUsersInfo{
				Groups: user.groups,
				Mode:   services.HostUserModeDrop,
			})
		} else {
			mgmt.CreateUser(user.user, user.groups, "", "", "")
		}
	}
	require.NoError(t, users.DeleteAllUsers())
	resultingUsers, err := mgmt.GetAllUsers()
	require.NoError(t, err)

	require.ElementsMatch(t, remainingUsers, resultingUsers)

	users = HostUserManagement{
		backend: newTestUserMgmt(),
		storage: pres,
	}
	// teleport-system group doesnt exist, DeleteAllUsers will return nil, instead of erroring
	require.NoError(t, users.DeleteAllUsers())
}

func TestSudoersSanitization(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		user         string
		userExpected string
	}{
		{
			user:         "testuser",
			userExpected: "testuser",
		},
		{
			user:         "test.user",
			userExpected: "test_user",
		},
		{
			user:         "test.us~er",
			userExpected: "test_us_er",
		},
		{
			user:         "test../../us~er",
			userExpected: "test______us_er",
		},
	} {
		actual := sanitizeSudoersName(tc.user)
		require.Equal(t, tc.userExpected, actual)
	}
}

func TestIsUnknownGroupError(t *testing.T) {
	unknownGroupName := "unknown"
	for _, tc := range []struct {
		err                 error
		isUnknownGroupError bool
	}{
		{
			err:                 user.UnknownGroupError(unknownGroupName),
			isUnknownGroupError: true,
		}, {
			err:                 fmt.Errorf("lookup groupname %s: no such file or directory", unknownGroupName),
			isUnknownGroupError: true,
		},
	} {
		require.Equal(t, tc.isUnknownGroupError, isUnknownGroupError(tc.err, unknownGroupName))
	}
}

func Test_UpdateUserGroups_Keep(t *testing.T) {
	t.Parallel()

	allGroups := []string{"foo", "bar", "baz", "quux"}
	users, backend := initBackend(t, allGroups)

	userinfo := services.HostUsersInfo{
		Groups: slices.Clone(allGroups[:2]),
		Mode:   services.HostUserModeKeep,
	}

	// Create user
	closer, err := users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportKeepGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportDropGroup)

	// Update user with new groups.
	userinfo.Groups = slices.Clone(allGroups[2:])

	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportKeepGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportDropGroup)

	// Upsert again with same groups should not call SetUserGroups.
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportKeepGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportDropGroup)

	// Do not convert the managed user to static.
	userinfo.Mode = services.HostUserModeStatic
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportKeepGroup), backend.users["alice"])

	// Updates with INSECURE_DROP mode should convert the managed user
	userinfo.Mode = services.HostUserModeDrop
	userinfo.Groups = slices.Clone(allGroups[:2])
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Equal(t, 2, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportDropGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportKeepGroup)
}

func Test_UpdateUserGroups_Drop(t *testing.T) {
	t.Parallel()

	allGroups := []string{"foo", "bar", "baz", "quux"}
	users, backend := initBackend(t, allGroups)

	userinfo := services.HostUsersInfo{
		Groups: slices.Clone(allGroups[:2]),
		Mode:   services.HostUserModeDrop,
	}

	// Create user
	closer, err := users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportDropGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportKeepGroup)

	// Update user with new groups.
	userinfo.Groups = slices.Clone(allGroups[2:])

	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportDropGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportKeepGroup)

	// Upsert again with same groups should not call SetUserGroups.
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportDropGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportKeepGroup)

	// Do not convert the managed user to static.
	userinfo.Mode = services.HostUserModeStatic
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportDropGroup), backend.users["alice"])

	// Updates with KEEP mode should convert the ephemeral user
	userinfo.Mode = services.HostUserModeKeep
	userinfo.Groups = slices.Clone(allGroups[:2])
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 2, backend.setUserGroupsCalls)
	assert.Equal(t, 1, backend.createHomeDirectoryCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportKeepGroup), backend.users["alice"])
	assert.NotContains(t, backend.users["alice"], types.TeleportDropGroup)
}

func Test_UpdateUserGroups_Static(t *testing.T) {
	t.Parallel()

	allGroups := []string{"foo", "bar", "baz", "quux"}
	users, backend := initBackend(t, allGroups)
	userinfo := services.HostUsersInfo{
		Groups: slices.Clone(allGroups[:2]),
		Mode:   services.HostUserModeStatic,
	}

	// Create user.
	closer, err := users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportStaticGroup), backend.users["alice"])

	// Update user with new groups.
	userinfo.Groups = slices.Clone(allGroups[2:])
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportStaticGroup), backend.users["alice"])

	// Upsert again with same groups should not call SetUserGroups.
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(userinfo.Groups, types.TeleportStaticGroup), backend.users["alice"])

	// Do not convert to KEEP.
	userinfo.Mode = services.HostUserModeKeep
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(slices.Clone(allGroups[2:]), types.TeleportStaticGroup), backend.users["alice"])

	// Do not convert to INSECURE_DROP.
	userinfo.Mode = services.HostUserModeDrop
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, append(slices.Clone(allGroups[2:]), types.TeleportStaticGroup), backend.users["alice"])
}

func Test_DontManageExistingUser(t *testing.T) {
	t.Parallel()

	allGroups := []string{"foo", "bar", "baz", "quux"}
	users, backend := initBackend(t, allGroups)

	assert.NoError(t, backend.CreateUser("alice", allGroups, "", "", ""))

	userinfo := services.HostUsersInfo{
		Groups: allGroups[:2],
		Mode:   services.HostUserModeDrop,
	}

	// Update user in DROP mode
	closer, err := users.UpsertUser("alice", userinfo)
	assert.ErrorIs(t, err, unmanagedUserErr)
	assert.Equal(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, allGroups, backend.users["alice"])

	// Update user in KEEP mode
	userinfo.Mode = services.HostUserModeKeep
	closer, err = users.UpsertUser("alice", userinfo)
	assert.ErrorIs(t, err, unmanagedUserErr)
	assert.Equal(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, allGroups, backend.users["alice"])

	// Update static user
	userinfo.Mode = services.HostUserModeStatic
	closer, err = users.UpsertUser("alice", userinfo)
	assert.ErrorIs(t, err, unmanagedUserErr)
	assert.Equal(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, allGroups, backend.users["alice"])
}

func Test_DontUpdateUnmanagedUsers(t *testing.T) {
	t.Parallel()

	allGroups := []string{"foo", "bar", "baz", "quux"}
	users, backend := initBackend(t, allGroups)

	assert.NoError(t, backend.CreateUser("alice", allGroups[2:], "", "", ""))
	tests := []struct {
		name     string
		userinfo services.HostUsersInfo
	}{
		{
			name: "keep",
			userinfo: services.HostUsersInfo{
				Groups: allGroups[:2],
				Mode:   services.HostUserModeKeep,
			},
		},
		{
			name: "drop",
			userinfo: services.HostUsersInfo{
				Groups: allGroups[:2],
				Mode:   services.HostUserModeDrop,
			},
		},
		{
			name: "static",
			userinfo: services.HostUsersInfo{
				Groups: allGroups[:2],
				Mode:   services.HostUserModeStatic,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			closer, err := users.UpsertUser("alice", tc.userinfo)
			assert.ErrorIs(t, err, unmanagedUserErr)
			assert.Equal(t, nil, closer)
			assert.Zero(t, backend.setUserGroupsCalls)
			assert.ElementsMatch(t, allGroups[2:], backend.users["alice"])
		})
	}
}

// teleport-keep can be included explicitly in the Groups slice in order to flag an
// existing user as being managed by teleport
func Test_AllowExplicitlyManageExistingUsers(t *testing.T) {
	t.Parallel()

	allGroups := []string{"foo", types.TeleportKeepGroup, types.TeleportDropGroup}
	users, backend := initBackend(t, allGroups)

	assert.NoError(t, backend.CreateUser("alice-keep", []string{}, "", "", ""))
	assert.NoError(t, backend.CreateUser("alice-drop", []string{}, "", "", ""))
	userinfo := services.HostUsersInfo{
		Groups: slices.Clone(allGroups),
		Mode:   services.HostUserModeKeep,
	}

	// Take ownership of existing user when in KEEP mode
	closer, err := users.UpsertUser("alice-keep", userinfo)
	assert.NoError(t, err)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	// slice off the end because teleport-system should be explicitly excluded
	assert.ElementsMatch(t, allGroups[:2], backend.users["alice-keep"])
	assert.NotContains(t, backend.users["alice-keep"], types.TeleportDropGroup)

	// Don't take ownership of existing user when in DROP mode
	userinfo.Mode = services.HostUserModeDrop
	closer, err = users.UpsertUser("alice-drop", userinfo)
	assert.ErrorIs(t, err, unmanagedUserErr)
	assert.Equal(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.Empty(t, backend.users["alice-drop"])

	// Don't assign teleport-keep to users created in DROP mode
	userinfo.Mode = services.HostUserModeDrop
	closer, err = users.UpsertUser("bob", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Equal(t, 1, backend.setUserGroupsCalls)
	assert.ElementsMatch(t, []string{"foo", types.TeleportDropGroup}, backend.users["bob"])
	assert.NotContains(t, backend.users["bob"], types.TeleportKeepGroup)
}

func initBackend(t *testing.T, groups []string) (HostUserManagement, *testHostUserBackend) {
	backend := newTestUserMgmt()
	bk, err := memory.New(memory.Config{})
	require.NoError(t, err)
	pres := local.NewPresenceService(bk)
	users := HostUserManagement{
		backend: backend,
		storage: pres,
	}

	for _, group := range groups {
		require.NoError(t, backend.CreateGroup(group, ""))
	}
	return users, backend
}

func TestCreateUserWithExistingPrimaryGroup(t *testing.T) {
	t.Parallel()
	backend := newTestUserMgmt()
	bk, err := memory.New(memory.Config{})
	require.NoError(t, err)
	pres := local.NewPresenceService(bk)
	users := HostUserManagement{
		backend: backend,
		storage: pres,
	}

	existingGroups := []string{"alice", "simon"}
	for _, group := range existingGroups {
		require.NoError(t, backend.CreateGroup(group, ""))
	}

	userinfo := services.HostUsersInfo{
		Groups: []string{},
		Mode:   services.HostUserModeDrop,
	}

	// create a user without an existing primary group
	closer, err := users.UpsertUser("bob", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)

	// create a user with primary group defined in userinfo.Groups, but not yet on the host
	userinfo.Groups = []string{"fred"}
	closer, err = users.UpsertUser("fred", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)

	// create a user with primary group defined in userinfo.Groups that already exists on the host
	userinfo.Groups = []string{"alice"}
	closer, err = users.UpsertUser("alice", userinfo)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)

	// create a user with primary group that already exists on the host but is not defined in userinfo.Groups
	userinfo.Groups = []string{""}
	closer, err = users.UpsertUser("simon", userinfo)
	assert.True(t, trace.IsAlreadyExists(err))
	assert.Contains(t, err.Error(), "conflicts with an existing group")
	assert.Equal(t, nil, closer)
	assert.Zero(t, backend.setUserGroupsCalls)
}
