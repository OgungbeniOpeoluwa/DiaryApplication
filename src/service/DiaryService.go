package service

import (
	"Diary/src/data/model"
	"Diary/src/data/repository"
	"Diary/src/exception"
)

type DiaryService interface {
	CreateDiary(username string, password string) (*model.Diary, error)
	UnlockDiary(username string, password string)
	LockDiary(username string)
	AddEntry(username string, title string, body string) (string, error)
	UpdateEntry(username string, title string, body string)
	DeleteEntryBy(username string, title string)
	DeleteAllEntry(username string)
	FindEntryBy(title string, username string) *model.Entry
	FindAllEntry(username string) []model.Entry
	FindDiaryByUsername(username string) *model.Diary
	DeleteDiary(username string)
	CheckIfDiaryIsUnlock(username string) bool
}

type DiaryServiceImpl struct {
	repository repository.DiaryRepository
	entry      EntryService
}

func NewDiaryServiceImpl() *DiaryServiceImpl {
	return &DiaryServiceImpl{repository: new(repository.DiaryRepositoryImpl), entry: NewEntryServiceImpl()}
}

func (d *DiaryServiceImpl) CreateDiary(username string, password string) (*model.Diary, error) {
	diary := d.FindDiaryByUsername(username)
	if diary == nil && d.CheckIfDiaryIsUnlock(username) == false {
		newDiary := model.NewDiary(username, password)
		d.repository.Save(newDiary)
		return newDiary, nil
	}
	return nil, exception.NewDiaryException("Account already exist")

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

func (d *DiaryServiceImpl) AddEntry(username string, title string, body string) (string, error) {
	diary := d.FindDiaryByUsername(username)
	if diary != nil {
		return d.entry.CreateEntry(title, body, diary.Id())
	} else {
		return "nil", exception.NewDiaryException("diary doesn't exist")
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
func (d *DiaryServiceImpl) CheckIfDiaryIsUnlock(username string) bool {
	diary := d.FindDiaryByUsername(username)
	if diary != nil && diary.IsLocked() == false {
		return true
	} else {
		return false
	}

}
