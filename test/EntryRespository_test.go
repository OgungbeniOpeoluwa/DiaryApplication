package test

import (
	"Diary/src/data/model"
	"Diary/src/data/repository"
	"testing"
)

func TestEntryRepository_save(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	firstEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(firstEntry)
	expected := len(*repo.FindAll())
	result := 1
	if expected != result {
		t.Errorf("expected %q but got %q", result, expected)
	}

}

func TestEntryRepository_ResavingExistingEntryAfterUpdating(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)

	firstEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(firstEntry)

	findEntry := repo.FindById("1")
	findEntry.SetBody("my love for school")
	repo.Save(findEntry)

	expected := len(*repo.FindAll())
	result := 1
	if expected != result {
		t.Errorf("expected %d but got %d", result, expected)
	}

}
func TestEntryRepository_DeleteAll(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	firstEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(firstEntry)
	secondEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(secondEntry)
	repo.DeleteAll()

	expected := len(*repo.FindAll())
	result := 0
	if expected != result {
		t.Errorf("expected %q but got %q", result, expected)
	}

}

func TestEntryRepository_Count(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	firstEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(firstEntry)
	secondEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(secondEntry)

	expected := repo.Count()
	result := 2
	if expected != result {
		t.Errorf("expected %q but got %q", result, expected)
	}

}
func TestEntryRepository_DeleteById(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	firstEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(firstEntry)
	secondEntry := model.NewEntry("My last day in school", "i had fun", "1")
	repo.Save(secondEntry)

	repo.DeleteById("1")
	expected := len(*repo.FindAll())
	result := 1
	if expected != result {
		t.Errorf("expected %q but got %q", result, expected)
	}

}
