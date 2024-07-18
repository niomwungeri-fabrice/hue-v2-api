package hue

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Lights list of a light types
type Lights struct {
	Errors []interface{} `json:"errors"`
	Data   []Light       `json:"data"`
}

// Light an instant of hue light
type Light struct {
	ID    string `json:"id"`
	IDV1  string `json:"id_v1"`
	Owner struct {
		RID   string `json:"rid"`
		RType string `json:"rtype"`
	} `json:"owner"`
	Metadata struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
		Function  string `json:"function"`
	} `json:"metadata"`
	ProductData struct {
		Function string `json:"function"`
	} `json:"product_data"`
	Identify  interface{} `json:"identify"`
	ServiceID int         `json:"service_id"`
	On        struct {
		On bool `json:"on"`
	} `json:"on"`
	Dimming struct {
		Brightness  float64 `json:"brightness"`
		MinDimLevel float64 `json:"min_dim_level"`
	} `json:"dimming"`
	DimmingDelta     interface{} `json:"dimming_delta"`
	ColorTemperature *struct {
		Mirek       int  `json:"mirek"`
		MirekValid  bool `json:"mirek_valid"`
		MirekSchema struct {
			MirekMinimum int `json:"mirek_minimum"`
			MirekMaximum int `json:"mirek_maximum"`
		} `json:"mirek_schema"`
	} `json:"color_temperature"`
	ColorTemperatureDelta interface{} `json:"color_temperature_delta"`
	Color                 *struct {
		XY struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"xy"`
		Gamut struct {
			Red struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"red"`
			Green struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"green"`
			Blue struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"blue"`
		} `json:"gamut"`
		GamutType string `json:"gamut_type"`
	} `json:"color"`
	Dynamics struct {
		Status       string   `json:"status"`
		StatusValues []string `json:"status_values"`
		Speed        float64  `json:"speed"`
		SpeedValid   bool     `json:"speed_valid"`
	} `json:"dynamics"`
	Alert struct {
		ActionValues []string `json:"action_values"`
	} `json:"alert"`
	Signaling struct {
		SignalValues []string `json:"signal_values"`
	} `json:"signaling"`
	Mode    string `json:"mode"`
	Effects struct {
		StatusValues []string `json:"status_values"`
		Status       string   `json:"status"`
		EffectValues []string `json:"effect_values"`
	} `json:"effects"`
	TimedEffects struct {
		StatusValues []string `json:"status_values"`
		Status       string   `json:"status"`
		EffectValues []string `json:"effect_values"`
	} `json:"timed_effects"`
	PowerUp struct {
		Preset     string `json:"preset"`
		Configured bool   `json:"configured"`
		On         struct {
			Mode string `json:"mode"`
			On   struct {
				On bool `json:"on"`
			} `json:"on"`
		} `json:"on"`
		Dimming struct {
			Mode    string `json:"mode"`
			Dimming struct {
				Brightness float64 `json:"brightness"`
			} `json:"dimming"`
		} `json:"dimming"`
		Color struct {
			Mode             string `json:"mode"`
			ColorTemperature struct {
				Mirek int `json:"mirek"`
			} `json:"color_temperature"`
		} `json:"color"`
	} `json:"powerup"`
	Type string `json:"type"`
}

// GetLights client to get light services. These are offered by devices with lighting capabilities.
func (c *Client) GetLights() (*Lights, error) {
	req, err := c.newRequest("GET", "/clip/v2/resource/light")
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get light with status " + strconv.Itoa(res.StatusCode))
	}

	var lights *Lights
	err = json.NewDecoder(res.Body).Decode(&lights)
	if err != nil {
		return nil, err
	}
	return lights, nil
}
