package alexa_maker

import (
	"cocktail/bin/models"
	"cocktail/bin/services/database"
	"encoding/json"
	"fmt"
	"os"
)

func DevicesCreate() {

	var maker []models.DrinkModel

	database.Db.Find(&maker)

	filePath := "/opt/api_cocktail/alexa_maker/config.json"

	devices := make([]models.DeviceConfig, len(maker))

	for i, value := range maker {

		device := models.DeviceConfig{
			Port:    888 + i,
			OnCmd:   "http://localhost/myserver/fakeswitches/01/on",
			OffCmd:  "http://192.168.1.104/myserver/fakeswitches/01/off",
			OnData:  models.DrinkMaker{Recipe: value.Recipe},
			OffData: models.DrinkMaker{Recipe: value.Recipe},
			Method:  "POST",
			Name:    value.Name,
		}

		devices[i] = device
	}

	config := models.Config{
		Fauxmo: models.FauxmoConfig{
			IPAddress: "auto",
		},
		Plugins: models.PluginsConfig{
			SimpleHTTPPlugin: models.SimpleHTTPPluginConfig{
				Devices: devices,
			},
		},
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo %s criado com sucesso!\n", filePath)

}
