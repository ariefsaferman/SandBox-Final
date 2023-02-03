package social

import "errors"

var ErrFollower = errors.New("you already followed the user")
var ErrInvalidKeyword = errors.New("invalid keyword")
var ErrLikesPhoto = errors.New("unable to like")
var ErrUploadTwice = errors.New("you cannot upload more than once")
var ErrDoubleLike = errors.New("you already like the photo")
