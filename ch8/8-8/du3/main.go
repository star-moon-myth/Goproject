package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messsage")

func main() {
	flag.Parse()
	roots:=flag.Args()
	if len(roots)==0{
		roots=[]string{"."}
	}
	fileSizes:=make(chan int64)
	var n sync.WaitGroup
	for _,root:=range roots{
		n.Add(1)
		go walkDir(root,&n,fileSizes)
	}
	go func ()  {
		n.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose{
		tick=time.Tick(100*time.Second)
	}
	var nfiles,nbytes int64
loop:
	for{
		select {
		case size,ok:=<-fileSizes:
			if !ok{
				break loop
			}
			nfiles++
			nbytes+=size
		case <-tick:
			printDiskUsage(nfiles,nbytes)
		}

	}

	printDiskUsage(nfiles,nbytes)
}
func walkDir(dir string,n *sync.WaitGroup,fileSizes chan<-int64){
	defer n.Done()
	for _,entry:=range dirents(dir){
		if entry.IsDir(){
			n.Add(1)
			subdir:=filepath.Join(dir,entry.Name())
			walkDir(subdir,n,fileSizes)
		}
	}
}
var sema = make(chan struct{}, 20) 
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}