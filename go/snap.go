package main

import (
	"fmt"
)

type Elem struct {
	value   int
	snap_id int
}

type Storage map[int][]Elem

type SnapshotArray struct {
	length_       int
	snapshots_    Storage
	current_snap_ map[int]int
	current_idx_  int
}

func Constructor(length int) SnapshotArray {
	snapshots := make(Storage)
	current_snap := make(map[int]int)
	return SnapshotArray{length_: length,
		snapshots_:    snapshots,
		current_snap_: current_snap,
		current_idx_:  -1}
}

func (this *SnapshotArray) Set(index int, val int) {
	if index < 0 || index >= this.length_ {
		return
	}
	this.current_snap_[index] = val
}

func (this *SnapshotArray) Snap() int {
	new_index := this.current_idx_ + 1
	for key, value := range this.current_snap_ {
		slice := this.snapshots_[key]
		this.snapshots_[key] = append(slice, Elem{value, new_index})
	}

	for key := range this.current_snap_ {
		delete(this.current_snap_, key)
	}

	this.current_idx_++
	return new_index
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if index < 0 || index >= this.length_ {
		return 0
	}
	if snap_id > this.current_idx_ {
		return 0
	}

	slice := this.snapshots_[index]
	for i := 0; i < len(slice); i++ {
		if slice[i].snap_id <= snap_id {
			return slice[i].value
		}
	}

	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

func main() {
	length := 10
	obj := Constructor(length)
	obj.Set(0, 15)
	obj.Snap()
	obj.Snap()
	obj.Snap()
	fmt.Println(obj.Get(0, 2))
	obj.Snap()
	obj.Snap()
	fmt.Println(obj.Get(0, 0))
}
