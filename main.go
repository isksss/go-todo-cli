package main

// cliでToDoを管理するアプリケーション
// 1. ToDoを追加する
// 2. ToDoを完了にする
// 3. ToDoを削除する
// 4. ToDoを一覧表示する
// 5. ToDoを完了済みのみ表示する
// 6. ToDoを未完了のみ表示する
// 7. ToDoを完了済みにする
// 8. ToDoを未完了にする
// 9. ToDoを編集する
// 10. ToDoを検索する
// 11. ToDoをソートする
// 12. ToDoをファイルに保存する
// 13. ToDoをファイルから読み込む

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ToDoの構造体
// jsonタグを付けることで、json形式に変換できる

type ToDo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Done   bool   `json:"done"`
	// doneDate time.Time `json:"doneDate"`
}

// ToDoのスライス
var toDoList []ToDo

func main() {
	// メニューを表示する
	showMenu()
}

// メニューを表示する
func showMenu() {
	fmt.Println("ToDo管理アプリ----------")
	fmt.Println("1. ToDoを追加する")
	fmt.Println("2. ToDoを完了にする")
	fmt.Println("3. ToDoを削除する")
	fmt.Println("4. ToDoを一覧表示する")
	fmt.Println("5. ToDoを完了済みのみ表示する")
	fmt.Println("6. ToDoを未完了のみ表示する")
	fmt.Println("7. ToDoを完了済みにする")
	fmt.Println("8. ToDoを未完了にする")
	fmt.Println("9. ToDoを編集する")
	fmt.Println("10. ToDoを検索する")
	fmt.Println("11. ToDoをソートする")
	fmt.Println("12. ToDoをファイルに保存する")
	fmt.Println("13. ToDoをファイルから読み込む")
	fmt.Println("0. 終了する")
	fmt.Print("メニューを選択してください: ")

	// 入力を受け取る
	input := input()

	// 入力された値に応じて処理を分岐する
	switch input {
	case "1":
		addToDo()
	case "2":
		doneToDo()
	case "3":
		deleteToDo()
	case "4":
		showToDoList()
	case "5":
		showDoneToDoList()
	case "6":
		showUndoneToDoList()
	case "7":
		doneAllToDo()
	case "8":
		undoneAllToDo()
	case "9":
		editToDo()
	case "10":
		searchToDo()
	case "11":
		sortToDo()
	case "12":
		saveToDo()
	case "13":
		loadToDo()
	case "0":
		fmt.Println("アプリケーションを終了します")
		os.Exit(0)
	default:
		fmt.Println("入力が正しくありません")
	}

	showMenu()
}

// ToDoを追加する
func addToDo() {
	// ToDoのタイトルを入力する
	fmt.Print("ToDoのタイトルを入力してください: ")
	title := input()

	// ToDoの詳細を入力する
	fmt.Print("ToDoの詳細を入力してください: ")
	detail := input()

	// ToDoの構造体を作成する
	toDo := ToDo{
		Id:     len(toDoList) + 1,
		Title:  title,
		Detail: detail,
	}

	// ToDoのスライスに追加する
	toDoList = append(toDoList, toDo)
}

// ToDoを完了にする
func doneToDo() {
	// ToDoのIDを入力する
	fmt.Print("完了にするToDoのIDを入力してください: ")
	id, err := strconv.Atoi(input())
	if err != nil {
		fmt.Println("入力が正しくありません")
		return
	}

	// ToDoのスライスからIDに一致するToDoを探す
	for i, toDo := range toDoList {
		if toDo.Id == id {
			// ToDoを完了にする
			toDoList[i].Done = true
			// toDoList[i].doneDate = time.Now()
			return
		}
	}

	fmt.Println("入力が正しくありません")
}

// ToDoを削除する
func deleteToDo() {
	// ToDoのIDを入力する
	fmt.Print("削除するToDoのIDを入力してください: ")
	id, err := strconv.Atoi(input())
	if err != nil {
		fmt.Println("入力が正しくありません")
		return
	}

	// ToDoのスライスからIDに一致するToDoを探す
	for i, toDo := range toDoList {
		if toDo.Id == id {
			// ToDoを削除する
			toDoList = append(toDoList[:i], toDoList[i+1:]...)
			return
		}
	}

	fmt.Println("入力が正しくありません")
}

// ToDoを一覧表示する
func showToDoList() {
	// ToDoのスライスをループして表示する
	for _, toDo := range toDoList {
		fmt.Println(toDo)
	}
}

// ToDoを完了済みのみ表示する
func showDoneToDoList() {
	// ToDoのスライスをループして表示する
	for _, toDo := range toDoList {
		if toDo.Done {
			fmt.Println(toDo)
		}
	}
}

// ToDoを未完了のみ表示する
func showUndoneToDoList() {
	// ToDoのスライスをループして表示する
	for _, toDo := range toDoList {
		if !toDo.Done {
			fmt.Println(toDo)
		}
	}
}

// ToDoを完了済みにする
func doneAllToDo() {
	// ToDoのスライスをループして完了にする
	for i := range toDoList {
		toDoList[i].Done = true
	}
}

// ToDoを未完了にする
func undoneAllToDo() {
	// ToDoのスライスをループして未完了にする
	for i := range toDoList {
		toDoList[i].Done = false
	}
}

// ToDoを編集する
func editToDo() {
	// ToDoのIDを入力する
	fmt.Print("編集するToDoのIDを入力してください: ")
	id, err := strconv.Atoi(input())
	if err != nil {
		fmt.Println("入力が正しくありません")
		return
	}

	// ToDoのスライスからIDに一致するToDoを探す
	for i, toDo := range toDoList {
		if toDo.Id == id {
			// ToDoのタイトルを入力する
			fmt.Print("ToDoのタイトルを入力してください: ")
			title := input()

			// ToDoの詳細を入力する
			fmt.Print("ToDoの詳細を入力してください: ")
			detail := input()

			// ToDoを編集する
			toDoList[i].Title = title
			toDoList[i].Detail = detail
			return
		}
	}

	fmt.Println("入力が正しくありません")
}

// ToDoを検索する
func searchToDo() {
	// ToDoのタイトルを入力する
	fmt.Print("検索するToDoのタイトルを入力してください: ")
	title := input()

	// ToDoのスライスをループして表示する
	for _, toDo := range toDoList {
		if strings.Contains(toDo.Title, title) {
			fmt.Println(toDo)
		}
	}
}

// ToDoをソートする
func sortToDo() {
	// ToDoのスライスをソートする
	sort.Slice(toDoList, func(i, j int) bool {
		return toDoList[i].Id < toDoList[j].Id
	})
}

// ToDoを保存する
func saveToDo() {
	// ToDoのスライスをJSONに変換する
	file, err := json.Marshal(toDoList)
	if err != nil {
		fmt.Println("ToDoの保存に失敗しました")
		return
	}

	// ToDoのJSONをファイルに保存する
	err = ioutil.WriteFile("todo.json", file, 0644)
	if err != nil {
		fmt.Println("ToDoの保存に失敗しました")
		return
	}

	fmt.Println("ToDoを保存しました")
}

// ToDoを読み込む
func loadToDo() {
	// ToDoのJSONをファイルから読み込む
	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		fmt.Println("ToDoの読み込みに失敗しました")
		return
	}

	// ToDoのJSONをスライスに変換する
	err = json.Unmarshal(file, &toDoList)
	if err != nil {
		fmt.Println("ToDoの読み込みに失敗しました")
		return
	}

	fmt.Println("ToDoを読み込みました")
}

// ToDoを削除する
func deleteAllToDo() {
	// ToDoのスライスを空にする
	toDoList = []ToDo{}
}

// ToDoのIDを生成する
func generateToDoID() int {
	// ToDoのスライスの長さをIDにする
	return len(toDoList)
}

// ToDoのタイトルを入力する
func inputToDoTitle() string {
	fmt.Print("ToDoのタイトルを入力してください: ")
	return input()
}

// ToDoの詳細を入力する
func inputToDoDetail() string {
	fmt.Print("ToDoの詳細を入力してください: ")
	return input()
}

// 入力を受け付ける
func input() string {
	var input string
	fmt.Scan(&input)
	return input
}

// ToDoの文字列を返す
func (toDo ToDo) String() string {
	return fmt.Sprintf("%d: %s %s %t", toDo.Id, toDo.Title, toDo.Detail, toDo.Done)
}
