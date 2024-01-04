package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task struct'u, görevlerin özelliklerini temsil eder
type Task struct {
	id       int
	title    string
	completed bool
}

// TodoList struct'ı, görevleri saklamak için kullanılır
type TodoList struct {
	tasks []Task
}

// Yeni bir görev oluşturur ve TodoList'e ekler
func (list *TodoList) addTask(title string) {
	newID := len(list.tasks) + 1
	newTask := Task{id: newID, title: title, completed: false}
	list.tasks = append(list.tasks, newTask)
	fmt.Printf("'%s' görevi eklendi!\n", title)
}

// Görevleri listeler
func (list TodoList) listTasks() {
	fmt.Println("Yapılacaklar Listesi:")
	for _, task := range list.tasks {
		status := "Tamamlanmadı"
		if task.completed {
			status = "Tamamlandı"
		}
		fmt.Printf("%d. [%s] %s\n", task.id, status, task.title)
	}
}

// Görevi tamamlanmış olarak işaretler
func (list *TodoList) completeTask(id int) {
	for i := range list.tasks {
		if list.tasks[i].id == id {
			list.tasks[i].completed = true
			fmt.Printf("'%s' görevi tamamlandı!\n", list.tasks[i].title)
			return
		}
	}
	fmt.Println("Görev bulunamadı!")
}

func main() {
	myList := TodoList{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Yapılacaklar Listesi Uygulamasına Hoş Geldiniz!")
	fmt.Println("Komutlar: 'ekle', 'liste', 'tamamla', 'çıkış'")

	for {
		fmt.Print("Komut: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "ekle":
			fmt.Print("Yeni görev: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			myList.addTask(task)
		case "liste":
			myList.listTasks()
		case "tamamla":
			fmt.Print("Tamamlanan görevin ID'si: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)
			id, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Geçersiz ID!")
				continue
			}
			myList.completeTask(id)
		case "çıkış":
			fmt.Println("Yapılacaklar Listesi Uygulaması Sonlandırıldı.")
			return
		default:
			fmt.Println("Geçersiz komut! Lütfen tekrar deneyin.")
		}
	}
}
