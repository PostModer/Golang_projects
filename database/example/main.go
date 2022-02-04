package main

import (
	"bufio"
	"fmt"
	"kvdb/example/consts"
	"kvdb/example/database"
	"os"
)

func main() {
	base := database.NewDatabase()

	if err := base.Load(); err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		command, key, value := Handler(text)

		switch command {
		case consts.Create:
			fmt.Println(base.Create(key, value))

		case consts.Read:
			fmt.Println(base.Read(key))

		case consts.Update:
			fmt.Println(base.Update(key, value))

		case consts.Delete:
			fmt.Println(base.Delete(key))

		case consts.Exist:
			fmt.Println(base.Exist(key))

		case consts.Quit:
			fmt.Println(base.Quit())

		case consts.Sum:
			fmt.Println(base.Sum())

		case consts.Avg:
			fmt.Println(base.Avg())

		case consts.Gt:
			fmt.Println(base.Gt(key))

		case consts.Lt:
			fmt.Println(base.Lt(key))

		case consts.Eq:
			fmt.Println(base.Eq(key))

		case consts.Count:
			fmt.Println(base.Count())

		default:
			fmt.Println(consts.Invalid)
		}

		base.Save()
	}
}
