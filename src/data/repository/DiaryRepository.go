package repository

import (
	"Diary/src/data/model"
	"strconv"
)

type DiaryRepository interface {
	Save(diary *model.Diary)
	FindById(id string) *model.Diary
	DeleteById(id string)
	DeleteAll()
	FindAll() *[]model.Diary
	Count() int
}

type DiaryRepositoryImpl struct {
	Diaries []model.Diary
	count   int
}

func (e *DiaryRepositoryImpl) Save(diary *model.Diary) {
	user := e.FindById(diary.Id())
	if user == nil {
		id := e.generateId()
		diary.SetId(strconv.Itoa(id))
		e.Diaries = append(e.Diaries, *diary)
	} else {
		e.DeleteById(user.Id())
		e.Diaries = append(e.Diaries, *diary)
	}

}
func (e *DiaryRepositoryImpl) FindById(id string) *model.Diary {
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

func (e *DiaryRepositoryImpl) DeleteById(id string) {
	for index, diary := range e.Diaries {
		if diary.Id() == id {
			e.Diaries = append(e.Diaries[:index], e.Diaries[index+1:]...)
		}
	}

}

func (e *DiaryRepositoryImpl) FindAll() *[]model.Diary {
	return &e.Diaries

}

func (e *DiaryRepositoryImpl) DeleteAll() {
	var diary []model.Diary
	e.Diaries = diary

}

func (e *DiaryRepositoryImpl) Count() int {
	return len(e.Diaries)

}
func (e *DiaryRepositoryImpl) generateId() int {
	e.count += 1
	return e.count

}
