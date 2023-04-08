package handlers

import (
	"container/list"
	"errors"
	"log"
)

var cl *list.List

type Clip struct {
}

func (c Clip) AddTop(r *string, reply *string) (err error) {

	log.Println("add to top :", *r)

	cl.PushFront(*r)

	return
}

func (c Clip) AddBottom(r *string, reply *string) (err error) {

	log.Println("add to bottom :", *r)

	cl.PushBack(*r)
	return
}

func (c Clip) GetTop(r *string, reply *string) (err error) {
	data, ok := cl.Front().Value.(string)
	if !ok {
		err = errors.New("error : type casting error")
	}
	*reply = data
	log.Println("get to top :", data)
	return
}
func (c Clip) GetBottom(r *string, reply *string) (err error) {
	data, ok := cl.Back().Value.(string)
	if !ok {
		err = errors.New("error : type casting error")
	}
	*reply = data
	log.Println("get to bottom :", data)

	return
}

func (c Clip) RemoveTop(r *string, reply *string) (err error) {
	log.Println("remove to top")

	if cl.Len() <= 0 {
		err = errors.New("clipboard empty")
		return
	}
	cl.Remove(cl.Front())
	return
}

func (c Clip) RemoveBottom(r *string, reply *string) (err error) {
	log.Println("remove to bottom")

	if cl.Len() <= 0 {
		err = errors.New("clipboard empty")
		return
	}
	cl.Remove(cl.Back())
	return
}

func init() {
	cl = list.New()
}
