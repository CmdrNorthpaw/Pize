package main

import (
	"fmt"
	"strconv"
)

var Commands = map[string]func(*DummyDrone, ...string) (bool, string) {
	"takeoff": func(dd *DummyDrone, s ...string) (bool, string) {
		if dd.Airborne {
			return false, "Drone is already airborne"
		}
		dd.Airborne = true
		return true, "Takeoff successful, drone is airborne"
	},
	"land": func(dd *DummyDrone, s ...string) (bool, string) {
		if !dd.Airborne {
			return false, "Drone is already on the ground"
		}
		dd.Airborne = false
		return true, "Landing successful, drone is on the ground"
	},
	"streamon": func(dd *DummyDrone, s ...string) (bool, string) {
		if dd.VideoOn {
			return false, "The video stream is already active"
		}
		dd.VideoOn = true
		return true, "Video stream activated"
	},
	"streamoff": func(dd *DummyDrone, s ...string) (bool, string) {
		if !dd.VideoOn {
			return false, "The video stream isn't on."
		}
		dd.VideoOn = false
		return true, "Video stream deactivated"
	},

}

func changePosition(changed *float64, changeTo string, min float64, max float64) error {
	converted, err := strconv.ParseFloat(changeTo, 64)
	if err != nil { 
		return err
	 } else if converted < min || converted > min {
		return fmt.Errorf("%f is not within the required range", converted)
	 } else {
		 changed = &converted
		 return nil
	 }
}