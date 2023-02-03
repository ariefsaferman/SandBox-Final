package social

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserToList(t *testing.T) {
	t.Run("should return user when register user is not registered", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")
		user3 := NewUser("John")
		listUser := NewListUser([]*User{})

		listUser.AddUserToUsers(user1.name)
		listUser.AddUserToUsers(user2.name)
		listUser.AddUserToUsers(user3.name)

		result := listUser.GetUser(user1.name)

		assert.Equal(t, user1, result)

	})
}

func TestTopPhoto(t *testing.T) {
	t.Run("should return user when register user is not registered", func(t *testing.T) {
		foto := NewPhoto()
		users := NewListUser([]*User{})
		bob := users.AddUserToUsers("Bob")
		alice := users.AddUserToUsers("Alice")
		john := users.AddUserToUsers("John")
		bill := users.AddUserToUsers("Bill")

		bob.SetUserFollower(alice)
		bob.SetUserFollower(john)
		bob.SetUserFollower(bill)
		alice.SetUserFollower(bob)
		alice.SetUserFollower(bill)
		alice.SetUserFollower(john)
		john.SetUserFollower(alice)
		john.SetUserFollower(bob)
		john.SetUserFollower(bill)
		bill.SetUserFollower(alice)
		bill.SetUserFollower(bob)
		bill.SetUserFollower(john)
		bill.SetUserPhoto(foto)
		alice.SetUserPhoto(foto)
		john.SetUserPhoto(foto)
		bill.SetLikePhoto(alice)
		bill.SetLikePhoto(john)
		alice.SetLikePhoto(john)
		users.SortPhotoDesc()
		trendPhoto := users.GetTopThreePhotos()

		assert.Len(t, trendPhoto, 3)
		assert.Equal(t, trendPhoto[0], bill)
		assert.Equal(t, trendPhoto[0].GetPhotoUser().GetTotalLikesPhoto(), 2)
		assert.Equal(t, trendPhoto[1], alice)
		assert.Equal(t, trendPhoto[1].GetPhotoUser().GetTotalLikesPhoto(), 1)
		assert.Equal(t, trendPhoto[2], john)
		assert.Equal(t, trendPhoto[2].GetPhotoUser().GetTotalLikesPhoto(), 0)
		assert.NotContains(t, trendPhoto, bob)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("should return nil when register user is not registered", func(t *testing.T) {

		listUser := NewListUser([]*User{})

		assert.Nil(t, listUser.GetUser("test-user"))

	})

	t.Run("should return user when register use is registered", func(t *testing.T) {
		bob := NewUser("Bob")
		alice := NewUser("Alice")
		listUser := NewListUser([]*User{bob, alice})

		user := listUser.AddUserToUsers(bob.GetName())

		assert.Equal(t, user, bob)
	})
}
