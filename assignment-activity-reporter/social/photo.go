package social

type Photo struct {
	isAvailable bool
	likeBy      []*User
}

func NewPhoto() Photo {
	return Photo{
		isAvailable: false,
		likeBy:      []*User{},
	}
}

func (photo *Photo) GetTotalLikesPhoto() int {
	return len(photo.likeBy)
}
