package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func repetition(st string) map[string]int {
	st = strings.ToUpper(st)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	st1 := reg.ReplaceAllString(st, " ")
	// using strings.Field Function
	input := strings.Fields(st1)
	fmt.Println(input)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	keys := make([]int, 0, len(wc))
	for k := range wc {
		keys = append(keys, wc[k])
	}
	sort.Ints(keys)
	fmt.Println(keys)
	unque_slice := unique(keys)
	fmt.Println(unque_slice)
	if len(unque_slice) < 11 {
		return wc
	} else {
		fina_wc := make(map[string]int)
		for _, i := range unque_slice[len(unque_slice)-10:] {
			for k, v := range wc {
				if v == i {
					fina_wc[k] = v
				}
			}
		}
		return fina_wc
	}
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	router := gin.Default()

	router.POST("/text", func(c *gin.Context) {

		p, _ := ioutil.ReadAll(c.Request.Body)
		println(string(p))
		err := c.Bind(&p)
		if err != nil {
			log.Fatalln(err)
		}

		aval := repetition(string(p))
		fmt.Println(aval)

		c.JSON(http.StatusOK, gin.H{
			"Result": aval,
		})

	})

	router.Run(":8000")
}
