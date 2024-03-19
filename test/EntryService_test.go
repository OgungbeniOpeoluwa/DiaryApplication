package test

import (
	"Diary/src/service"
	"testing"
)

func TestCreateEntry(t *testing.T) {
	entry := service.NewEntryServiceImpl()
	title := "I like me"
	body := "why i like me ooooo"
	entry.CreateEntry(title, body, "1")
	actual := len(entry.FindAllEntryBelongingTo("1"))
	output := 1
	if actual != output {
		t.Errorf("expected %q but got %q", output, actual)
	}

}

func TestTitleCantBeDuplicated(t *testing.T) {
	title := "I like me"
	body := "why i like me ooooo"
	service := service.NewEntryServiceImpl()
	service.CreateEntry(title, body, "1")
	title = "I like me"
	body = "why i like me and me ooooo"
	service.CreateEntry(title, body, "1")
	actual := service.FindAllEntryBelongingTo("1")
	output := 1
	if len(actual) != output {
		t.Errorf("expected %q but got %q", output, actual)
	}

}
func TestWhenEntryIsUpdatedAndSaved_RepositoryStillMaintainItLength(t *testing.T) {
	title := "I like me"
	body := "why i like me"
	service := service.NewEntryServiceImpl()
	service.CreateEntry(title, body, "1")
	body = "total me"
	service.UpdateEntry(title, "1", body)
	actual := service.FindAnEntry(title, "1")
	output := "why i like me" + "\n" + "total me"
	if actual.Body() != output {
		t.Errorf("expected %q but got %q", output, actual)
	}

}

func TestThatUserCanDeleteAnEntry(t *testing.T) {
	entry := service.NewEntryServiceImpl()
	title := "I like me"
	body := "why i like me ooooo"
	entry.CreateEntry(title, body, "1")
	entry.DeleteAnEntry(title, "1")
	actual := len(entry.FindAllEntryBelongingTo("1"))
	output := 0
	if actual != output {
		t.Errorf("expected %q but got %q", output, actual)
	}

}

func TestDeleteAll(t *testing.T) {
	entry := service.NewEntryServiceImpl()
	title := "I like me"
	body := "why i like me ooooo"
	entry.CreateEntry(title, body, "1")
	entry.DeleteAnEntry(title, "1")
	actual := len(entry.FindAllEntryBelongingTo("1"))
	output := 0
	if actual != output {
		t.Errorf("expected %q but got %q", output, actual)
	}

}
