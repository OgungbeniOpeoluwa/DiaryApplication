package service

import (
	"Diary/src/data/model"
	"Diary/src/data/repository"
)

type DiaryService interface {
	CreateDiary(username string, password string)
	UnlockDiary(username string, password string)
	LockDiary(username string)
	AddEntry(username string, title string, body string)
	UpdateEntry(username string, title string, body string)
	DeleteEntryBy(username string, title string)
	DeleteAllEntry(username string)
	FindEntryBy(title string, username string) *model.Entry
	FindAllEntry(username string) []model.Entry
	FindDiaryByUsername(username string) *model.Diary
	DeleteDiary(username string)
}

type DiaryServiceImpl struct {
	repository repository.DiaryRepository
	entry      EntryService
}

func NewDiaryServiceImpl() *DiaryServiceImpl {
	return &DiaryServiceImpl{repository: new(repository.DiaryRepositoryImpl), entry: NewEntryServiceImpl()}
}

func (d *DiaryServiceImpl) CreateDiary(username string, password string) {
	diary := d.FindDiaryByUsername(username)
	if diary == nil {
		newDiary := model.NewDiary(username, password)
		d.repository.Save(newDiary)
	}

}
func (d *DiaryServiceImpl) UnlockDiary(username string, password string) {
	diary := d.FindDiaryByUsername(username)
	if diary != nil && diary.IsLocked() && diary.Password() == password {
		diary.SetIsLocked(false)
		d.repository.Save(diary)

	}

}

func (d *DiaryServiceImpl) LockDiary(username string) {
	diary := d.FindDiaryByUsername(username)
	if diary != nil && !diary.IsLocked() {
		diary.SetIsLocked(true)
		d.repository.Save(diary)

	}

}

func (d *DiaryServiceImpl) AddEntry(username string, title string, body string) {
	diary := d.FindDiaryByUsername(username)
	if diary != nil {
		d.entry.CreateEntry(title, body, diary.Id())

	}

}

func (d *DiaryServiceImpl) UpdateEntry(username string, title string, body string) {
	foundDiary := d.FindDiaryByUsername(username)
	if foundDiary != nil {
		d.entry.UpdateEntry(title, foundDiary.Id(), body)
	}
}

func (d *DiaryServiceImpl) DeleteEntryBy(username string, title string) {
	foundDiary := d.FindDiaryByUsername(username)
	if foundDiary != nil {
		d.entry.DeleteAnEntry(title, foundDiary.Id())
	}

}

func (d *DiaryServiceImpl) DeleteAllEntry(username string) {
	foundDiary := d.FindDiaryByUsername(username)
	if foundDiary != nil {
		d.entry.DeleteAllEntries(foundDiary.Id())
	}

}

func (d *DiaryServiceImpl) FindEntryBy(title string, username string) *model.Entry {
	diary := d.FindDiaryByUsername(username)
	if diary != nil {
		return d.entry.FindAnEntry(title, diary.Id())
	}
	return nil

}
func (d *DiaryServiceImpl) FindAllEntry(username string) []model.Entry {
	foundDiary := d.FindDiaryByUsername(username)
	if foundDiary != nil {
		entries := d.entry.FindAllEntryBelongingTo(foundDiary.Id())
		return entries

	}
	return nil

}
func (d *DiaryServiceImpl) FindDiaryByUsername(username string) *model.Diary {
	allDairies := d.repository.FindAll()
	if allDairies != nil {
		for _, diary := range *allDairies {
			if diary.Username() == username {
				return &diary
			}

		}
	}
	return nil
}
func (d *DiaryServiceImpl) DeleteDiary(username string) {
	foundDiary := d.FindDiaryByUsername(username)
	if foundDiary != nil {
		d.repository.DeleteById(foundDiary.Id())

	}

}
