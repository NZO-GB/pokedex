package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/NZO-GB/pokedex/internal/cache"
	"github.com/NZO-GB/pokedex/internal/pokeapi"
)


type configStruct struct {
	mapCall mapCall
	cache pokecache.Cache
	pokeClient *pokeapi.Client
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *configStruct) {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading input:", err)
				} else {
					fmt.Println("EOF, exiting REPL")
				}
				return
			}

			words := cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}

			command := words[0]
			var argument string
			if len(words) == 2 {
				argument = words[1]
			} else if len(words) > 2 {
				fmt.Println("Too many arguments")
				continue
			}


			cmd, exists := supportedCommands[command]
			if exists {
				fmt.Printf("Executing %s \n", command)
				if err := cmd.callback(cfg, argument); err != nil{
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
}