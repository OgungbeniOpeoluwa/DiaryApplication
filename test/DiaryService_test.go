package test

import (
	"Diary/src/service"
	"fmt"
	"testing"
)

func TestCreateDiary(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("shola", "opemipo")
	foundDiary := diary.FindDiaryByUsername("shola")
	name := "shola"
	if foundDiary == nil {
		t.Errorf("expected %q but actual %q ", name, foundDiary.Username())
	}

}

func TestUnlockDiary_lockDiary(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("shola", "username")
	diary.LockDiary("shola")
	foundDiary := diary.FindDiaryByUsername("shola")
	if foundDiary.IsLocked() == false {
		t.Errorf("expected true got false")
	}

	diary.UnlockDiary("shola", "username")
	foundDiary1 := diary.FindDiaryByUsername("shola")
	if foundDiary1.IsLocked() == true {
		t.Errorf("expected false got true")
	}

}

func TestAddEntry(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("username", "password")

	diary.AddEntry("username", "holiday", "vacation at maldives")

	foundFindEntry := diary.FindEntryBy("holiday", "username")

	fmt.Println(foundFindEntry)
	if foundFindEntry == nil {
		t.Errorf("expected Entry got nil")
	}
}
func TestUpdateDiary(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("username", "password")
	diary.AddEntry("username", "my home", "i love my home")

	diary.UpdateEntry("username", "my home", "i love my home !!")

	foundDiary := diary.FindEntryBy("my home", "username")

	updatedDiary := "i love my home" + "\n" + "i love my home !!"
	if foundDiary.Body() != updatedDiary {
		t.Errorf("expected is not actual")
	}

}
func TestDeleteAnEntry(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("username", "password")
	diary.AddEntry("username", "my home", "i love home")
	diary.AddEntry("username", "i want peace", "just peace im not in for this stress")

	diary.DeleteEntryBy("username", "my home")

	entry := diary.FindAllEntry("username")
	fmt.Println(entry)
	if len(entry) != 1 {
		t.Errorf("Expected 1 actual is %q", len(entry))

	}

}

func TestDeleteAllEntry(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("username", "password")
	diary.AddEntry("username", "i want peace", "peace is all i want")

	diary.DeleteAllEntry("username")

	entry := diary.FindAllEntry("username")

	if len(entry) != 0 {
		t.Errorf("expected 0 actual %q", len(entry))
	}

}

func TestDeleteDiary(t *testing.T) {
	var diary service.DiaryService = service.NewDiaryServiceImpl()
	diary.CreateDiary("username", "password")
	diary.DeleteDiary("username")
	foundDiary := diary.FindDiaryByUsername("username")
	if foundDiary != nil {
		t.Errorf("expected nil but actual is not nil")
	}

}
