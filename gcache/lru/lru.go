package lru

import "container/list"

type Cache struct{
	maxBytes int64
	nBytes int64
	ll *list.List
	cahce map[string] *list.Element
	OnEvicted func(key string, value Value)
}

type entry struct{
	key string
	value Value
}

type Value interface{
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache{
	return &Cache{
		maxBytes:maxBytes ,
		ll: list.New(),
		cahce: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache)Get(key string)(value Value,ok bool){
	if ele,ok := c.cahce[key];ok{
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value,true
	}
	return
}

func (c *Cache) RemoveOldest(){
	ele:=c.ll.Back()
	if ele !=nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cahce,kv.key)
		c.nBytes -= (int64(len(kv.key))+ int64(kv.value.Len()))
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key,kv.value)
		}
	}
}

func (c *Cache)Add(key string,value Value){
	if ele,ok :=c.cahce[key];ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes +=int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	}else{
		ele := c.ll.PushFront(&entry{key,value})
		c.cahce[key] = ele
		c.nBytes +=int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes !=0 && c.maxBytes <c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}