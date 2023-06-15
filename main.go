package main

import (
	"ChangePicBed/model"
	"ChangePicBed/utils"
	"ChangePicBed/yuque"
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	config     model.Config
	inputFiles []string
	reader     *bufio.Reader
	err        error
)

func init() {

	fmt.Println("Loading configuration file...")
	configPath := "config"
	err = utils.CheckDir(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = utils.PrintConfig(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	inputPath := config.InputDir
	err = utils.CheckDir(inputPath)
	outputPath := config.OutputDir
	err = utils.CheckDir(outputPath)
	tempPath := config.TempDir
	err = utils.CheckDir(tempPath)
	yuqueExportPath := config.YuqueConfig.ExportPath
	err = utils.CheckDir(yuqueExportPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("\nDo you want to continue?(y/n)")
	reader = bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadRune()
	if input != 121 {
		os.Exit(0)
	}
}

func changePicBed() {
	fmt.Println("Reading input files...")
	inputPath := config.InputDir
	count := 0
	err = filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		if !info.IsDir() {
			infoArr := strings.Split(info.Name(), ".")
			pattern := infoArr[len(infoArr)-1]
			if pattern == "md" {
				count += 1
				fmt.Println("Found file:", count, "==>", info.Name())
				inputFiles = append(inputFiles, path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	startTamp := time.Now().UnixNano()
	fmt.Println("Processing files...")
	for _, inputFile := range inputFiles {
		err = utils.ChangePicBed(inputFile, config)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = utils.Clear(config.TempDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	endTamp := time.Now().UnixNano()
	duration := float64(endTamp-startTamp) / 1000000000
	fmt.Printf("Total time spent: %.2f seconds\n", duration)

	return
}

func getYuqueBookStacks(options int) {
	var bookStacks model.YuqueBookStacks
	bookStacks, err = yuque.GetBookStacks(config)

	if len(bookStacks.Data) <= 0 {
		fmt.Println("Knowledge base not found.")
		return
	}

	fmt.Println("The Yuque knowledge base data is as follows:")
	for _, data := range bookStacks.Data {
		fmt.Printf("-%s\n", data.Name)
		for _, book := range data.Books {
			fmt.Printf("|----[BookID: %d] %s\n", book.Id, book.Name)
			for _, summary := range book.Summary {
				fmt.Printf("\t|----[SummaryID:%d] %s\n", summary.Id, summary.Title)
			}
		}
	}

	fmt.Println("\nPlease ensure that the relevant configuration data is correctly filled in the configuration file before exporting Yuque notes.")
	fmt.Printf("\nDo you want to expor all books?(y/n)")

	reader.Reset(os.Stdin)
	input, _, _ := reader.ReadRune()
	if input == 121 {
		for _, data := range bookStacks.Data {
			for _, book := range data.Books {
				yuque.ExportBook(bookStacks, book.Id, config, options)
			}
		}
	}
}

func main() {
	for {
		fmt.Println("\nThe function list is as follows:")
		fmt.Printf("\t1、Export yuque books and changing the images bed\n")
		fmt.Printf("\t2、Export yuque books but without changing the images bed\n")
		fmt.Printf("\t3、Change markdown's images bed only\n")
		fmt.Printf("\t4、Exit")
		fmt.Printf("\nPlease enter the serial number corresponding to the above function to proceed:")

		reader = bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadRune()

		if input == 49 {
			getYuqueBookStacks(1)
		}
		if input == 50 {
			getYuqueBookStacks(0)
		}
		if input == 51 {
			changePicBed()
		}
		if input == 52 {
			_ = utils.Clear(config.TempDir)
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}
}
