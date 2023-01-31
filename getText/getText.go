package main

import (
	"log"
)

func main() {
	inputText := "想知道如果提高王者荣耀水平,上分把妹不是梦,加微信:17252sugats78,加QQ:34676328889"
	inputTags := "0000000000000000000000012034444444440012034444444444"
	cnt := 0
	for _, c := range inputTags {
		if c != '0' {
			cnt++
		}

		//fmt.Println(fmt.Sprintf("text:[%+v] , tags:[%+v]", string(inputText[i]), string(inputTags[i])))
	}
	log.Println(cnt)
	log.Println(inputText[cnt : cnt+2])
}
func getText(inputText, inputTags string) (keyWords, tips []string) {
	keyWords = make([]string, 0)
	tips = make([]string, 0)
	left := 0
	for idx, c := range inputText {
		if idx > len(inputTags) {
			break
		}
		switch c {
		case '1':
			left = idx
		case '2':
			if inputText[idx-1] != '2' || idx == len(inputText)-1 {
				keyWords = append(keyWords, inputText[left:idx])
			}
		case '3':
			left = idx
		case '4':
			if inputText[idx-1] != '4' || idx == len(inputText)-1 {
				tips = append(tips, inputText[left:idx])
			}
		}
	}
	return
}
