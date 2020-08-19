package ad_engine

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo"
)

//------------------------------------------

type sample_Ad struct {
	Age struct {
		young float64
		mid   float64
		old   float64
	}
	Gender struct {
		male   float64
		female float64
	}
	content string
}

type user_det struct {
	age_group string
	gender    string
}

func Engine(c echo.Context) error {
	request := new(struct {
		Uid int `json:"Uid"`
	})

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Uid, validation.Required),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No ID")
	}

	//------------------------------------------------------API Call to get User Details
	var user_1 user_det // Here User 1 is mid aged, Female Customer
	user_1.age_group = "mid"
	user_1.gender = "female"

	//------------------------------------------------------API Call to get Ads
	var ad_1 sample_Ad // Ads will have to be scored from client facing team. This is just a sample scoring
	ad_1.Age.mid = 0.3
	ad_1.Age.old = 0.2
	ad_1.Age.young = 0.1
	ad_1.Gender.female = 0.3
	ad_1.Gender.male = 0.5

	var ad_2 sample_Ad // Ads will have to be scored from client facing team. This is just a sample scoring
	ad_2.Age.mid = 0.4
	ad_2.Age.old = 0.2
	ad_2.Age.young = 0.4
	ad_2.Gender.female = 0.5
	ad_2.Gender.male = 0.3

	var ads []sample_Ad
	ads[0] = ad_1
	ads[1] = ad_2

	//Calculate Score for each Ad
	var total float64
	var result string

	for i := range ads {
		if user_1.age_group == "mid" {
			total = total + ads[i].Age.mid
		} else if user_1.age_group == "old" {
			total = total + ads[i].Age.old
		} else if user_1.age_group == "young" {
			total = total + ads[i].Age.young
		}

		if user_1.gender == "male" {
			total = total + ads[i].Gender.male
		} else if user_1.gender == "female" {
			total = total + ads[i].Gender.female
		}

		if total > 0.7 {
			result = ads[i].content
		}
	}
	return c.JSON(http.StatusOK, result)
}
