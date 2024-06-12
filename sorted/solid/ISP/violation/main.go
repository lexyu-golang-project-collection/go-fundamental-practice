package main

import "fmt"

type Content struct {
	Title string
	Body  string
}

type ContentManager interface {
	Edit(content *Content)
	Publish(content *Content)
	Play(content *Content)
}

type ArticleManager struct{}

func (a *ArticleManager) Edit(content *Content) {
	fmt.Println("Article Editing...")
	content.Title = "Greeing"
	content.Body = "Some Description"
}

func (a *ArticleManager) Publish(content *Content) {
	fmt.Println("Article Publishing...")
}

// ArticleManager implement Play method, which is irrelevant for articles.
func (a *ArticleManager) Play(content *Content) {
	fmt.Println("Article Doesn't Support Playing...")
}

func main() {
	c := &Content{
		Title: "Title",
		Body:  "Body",
	}
	a := &ArticleManager{}

	a.Edit(c)
	a.Publish(c)
	a.Play(c)

	fmt.Printf("%+v", c)
}
