package social

import (
	"strings"
)

const LIKES = "likes"
const UPLOADED = "uploaded"

type User struct {
	name       string
	followers  []*User
	activities []string
	photo      *Photo
}

func NewUser(name string) *User {
	user := User{
		name:       name,
		followers:  []*User{},
		activities: make([]string, 0),
		photo:      &Photo{},
	}
	return &user
}

func (user *User) SetUserPhoto(photo Photo) error {
	if user.photo.isAvailable {
		return ErrUploadTwice
	}
	user.photo.isAvailable = true
	return nil
}

func (user *User) SetUserFollower(p *User) error {

	// cek user follower
	for _, v := range user.followers {
		if v.name == p.name {
			return ErrFollower
		}
	}

	// cek user list
	user.followers = append(user.followers, p)
	return nil
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) CheckLikes(p *User) error {
	for _, v := range user.followers {
		if v.name == p.name {
			return nil
		}
	}
	return ErrLikesPhoto
}

func (user *User) SetLikePhoto(p *User) error {
	for _, v := range user.photo.likeBy {
		if v == p {
			return ErrDoubleLike
		}
	}
	user.photo.likeBy = append(user.photo.likeBy, p)
	return nil
}

func (user *User) NotifyAll(strSlice []string) {

	//append follower activity
	for _, v := range user.followers {
		v.AddActivity(strSlice)
	}
}

func (user *User) GetActivities() []string {
	return user.activities
}

func (user *User) AddActivity(strSlice []string) {
	activity := strings.Join(strSlice, " ")
	user.activities = append(user.activities, activity)
}

func (user *User) GetPhotoUser() *Photo {
	return user.photo
}
