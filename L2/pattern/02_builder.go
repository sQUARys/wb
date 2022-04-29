package main

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Builder interface {
	makeName(str string)
	makeDescription(str string)
	makeConclusion(str string)
}

type Text struct {
	builder Builder
}

func (t *Text) Collect() {
	t.builder.makeName("Text about plants")
	t.builder.makeDescription("All plants are green")
	t.builder.makeConclusion("Plants has a green color")
}

type fieldsBuilder struct {
	content *Content
}

type Content struct {
	str string
}

func (fb *fieldsBuilder) makeName(str string) {
	fb.content.str += "<header>" + str + "</header>"
}

func (fb *fieldsBuilder) makeDescription(str string) {
	fb.content.str += "\n<desc>" + str + "</desc>"
}

func (fb *fieldsBuilder) makeConclusion(str string) {
	fb.content.str += "\n<concl>" + str + "</concl>"
}

//func main() {
//	content := new(Content)
//
//	text := Text{
//		&fieldsBuilder{
//			content,
//		},
//	}
//
//	text.Collect()
//
//	fmt.Println(content.str)
//
//}
