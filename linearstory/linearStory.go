package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	page.nextPage.playStory()
}

func main() {
	page1 := storyPage{"ceci est la première page", nil}
	page2 := storyPage{"ceci est la deuxieme page", nil}
	page3 := storyPage{"ceci est la troisième page", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3
}
