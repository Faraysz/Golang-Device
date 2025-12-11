package repositories

import (
    "golang-api/models"
)

var Devices []models.Device

func AddDevice(device models.Device) {
    Devices = append(Devices, device)
}

func GetDevices() []models.Device {
    return Devices
}
