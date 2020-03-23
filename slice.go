package gohelper

/*
结论： <=1024 , 按 2 倍扩容；> 1024 按 1.25 倍扩容


为什么？


type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度
    cap   int // 容量
}

func growslice(et *_type, old slice, cap int) slice {
    // ……
    newcap := old.cap
    doublecap := newcap + newcap
    if cap > doublecap {
        newcap = cap
    } else {
        if old.len < 1024 {
            newcap = doublecap
        } else {
            for newcap < cap {
                newcap += newcap / 4
            }
        }
    }
    // ……

    capmem = roundupsize(uintptr(newcap) * ptrSize)
    newcap = int(capmem / ptrSize)

}
*/

type Slice struct {
	v []interface{}
}

func NewSlice(len int, cap int) *Slice {
	return &Slice{v: make([]interface{},len,cap),
	}
}


func (s Slice) Len() int {
	return len(s.v)
}

func (s Slice) Cap() int{
	return cap(s.v)
}