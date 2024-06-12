package main

import "fmt"

type Content struct {
	Title string
	Body  string
}

// Segregated interfaces
type Editor interface {
	Edit(content *Content)
}

type Publisher interface {
	Publish(content *Content)
}

type Player interface {
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

type VideoManager struct{}

func (v *VideoManager) Edit(content *Content) {
	// Implement editing logic specific to videos
}

func (v *VideoManager) Publish(content *Content) {
	// Implement publishing logic for videos
}

func (v *VideoManager) Play(content *Content) {
	// Implement play functionality for videos
}

func main() {
	// ...
}
