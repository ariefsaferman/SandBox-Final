package socialmedia

type User struct {
	name       string
	follower   []*User
	activities []*Activity
	photo      Photo
}

func NewUser(name string) *User {
	newUser := User{
		name: name,
	}
	return &newUser
}

func (user *User) GetName(u User) string {
	if user.searchUser(u) {
		return u.name
	}
	return ""
}

func (user *User) searchUser(u User) bool {
	for _, v := range user.follower {
		if v.name == u.name {
			return true
		}
	}
	return false
}

func (user *User) AddActivity(act Activity) {
	user.activities = append(user.activities, &act)
}
