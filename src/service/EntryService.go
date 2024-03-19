package service

import (
	"Diary/src/data/model"
	"Diary/src/data/repository"
)

type EntryService interface {
	CreateEntry(title string, body string, id string)
	FindAllEntryBelongingTo(diaryId string) []model.Entry
	UpdateEntry(title string, diaryId string, updatedBody string)
	FindAnEntry(title string, diaryId string) *model.Entry
	DeleteAnEntry(title string, diaryId string)
	DeleteAllEntries(diaryId string)
}

type EntryServiceImpl struct {
	repository repository.EntryRepository
}

func NewEntryServiceImpl() *EntryServiceImpl {
	return &EntryServiceImpl{repository: new(repository.EntryRepositoryImpl)}
}

func (e *EntryServiceImpl) CreateEntry(title string, body string, diaryId string) {
	if !e.checkTitleExist(title, diaryId) {
		entry := model.NewEntry(title, body, diaryId)
		e.repository.Save(entry)
	}

}

func (e *EntryServiceImpl) checkTitleExist(title string, id string) bool {
	entries := e.FindAllEntryBelongingTo(id)
	for _, entry := range entries {
		if entry.Title() == title {
			return true
		}

	}
	return false

}

func (e *EntryServiceImpl) FindAllEntryBelongingTo(id string) []model.Entry {
	var userEntries []model.Entry
	entries := e.repository.FindAll()
	for _, entry := range *entries {
		if entry.DiaryId() == id {
			userEntries = append(userEntries, entry)
		}

	}
	return userEntries

}

func (e *EntryServiceImpl) UpdateEntry(title string, diaryId string, updateBody string) {
	oldEntry := e.FindAnEntry(title, diaryId)
	if oldEntry != nil {
		oldBody := oldEntry.Body() + "\n" + updateBody
		oldEntry.SetBody(oldBody)
		e.repository.Save(oldEntry)

	}

}

func (e *EntryServiceImpl) FindAnEntry(title string, diaryId string) *model.Entry {
	allEntry := e.FindAllEntryBelongingTo(diaryId)
	for _, entry := range allEntry {
		if entry.Title() == title {
			return &entry
		}

	}
	return nil

}

func (e *EntryServiceImpl) DeleteAnEntry(title string, diaryId string) {
	entry := e.FindAnEntry(title, diaryId)
	if entry != nil {
		e.repository.DeleteById(entry.Id())
	}

}

func (e *EntryServiceImpl) DeleteAllEntries(diaryId string) {
	allEntry := e.FindAllEntryBelongingTo(diaryId)
	if allEntry != nil {
		for _, entry := range allEntry {
			e.repository.DeleteById(entry.Id())
		}
	}

}
