package service

import "cryptodini-robo-manager/model"

type Robot struct {
	User       model.User
	Cryptodini ICryptodiniService
}

// NewRobot strart new robot entity
func NewRobot(cryp ICryptodiniService) (*Robot, error) {
	return &Robot{
		Cryptodini: cryp,
	}, nil
}

// RobotAlgorithm robot will getport and adjust automatically
func (r *Robot) RobotAlgorithm() {
	r.Cryptodini.GetPort(r.User.UserID)
	r.Cryptodini.Adjust(r.User.UserID, &r.User.Ports[0])
}
