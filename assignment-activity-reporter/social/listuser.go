package social

import "sort"

type ListUser struct {
	listUser []*User
}

func NewListUser(listUser []*User) *ListUser {
	users := ListUser{
		listUser: listUser,
	}
	return &users
}

func (users *ListUser) GetUser(name string) *User {
	for _, v := range users.listUser {
		if v.name == name {
			return v
		}
	}
	return nil
}

func (users *ListUser) AddUserToUsers(name string) *User {
	for _, v := range users.listUser {
		if v.name == name {
			return v
		}
	}
	newUser := NewUser(name)
	users.listUser = append(users.listUser, newUser)
	return newUser
}

func (users *ListUser) SortPhotoDesc() {
	sort.Slice(users.listUser, func(i, j int) bool {
		return len(users.listUser[i].photo.likeBy) > len(users.listUser[j].photo.likeBy)
	})
}

func (users *ListUser) GetTopThreePhotos() []*User {
	topPhotos := []*User{}
	MAX := 3
	for _, v := range users.listUser {
		if v.photo.isAvailable {
			topPhotos = append(topPhotos, v)
		}

		if len(topPhotos) == MAX {
			break
		}
	}
	return topPhotos
}
