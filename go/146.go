type LRUCache struct {
    keys []int
    values []int
    marks []int
    recent int
}


func Constructor(capacity int) LRUCache {
    result := LRUCache{make([]int, capacity), make([]int, capacity), make([]int, capacity), 0}
    for i := 0; i < capacity; i++ {
        result.keys[i] = -1
        result.values[i] = -1
    }
    return result
}


func (this *LRUCache) Get(key int) int {
    for i := 0; i < len(this.keys); i++ {
        if key == this.keys[i] {
            this.recent = this.recent+1
            this.marks[i] = this.recent
            return this.values[i]
        }
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    /* the key exists */
    for i := 0; i < len(this.keys); i++ {
        if this.keys[i] == key {
            this.values[i] = value
            this.recent = this.recent+1
            this.marks[i] = this.recent
            return
        }
    }
    /* has free space */
    for i := 0; i < len(this.keys); i++ {
        if this.marks[i] == 0 {
            this.keys[i] = key
            this.values[i] = value
            this.recent = this.recent+1
            this.marks[i] = this.recent
            return
        }
    }
    /* has no free space */
    smallest_mark_index := 0
    for i := 0; i < len(this.keys); i++ {
        if this.marks[i] < this.marks[smallest_mark_index] {
            smallest_mark_index = i
        }
    }
    this.keys[smallest_mark_index] = key
    this.values[smallest_mark_index] = value
    this.recent = this.recent+1
    this.marks[smallest_mark_index] = this.recent
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
