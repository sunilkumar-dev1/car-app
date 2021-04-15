package entities

type Variant struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Displace   int    `json:"displace"`
	PeakPower  string `json:"peak_power"`
	PeakTorque string `json:"peak_torque"`
	ModelsId   int    `json:"models_id"`
}
