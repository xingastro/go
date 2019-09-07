package dictionary

import "sync"

type ValueDictionary struct {
	items map[string]string
	lock sync.RWMutex
}

func (d *ValueDictionary) Set(k, v string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.items == nil {
		d.items = make(map[string]string)
	}
	d.items[k] = v
}

func (d *ValueDictionary) Delete(k string) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	_, ok := d.items[k]
	if ok {
		delete(d.items, k)
	}
	return ok
}

func (d *ValueDictionary) Has(k string) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	_, ok := d.items[k]
	return ok
}

func (d *ValueDictionary) Get(k string) string {
	d.lock.Lock()
	defer d.lock.Unlock()

	return d.items[k]
}

func (d *ValueDictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.items = make(map[string]string)
}

func (d *ValueDictionary) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return len(d.items)
}

func (d *ValueDictionary) Keys() []string {
	d.lock.RLock()
	defer d.lock.RUnlock()

	keys := []string{}
	for i := range d.items {
		keys = append(keys, i)
	}
	return keys
}

func (d *ValueDictionary) Values() []string {
	d.lock.RLock()
	defer d.lock.RUnlock()

	values := []string{}
	for i := range d.items {
		values = append(values, d.items[i])
	}
	return values
}