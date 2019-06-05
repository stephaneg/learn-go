package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}

}

func main() {
	page1 := storyPage{"you are standing in an open field west of a white house", nil}
	page2 := storyPage{"you climb into the attic, it is pitch black, you can't see a thing!", nil}
	page3 := storyPage{"you are eaten by a Grue", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	page1.playStory()

}
