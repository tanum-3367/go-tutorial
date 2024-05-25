package author

import "fmt"

type Author struct {
	Name    string
	Contact string
}

func NewAuthor(name, contact string) *Author {
	return &Author{
		Name:    name,
		Contact: contact,
	}
}

func (a *Author) WriteChapter(chapterTitle string, content string) {
	fmt.Printf("Writing chapter %s with content: %s\n", chapterTitle, content)
}

func (a *Author) ReviewChapter(chapterTitle string, content string) {
	fmt.Printf("Reviewing chapter %s with content: %s\n", chapterTitle, content)
}

func (a *Author) FinalizeChapter(chapterTitle string) {
	fmt.Printf("Finalizing chapter %s\n", chapterTitle)
}

func main() {}
