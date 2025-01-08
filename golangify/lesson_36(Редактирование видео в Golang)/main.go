// package main

// import (
// 	"log"

// 	"github.com/mowshon/moviego"
// )

// // go get github.com/mowshon/moviego
// // Пример кода на Golang как менять ширину и высоту видео
// func task_1() {
// 	first, _ := moviego.Load("video.mp4")

// 	first.ResizeByWidth(500).Output("resize-by-width.mp4").Run()
// 	first.ResizeByHeight(500).Output("resize-by-height.mp4").Run()
// 	first.Resize(1000, 500).Output("resize.mp4").Run()
// }

// // Обрезать видео на фрагменты в Golang
// func task_2() {
// 	first, _ := moviego.Load("video.mp4")

// 	// Обрезаем видео
// 	err := first.SubClip(350, 363).Output("subclip.mp4").Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // Объединение нескольких видео в одно в Golang
// func task_3() {
// 	first, _ := moviego.Load("subclip.mp4")
// 	second, _ := moviego.Load("my_vidio.mp4")

// 	finalVidio, err := moviego.Concat([]moviego.Video{
// 		first,                         // Первое видео без эффектов.
// 		second,                        // Второе видео тоже без эффектов.
// 		first.SubClip(0, 10),          // Из первого видео обрезаем
// 		second.SubClip(0, 5),          // Из второго видео обрезаем
// 		first.FadeIn(0, 5).FadeOut(5), // Для оригинального первого видео добавляем несколько эффектов.
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	renderErr := finalVidio.Output("concat.mp4").Run()
// 	if renderErr != nil {
// 		log.Fatal(renderErr)
// 	}
// }

// // Добавление переходов Fade-in, Fade-out для видео
// func task_4() {
// 	first, _ := moviego.Load("subclip.mp4")

// 	// Добавляем оба эффекта Fade-in и Fade-out.
// 	first.FadeIn(0, 3).FadeOut(5).Output("fade-in-with-fade-out.mp4").Run()

// 	// Обрезаем видео и добавляем эффект Fade-in.
// 	first.SubClip(5.20, 10).FadeIn(0, 3).Output("cut-fade-in.mp4").Run()

// 	// Полностью убираем звук до 0.5 секунды видео
// 	// и постепенно возвращаем его к секунде 4.
// 	first.AudioFadeIn(0.5, 4).Output("audio-fade-in.mp4").Run()

// 	// Угасание видео и аудио в самом конце за последние 5 секунд видео.
// 	first.FadeOut(5).AudioFadeOut(5).Output("fade-out.mp4").Run()
// }

// // Создание скриншота из видео в Golang
// func task_5() {
// 	first, _ := moviego.Load("forest.mp4")

// 	// Создаем обычный скриншот.
// 	first.Screenshot(2, "simple-screen.png")

// 	// Применяем эффект FadeIn и FadeOut и скриншотим момент применения фильтра.
// 	first.FadeIn(0, 6).FadeOut(5).Screenshot(2, "screen.png")
// }

// func test() {
// 	first, _ := moviego.Load("subclip.mp4")

// 	first.FadeIn(0, 3).FadeOut(2).AudioFadeIn(0, 4).AudioFadeOut(3).Output("finaly.mp4").Run()
// }

// func main() {
// 	// task_1() // Пример кода на Golang как менять ширину и высоту видео
// 	// task_2() // Обрезать видео на фрагменты в Golang
// 	// task_3() // Объединение нескольких видео в одно в Golang
// 	// task_4() // Добавление переходов Fade-in, Fade-out для видео
// 	// test()
// 	task_5() // Создание скриншота из видео в Golang
// }
