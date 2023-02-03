package social

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowSuccess(t *testing.T) {
	t.Run("should return no error if the user never follow before", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")

		err := user2.SetUserFollower(user1)
		assert.Nil(t, err)
	})

	t.Run("should return no error if the user never like photo before", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")

		user2.SetUserFollower(user1)
		err := user2.SetLikePhoto(user1)

		assert.Nil(t, err)
	})
}

func TestFollowUnSuccess(t *testing.T) {
	t.Run("should return error  if the user already followed before", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")

		user2.SetUserFollower(user1)
		err := user2.SetUserFollower(user1)
		assert.NotNil(t, err)
	})

	t.Run("should return error  if the user already liked before", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")

		user2.SetUserFollower(user1)
		user2.SetLikePhoto(user1)
		err := user2.SetLikePhoto(user1)

		assert.NotNil(t, err)
	})
}

func TestCheckLikes(t *testing.T) {
	t.Run("should return error if the user like photo but he is not a follower", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")

		err := user2.CheckLikes(user1)

		assert.NotNil(t, err)
	})

	t.Run("should return nil if the user like photo but he is a follower", func(t *testing.T) {
		user1 := NewUser("Alice")
		user2 := NewUser("Bob")

		user2.SetUserFollower(user1)
		err := user2.CheckLikes(user1)

		assert.Nil(t, err)
	})
}

func TestUserPhoto(t *testing.T) {
	t.Run("should return ErrUploadTwice if the user like photo but already like the photo", func(t *testing.T) {
		user1 := NewUser("Alice")
		photo := NewPhoto()

		user1.SetUserPhoto(photo)
		err := user1.SetUserPhoto(photo)

		assert.Error(t, err, ErrUploadTwice)
	})
}

func TestNotifyAll(t *testing.T) {
	t.Run("should notify all when follower do action", func(t *testing.T) {
		alice := NewUser("Alice")
		bob := NewUser("Bob")

		alice.SetUserFollower(bob)
		alice.NotifyAll([]string{"alice", "uploaded", "photo"})

		assert.Contains(t, bob.GetActivities(), "alice uploaded photo")
	})
}
