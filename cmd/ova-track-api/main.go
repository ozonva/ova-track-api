package main

import (
	"fmt"
	"github.com/ozonva/ova-track-api/internal/utils"

	//	"github.com/golang/mock/gomock"
//	"github.com/ozonva/ova-track-api/internal/utils"
//	"github.com/ozonva/ova-track-api/internal/flusher"
//	"github.com/ozonva/ova-track-api/internal/mocks"
	"os"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
)

func main() {
	fmt.Println("Hi, i am ova-track-api!")
	if len (os.Args) != 2{
		fmt.Println("Path to config is strictly required")
		return
	}
	path := os.Args[1]
	utils.InitLibraryFromFile(path)


}
