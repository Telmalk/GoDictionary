package main

import (
	"flag"
	"fmt"
	"os"
	"training.go/Dictionary/dictionary"
)

func actionList(d *dictionary.Dictionary)  {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionary contains")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleErr(err error)  {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}

func actionAdd(d *dictionary.Dictionary, args []string)  {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the dictionary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string)  {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}

func actionDelete(d *dictionary.Dictionary, args []string)  {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' removed from the dictionary\n", word)
}

func main()  {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "delete":
		actionDelete(d, flag.Args())

	default:
		fmt.Printf("Unknow action: %v", *action)
	}
}
