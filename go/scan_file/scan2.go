package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//使用import语句导入需要使用的外部包，
	//分别是fmt、io/ioutil、log、os和path/filepath，
	//分别用于输出信息、读取文件信息、输出日志信息、读取系统信息和处理文件路径信息
)

/*定义一个printFileStats函数，输入路径参数path，
  递归遍历该路径下所有的文件和子目录，打印每个文件的名称和大小等信息*/
func printFileStats(path string) {
	/*使用os.Stat获取指定路径的文件信息，如果出现错误（如文件不存在），
	使用log.Printf输出错误信息*/
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Failed to read file info for %s: %v", path, err)
	}
	//如果该路径是一个目录，使用ioutil.ReadDir读取该目录下所有文件的信息
	if fileInfo.IsDir() {
		dirEntries, err := ioutil.ReadDir(path)
		if err != nil {
			log.Printf("Failed to read directory entries for %s: %v", path, err)
		}

		var dirCount, fileCount int
		var dirSize, fileSize int64
		/*使用循环处理该目录下的每个文件或子目录，对文件使用printFileStats函数
		  进行递归处理并统计目录下的文件数量、大小等信息*/
		for _, entry := range dirEntries {
			printFileStats(filepath.Join(path, entry.Name()))

			if entry.IsDir() {
				dirCount++
				dirSize += entry.Size()
			} else {
				fileCount++
				fileSize += entry.Size()
			}
		}
		//最后使用fmt.Printf输出目录的信息
		fmt.Printf("Directory %s contains %d subdirectories, %d files (total size %d bytes)\n", path, dirCount, fileCount, dirSize+fileSize)
	} else {
		fmt.Printf("File %s has size %d bytes\n", path, fileInfo.Size())
	}
}

func main() {
	var path string
	/*在main函数中，读取用户输入的命令行参数指定的路径（如果没有输入参数，
	则默认为当前目录），并将其传递给printFileStats函数进行处理*/
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "." // default to current directory
	}

	printFileStats(path)
}
