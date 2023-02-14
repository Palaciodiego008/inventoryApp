package actions_test

import (
	"inventoryApp/app/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
)

func (as *ActionSuite) Test_Show_Line() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newLine := models.Line{
		ID:               id,
		Carrier:          "Carrier Test",
		Iccid:            "233333",
		AssociatedTo:     "Test",
		AssociatedDevice: "Mobile test",
		EndContractDate:  "12-22-09",
		UpgradeEligibity: time.Now(),
		Status:           "Suspended",
	}

	as.NoError(as.DB.Create(&newLine))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	res = as.HTML("/lines/468d02bb-98ac-4496-af51-62c3c1f55530/show").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Carrier Test")
	as.Contains(res.Body.String(), "Test")
	as.Contains(res.Body.String(), "Mobile test")
	as.Contains(res.Body.String(), "12-22-09")
	as.Contains(res.Body.String(), "Suspended")

}

func (as *ActionSuite) Test_List_Line() {
	newLine := models.Line{
		Carrier:          "Carrier Test",
		AssociatedTo:     "Test",
		AssociatedDevice: "Mobile test",
		EndContractDate:  "12-22-09",
		UpgradeEligibity: time.Now(),
		Status:           "Suspended",
	}

	as.NoError(as.DB.Create(&newLine))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Carrier Test")
	as.Contains(res.Body.String(), "Test")
	as.Contains(res.Body.String(), "Mobile test")
	as.Contains(res.Body.String(), "12-22-09")
	as.Contains(res.Body.String(), "Suspended")
}

func (as *ActionSuite) Test_New_Line() {
	res := as.HTML("/line/new/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Carrier")
	as.Contains(res.Body.String(), "PhoneLine")
	as.Contains(res.Body.String(), "Iccid")
	as.Contains(res.Body.String(), "AssociatedTo")
	as.Contains(res.Body.String(), "AssociatedDevice")
	as.Contains(res.Body.String(), "Status")
	as.Contains(res.Body.String(), "EndContractDate")

}

func (as *ActionSuite) Test_Create_Line() {
	res := as.HTML("/line/new").Get()
	as.Equal(http.StatusOK, res.Code)
	line := url.Values{
		"Carrier":          []string{"Tigo"},
		"PhoneLine":        []string{"32434343"},
		"Iccid":            []string{"4343343"},
		"AssociatedTo":     []string{"Free"},
		"AssociatedDevice": []string{"Huawei"},
		"Status":           []string{"Deactivated"},
		"EndContractDate":  []string{"12-09-2020"},
	}

	as.TableChange("lines", 1, func() {
		res := as.HTML("/create/line").Post(line)
		as.Equal(http.StatusSeeOther, res.Code)
		as.Equal(res.Location(), "/")
	})

}

func (as *ActionSuite) Test_Edit_Line() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newLine := models.Line{
		ID:               id,
		Carrier:          "Carrier Test",
		Iccid:            "233300",
		AssociatedTo:     "Test",
		AssociatedDevice: "Mobile test",
		EndContractDate:  "12-22-09",
		UpgradeEligibity: time.Now(),
		Status:           "Suspended",
	}

	as.NoError(as.DB.Create(&newLine))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	res = as.HTML("/lines/468d02bb-98ac-4496-af51-62c3c1f55530/edit").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Edit Line:")

}
func (as *ActionSuite) Test_Update_Line() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newLine := models.Line{
		ID:               id,
		Carrier:          "Carrier Test2",
		Iccid:            "23939",
		AssociatedTo:     "Free",
		AssociatedDevice: "Mobile test huawei",
		EndContractDate:  "12-12-2021",
		UpgradeEligibity: time.Now(),
		Status:           "Active",
	}

	as.NoError(as.DB.Create(&newLine))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	res = as.HTML("/lines/468d02bb-98ac-4496-af51-62c3c1f55530/edit").Get()
	as.Equal(http.StatusOK, res.Code)
	values := url.Values{
		"Carrier":          []string{"Carrier"},
		"PhoneLine":        []string{"PhoneLine"},
		"Iccid":            []string{"23939"},
		"AssociatedTo":     []string{"AssociatedTo"},
		"AssociatedDevice": []string{"AssociatedDevice"},
		"Status":           []string{"Status"},
		"EndContractDate":  []string{"EndContractDate"},
	}
	as.TableChange("lines", 0, func() {
		res = as.HTML("/lines/468d02bb-98ac-4496-af51-62c3c1f55530/update").Put(values)
		as.Equal(http.StatusSeeOther, res.Code)
	})

	res = as.HTML("/").Get()
	as.Contains(res.Body.String(), "Carrier")
	as.Contains(res.Body.String(), "PhoneLine")
	as.Contains(res.Body.String(), "Iccid")
	as.Contains(res.Body.String(), "AssociatedTo")
	as.Contains(res.Body.String(), "AssociatedDevice")
	as.Contains(res.Body.String(), "Status")
	as.Contains(res.Body.String(), "EndContractDate")
}

func (as *ActionSuite) Test_Update_Task() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	as.NoError(err)
	newStatus := models.Line{
		ID:     id,
		Status: "Suspended",
	}

	as.NoError(as.DB.Create(&newStatus))

	res := as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)

	res = as.HTML("/lines/468d02bb-98ac-4496-af51-62c3c1f55530/edit").Get()
	as.Equal(http.StatusOK, res.Code)
	values := url.Values{
		"Status": []string{"Status"},
	}
	as.TableChange("lines", 0, func() {
		res = as.HTML("/lines/468d02bb-98ac-4496-af51-62c3c1f55530/change-status").Put(values)
		as.Equal(http.StatusSeeOther, res.Code)
	})

	res = as.HTML("/").Get()
	as.Contains(res.Body.String(), "Status")

}
