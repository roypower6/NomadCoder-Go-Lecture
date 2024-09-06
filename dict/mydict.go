package mydict

import (
	"errors"
	"fmt"
)

func main() {
	dictionary := Dictionary{"first": "First word"}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Update(baseword, "Second")
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}

type Dictionary map[string]string //사전 생성(이름이 Dictionary)

var (
	errNotFound    = errors.New("Not Found")                      //검색 시 단어 없음
	errWordExists  = errors.New("That word already exists")       //추가하려는 단어가 이미 존재함
	errCantUpdate  = errors.New("Can't update non-existing word") //업데이트하려는 단어가 없음
	errDoesntExist = errors.New("That word doesn't exist")        //삭제하려는 단어가 존재하지 않음
)

// 사전에서 검색
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// 사전 단어 추가
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// 사전 기존 단어 업데이트
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
	} else {
		return errCantUpdate
	}
	return nil
}

// 사전 단어 삭제
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
		return nil
	}
	return errDoesntExist
}
