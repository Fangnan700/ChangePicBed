package main

import (
	"ChangePicBed/model"
	"ChangePicBed/utils"
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
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Reading input files...")
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

	fmt.Printf("\nDo you want to continue?(y/n)")
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadRune()
	if input != 121 {
		os.Exit(0)
	}
}

func main() {
	startTamp := time.Now().UnixNano()
	totalFiles := 0
	successFiles := 0
	failedFiles := 0
	fmt.Println("Processing files...")
	for _, inputFile := range inputFiles {
		totalFiles += 1
		err = utils.ChangePicBed(inputFile, config)
		if err != nil {
			fmt.Println(err)
			failedFiles += 1
		}
		successFiles += 1
	}

	err = utils.Clear(config.TempDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	endTamp := time.Now().UnixNano()
	duration := float64(endTamp-startTamp) / 1000000000
	fmt.Printf("Total time spent: %.2f seconds\n", duration)
	fmt.Println("Total files:", totalFiles)
	fmt.Println("Success files:", successFiles)
	fmt.Println("Failed files:", failedFiles)
	fmt.Println("Press any key to exit...")
	reader := bufio.NewReader(os.Stdin)
	_, _, err = reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}
