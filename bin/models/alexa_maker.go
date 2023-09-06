package models

type Config struct {
	Fauxmo  FauxmoConfig  `json:"FAUXMO"`
	Plugins PluginsConfig `json:"PLUGINS"`
}

type FauxmoConfig struct {
	IPAddress string `json:"ip_address"`
}

type PluginsConfig struct {
	SimpleHTTPPlugin SimpleHTTPPluginConfig `json:"SimpleHTTPPlugin"`
}

type SimpleHTTPPluginConfig struct {
	Devices []DeviceConfig `json:"DEVICES"`
}

type DeviceConfig struct {
	Port             int               `json:"port"`
	OnCmd            string            `json:"on_cmd"`
	OffCmd           string            `json:"off_cmd"`
	Method           string            `json:"method"`
	Name             string            `json:"name"`
	OnData           DrinkMaker        `json:"on_data"`
	OffData          DrinkMaker        `json:"off_data"`
	User             string            `json:"user"`
	Password         string            `json:"password"`
	UseFakeState     bool              `json:"use_fake_state"`
	Headers          map[string]string `json:"headers"`
	StateMethod      string            `json:"state_method"`
	StateCmd         string            `json:"state_cmd"`
	StateResponseOn  string            `json:"state_response_on"`
	StateResponseOff string            `json:"state_response_off"`
}
