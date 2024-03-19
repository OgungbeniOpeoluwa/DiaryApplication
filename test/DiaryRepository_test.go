package test

import (
	"Diary/src/data/model"
	"Diary/src/data/repository"
	"testing"
)

func TestDiaryRepository_Save(t *testing.T) {
	var value repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	value.Save(model.NewDiary("ope", "123"))
	expected := len(*value.FindAll())
	output := 1

	if expected != 1 {
		t.Errorf("expected %q but actual %q", output, expected)

	}

}

func TestDiaryRepository_FindByUsername(t *testing.T) {
	var value repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	value.Save(model.NewDiary("ope", "123"))
	expected := value.FindById("1")
	output := "ope"
	if expected.Username() != output {
		t.Errorf("expected %v but %v", output, expected)
	}
}

func TestDiaryRepository_DeleteById(t *testing.T) {
	var value repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	value.Save(model.NewDiary("ope", "123"))
	value.DeleteById("1")
	output := len(*value.FindAll())
	expected := 0
	if output != expected {
		t.Errorf("expected %v but %v", expected, output)
	}
}

func TestDiaryRepository_DeleteAll(t *testing.T) {
	var value repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	value.Save(model.NewDiary("ope", "123"))
	value.Save(model.NewDiary("Tobi", "123"))
	value.DeleteAll()
	output := len(*value.FindAll())
	expected := 0
	if output != expected {
		t.Errorf("expected %v but %v", expected, output)
	}
}
