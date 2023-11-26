package db_test

import (
	"testing"
)

func TestGetProjectsByUserID(t *testing.T) {

	db := testDB()

	// create a random project
	project := db.CreateRandomProject()

	res, err := db.GetProjectsByUserID(project.UserID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if len(res) < 1 {
		t.Errorf("Expected at least 1 lang, got %d", len(res))
	}

	if res[0].ID != project.ID {
		t.Errorf("Expected project ID %d, got %d", project.ID, res[0].ID)
	}
}

func TestGetProject(t *testing.T) {

	db := testDB()

	// create a random project
	project := db.CreateRandomProject()

	res, err := db.GetProject(project.ID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if res.ID != project.ID {
		t.Errorf("Expected project ID %d, got %d", project.ID, res.ID)
	}

	if res.Name != project.Name {
		t.Errorf("Expected project name %s, got %s", project.Name, res.Name)
	}
}

func TestDeleteByIDProject(t *testing.T) {

	db := testDB()
	project := db.CreateRandomProject()

	err := db.DeleteProjectByID(project.ID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	res, err := db.GetProject(project.ID)

	if err == nil {
		t.Errorf("Expected error, got %v", res)
		return
	}
}

func TestUpdateProject(t *testing.T) {

	db := testDB()
	project := db.CreateRandomProject()

	project.Name = "new name"

	err := db.UpdateProject(project)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	res, err := db.GetProject(project.ID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if res.Name != project.Name {
		t.Errorf("Expected project name %s, got %s", project.Name, res.Name)
	}
}
