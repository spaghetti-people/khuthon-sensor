package api

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type Conditions struct {
	Temperature          float64 `json:"temperature"`
	Humidity             float64 `json:"humidity"`
	PH                   float64 `json:"ph"`
	Rainfall             float64 `json:"rainfall"`
	SoilMoisture         float64 `json:"soil_moisture"`
	SunlightExposure     float64 `json:"sunlight_exposure"`
	WaterUsageEfficiency float64 `json:"water_usage_efficiency"`
	N                    float64 `json:"N"`
	P                    float64 `json:"P"`
	K                    float64 `json:"K"`
	SoilType             float64 `json:"soil_type"`
	WindSpeed            float64 `json:"wind_speed"`
	CO2Concentration     float64 `json:"co2_concentration"`
	CropDensity          float64 `json:"crop_density"`
	PestPressure         float64 `json:"pest_pressure"`
	UrbanAreaProximity   float64 `json:"urban_area_proximity"`
	FrostRisk            float64 `json:"frost_risk"`
}

type DailyCondition struct {
	Day        int        `json:"day"`
	Conditions Conditions `json:"conditions"`
}

type CropData struct {
	CropName        string           `json:"crop_name"`
	PlantingDate    string           `json:"planting_date"`
	DailyConditions []DailyCondition `json:"daily_conditions"`
}

var (
	data     CropData
	index    = 0
	dataLock sync.Mutex
)

func loadTestData() error {
	file, err := os.Open("test_data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	return err
}

func NewRouter() *gin.Engine {
	err := loadTestData()
	if err != nil {
		panic("test_data.json 로드 실패: " + err.Error())
	}

	r := gin.Default()

	r.GET("/daily_conditions", func(c *gin.Context) {
		dataLock.Lock()
		current := data.DailyConditions[index].Conditions
		index = (index + 1) % len(data.DailyConditions)
		dataLock.Unlock()

		c.JSON(http.StatusOK, current)
	})

	return r
}
