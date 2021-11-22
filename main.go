package main

import (
	"context"
	"fmt"
	"strconv"
	"todolist/model"
	"todolist/repository"
	"todolist/service"
	"todolist/utils"
)

func main() {
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())
	todolistService := service.NewTodolistService(todolistRepository)

	fmt.Print("Program Sederhana Todolist" +
		"\n==========================")

	for {
		fmt.Println("\nOption :" +
			"\n1. Lihat Todolist" +
			"\n2. Buat Todolist" +
			"\n3. Hapus Todolist" +
			"\n4. Keluar")

		fmt.Print("Option : ")
		utils.Input.Scan()

		input := utils.Input.Text()

		if input == "1" {
			ctx := context.Background()
			todolistService.ShowTodolist(ctx)
		} else if input == "2" {
			var todolist model.Todolist
			ctx := context.Background()

			fmt.Println("\nBuat Todolist :  ")

			fmt.Print("Todolist : ")
			utils.Input.Scan()
			todolist.Name = utils.Input.Text()

			fmt.Print("Author : ")
			utils.Input.Scan()
			todolist.Author = utils.Input.Text()

			todolistService.CreateTodolist(ctx, &todolist)
		} else if input == "3" {
			ctx := context.Background()

			fmt.Print("\nHapus : ")
			utils.Input.Scan()

			input = utils.Input.Text()

			index, _ := strconv.Atoi(input)

			todolistService.DeleteTodolist(ctx, &index)
		} else if input == "4" {
			break
		} else {
			fmt.Println("Mohon masukkan angka 1 - 4")
		}
	}

	fmt.Println("Terima Kasih :3")
}
