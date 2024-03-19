package repository

import (
	"Diary/src/data/model"
	"strconv"
)

type EntryRepository interface {
	Save(entry *model.Entry)
	FindById(id string) *model.Entry
	DeleteById(id string)
	DeleteAll()
	FindAll() *[]model.Entry
	Count() int
}

type EntryRepositoryImpl struct {
	entries []model.Entry
	count   int
}

func (e *EntryRepositoryImpl) Save(entry *model.Entry) {
	user := e.FindById(entry.Id())
	if user == nil {
		id := e.generateId()
		entry.SetId(strconv.Itoa(id))
		e.entries = append(e.entries, *entry)
	} else {
		e.DeleteById(user.Id())
		e.entries = append(e.entries, *entry)
	}

}

func (e *EntryRepositoryImpl) FindById(id string) *model.Entry {
	var all = e.FindAll()
	if all != nil {
		for index := 0; index < len(*all); index++ {
			if (*all)[index].Id() == id {
				value := (*all)[index]
				return &value
			}

		}
	}
	return nil

}
func (e *EntryRepositoryImpl) FindAll() *[]model.Entry {
	return &e.entries

}

func (e *EntryRepositoryImpl) DeleteAll() {
	var entry []model.Entry
	e.entries = entry

}

func (e *EntryRepositoryImpl) DeleteById(id string) {
	for index, entry := range e.entries {
		if entry.Id() == id {
			e.entries = append(e.entries[:index], e.entries[index+1:]...)
		}
	}

}
func (e *EntryRepositoryImpl) Count() int {
	return len(e.entries)

}

func (e *EntryRepositoryImpl) generateId() int {
	e.count += 1
	return e.count

}
