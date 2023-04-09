package handlers

import (
	"container/list"
	"errors"
)

var cl *list.List

type Clip struct {
}

func (c Clip) AddTop(r *string, reply *string) (err error) {

	cl.PushFront(*r)

	return
}

func (c Clip) AddBottom(r *string, reply *string) (err error) {

	cl.PushBack(*r)
	return
}

func (c Clip) GetTop(r *string, reply *string) (err error) {
	if cl.Len() <= 0 {
		err = errors.New("clipboard empty")
		return
	}
	data, ok := cl.Front().Value.(string)
	if !ok {
		err = errors.New("error : type casting error")
	}
	*reply = data
	return
}
func (c Clip) GetBottom(r *string, reply *string) (err error) {
	if cl.Len() <= 0 {
		err = errors.New("clipboard empty")
		return
	}
	data, ok := cl.Back().Value.(string)
	if !ok {
		err = errors.New("error : type casting error")
	}
	*reply = data

	return
}

func (c Clip) RemoveTop(r *string, reply *string) (err error) {

	if cl.Len() <= 0 {
		err = errors.New("clipboard empty")
		return
	}
	cl.Remove(cl.Front())
	return
}

func (c Clip) RemoveBottom(r *string, reply *string) (err error) {

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
