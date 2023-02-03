package main

import (
	"assignment-activity-reporter/social"
	"assignment-activity-reporter/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FOLLOWS = "follows"
const UPLOADED = "uploaded"
const LIKES = "likes"
const PHOTO = "photo"
const ZERO = 0
const ONE = 1
const TWO = 2
const THREE = 3
const FOUR = 4

func menu() int {

	var input int
	fmt.Println("1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit")
	fmt.Print("Enter menu: ")
	fmt.Scanf("%d", &input)
	return input
}

func CheckCommand(command string) error {
	if command != FOLLOWS {
		return social.ErrInvalidKeyword
	}
	return nil
}

func setup(users *social.ListUser) {
	var uFollowing, command, uFollower string
	fmt.Print("Setup social graph: ")
	fmt.Scanf("%s %s %s", &uFollowing, &command, &uFollower)

	if err1 := CheckCommand(command); err1 != nil {
		return
	}

	userFollowing := users.AddUserToUsers(uFollowing)
	userFollower := users.AddUserToUsers(uFollower)
	err2 := userFollower.SetUserFollower(userFollowing)

	if err2 != nil {
		fmt.Println(err2)
	}

}

func action(foto social.Photo, users *social.ListUser) {
	var input, photo string
	fmt.Print("Enter user Actions: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print(social.ErrLikesPhoto)
		return
	}

	if len(input) < THREE {
		fmt.Println(social.ErrInvalidKeyword)
		fmt.Println("")
		return
	}

	strSlice := utils.SplitStrToSlice(input)
	uFollowing := users.GetUser(strSlice[ZERO])
	if uFollowing == nil {
		fmt.Println("unknown user " + strSlice[ZERO])
		return
	}
	command := strSlice[ONE]
	if len(strSlice) == THREE {
		photo = strSlice[TWO]

		if strings.Contains(photo, PHOTO) && strings.Contains(command, UPLOADED) {
			err = uFollowing.SetUserPhoto(foto)
			if err != nil {
				fmt.Println(err)
			}
			uFollowing.AddActivity(strSlice)
			uFollowing.NotifyAll(strSlice)
		}
	} else {
		uFollower := users.GetUser(strSlice[TWO])
		photo = strSlice[THREE]

		if uFollower != uFollowing {
			err = uFollower.CheckLikes(uFollowing)
			if err != nil {
				fmt.Println(err, uFollower.GetName(), photo)
				return
			}
		}

		err = uFollower.SetLikePhoto(uFollowing)
		if err != nil {
			fmt.Println(err)
		}
		uFollowing.AddActivity(strSlice)
		uFollowing.NotifyAll(strSlice)
		if uFollower != uFollowing {
			uFollower.AddActivity(strSlice)
		}
	}

}

func display(users *social.ListUser) {
	var input string
	fmt.Print("Display activity for: ")
	fmt.Scanf("%s", &input)
	fmt.Println(input + " activities: ")

	user := users.AddUserToUsers(input)
	for _, v := range user.GetActivities() {
		temp := strings.Split(v, " ")

		if len(temp) == THREE {
			if temp[ZERO] == input {
				fmt.Printf("You " + temp[ONE] + " " + temp[TWO])
				continue
			}
			fmt.Printf(v)
			continue
		} else {
			if input == temp[ZERO] && input == temp[TWO] {
				fmt.Printf("You liked your photo\n")
				continue
			} else if input == temp[ZERO] {
				fmt.Printf("You liked " + temp[TWO] + "'s photo\n")
				continue
			} else if input == temp[TWO] {
				fmt.Printf(temp[ZERO] + " liked your photo\n")
				continue
			}
			fmt.Printf(temp[ZERO] + " " + temp[ONE] + " " + temp[TWO] + "'s " + temp[THREE])
			continue
		}

	}
}

func trending(users *social.ListUser) {
	fmt.Println("Trending photos: ")
	users.SortPhotoDesc()
	trendPhotos := users.GetTopThreePhotos()

	for idx, v := range trendPhotos {
		fmt.Print(idx + 1)
		fmt.Print(". " + v.GetName() + " photo got ")
		fmt.Print(v.GetPhotoUser().GetTotalLikesPhoto())
		fmt.Println("")
	}
}

func main() {
	exit := false
	listUser := social.NewListUser([]*social.User{})
	foto := social.NewPhoto()

	for !exit {
		input := menu()

		switch input {
		case 1:
			setup(listUser)
		case 2:
			action(foto, listUser)
		case 3:
			display(listUser)
		case 4:
			trending(listUser)
		case 5:
			exit = true
		default:
			fmt.Println("invalid menu")
		}
		fmt.Println("------------------------------------------------------")
	}
}
