package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	payloadParseError = "payload parse error"
	jsonParseError    = "json parse error"
	writeBufferError  = "write buffer error"
)

type InputObject struct {
	Array []interface{} `json:"Array"`
}
type Node struct {
	Val  interface{}
	Next *Node
}

func (n *Node) ToString() string {
	result := ""
	current := n
	count := 0
	for current != nil {
		if count == 0 {
			result += fmt.Sprintf("head: %v\n", current.Val)
		} else if current.Next != nil {
			result += fmt.Sprintf("node%d: %v\n", count, current.Val)
		} else {
			result += fmt.Sprintf("tail: %v\n", current.Val)
		}
		count++
		current = current.Next
	}
	return result
}
func (n *Node) ShowValue() {
	current := n
	count := 0
	for current != nil {
		if count == 0 {
			fmt.Printf("head: %v\n", current.Val)
		} else if current.Next != nil {
			fmt.Printf("node%d: %v\n", count, current.Val)
		} else {
			fmt.Printf("tail: %v\n", current.Val)
		}
		count++
		current = current.Next
	}
}
func Test1(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Test 1:")
	// parse request body into linked list
	// 1 parse request body into object
	var object InputObject
	err := json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		http.Error(w, payloadParseError, http.StatusBadRequest)
		return
	}
	// 2 parse array into linkedList
	linkedList := ParseArrayIntoLinkedList(object.Array)
	linkedList.ShowValue()
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(&linkedList)
	// err = json.NewEncoder(w).Encode(&linkedList)
	if err != nil {
		http.Error(w, jsonParseError, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, writeBufferError, http.StatusInternalServerError)
	}
}

func ParseArrayIntoLinkedList(arr []interface{}) *Node {
	var head *Node
	var prev *Node
	for _, val := range arr {
		current := &Node{
			Val: val,
		}
		if prev == nil {
			prev = current
		} else {
			prev.Next = current
			prev = current
		}
		if head == nil {
			head = current
		}
	}
	return head
}
