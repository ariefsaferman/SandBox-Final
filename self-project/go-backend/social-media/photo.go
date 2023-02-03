package socialmedia

type Photo struct {
	isAvailable bool
	isLikedBy   []*User
}

func NewPhoto() *Photo {
	newPhoto := Photo{
		isAvailable: false,
		isLikedBy:   []*User{},
	}
	return &newPhoto
}

func (photo *Photo) GetStatusPhoto() bool {
	return photo.isAvailable
}

func (photo *Photo) SetStatusPhoto() {
	photo.isAvailable = true
}

func (photo *Photo) UploadPhoto(p *Photo) {
	if !p.GetStatusPhoto() {
		p.SetStatusPhoto()
	}
}
