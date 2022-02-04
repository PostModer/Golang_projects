package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

type Database struct {
	mutex *sync.RWMutex
	base  map[string]string
}

func NewDatabase() *Database {
	return &Database{
		mutex: new(sync.RWMutex),
		base:  make(map[string]string),
	}
}

func (o *Database) Create(key, value string) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.base[key] = value
	return o.base[key] == value
}

func (o *Database) Read(key string) (string, error) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	value, ok := o.base[key]
	if !ok {
		errMsg := fmt.Errorf("key %v not found", key)
		return "", errMsg
	}
	return value, nil
}

func (o *Database) Update(key, value string) (bool, error) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	_, ok := o.base[key]
	if !ok {
		errMsg := fmt.Errorf("key %v not found", key)
		return false, errMsg
	}
	o.base[key] = value
	return true, nil
}

func (o *Database) Delete(key string) (bool, error) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	_, ok := o.base[key]
	if !ok {
		errMsg := fmt.Errorf("key %v not found", key)
		return false, errMsg
	}
	delete(o.base, key)
	return true, nil
}

func (o *Database) Exist(key string) bool {
	_, ok := o.base[key]
	return ok
}

func (o *Database) Save() error {
	b, err := json.Marshal(o.base)
	if err != nil {
		return err
	}
	err13 := ioutil.WriteFile("MaxBase.txt", b, 0777)
	return err13
}

func (o *Database) Quit() error {
	defer o.mutex.Unlock()
	o.Save()
	os.Exit(0)

	return nil
}

func (o *Database) Load() error {
	data, err := ioutil.ReadFile("MaxBase.txt")
	if err != nil {
		return err
	}

	var m map[string]string

	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	o.base = m

	return nil
}

func (o *Database) Sum() (int, error) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	sum := 0
	for _, i := range o.base {
		value, _ := strconv.Atoi(i)
		sum += value
	}
	return sum, nil
}

func (o *Database) Avg() (int, error) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	sum := 0
	l := 0
	for _, i := range o.base {
		value, _ := strconv.Atoi(i)
		sum += value
		l += 1
	}
	avg := sum / l
	return avg, nil
}

func (o *Database) Gt(val string) error {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	for key, v := range o.base {
		value, _ := strconv.Atoi(v)
		val1, _ := strconv.Atoi(val)
		if value > val1 {
			fmt.Println("Key: ", key, "Value: ", value)
		}
	}
	return nil
}

func (o *Database) Lt(val string) error {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	for key, v := range o.base {
		value, _ := strconv.Atoi(v)
		val, _ := strconv.Atoi(val)
		if value < val {
			fmt.Println("Key: ", key, "Value: ", value)
		}
	}
	return nil
}

func (o *Database) Eq(val string) error {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	for key, v := range o.base {
		value, _ := strconv.Atoi(v)
		val, _ := strconv.Atoi(val)
		if value == val {
			fmt.Println("Key: ", key, "Value: ", value)
		}
	}
	return nil
}

func (o *Database) Count() int {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	i := 0
	for range o.base {
		i += 1
	}
	return i
}
