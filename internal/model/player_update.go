package model

import "online_shooter/internal/game/geometry"

type PlayerUpdateMessage struct {
	LeftKeyPressed  bool           `json:"left_key_pressed"`
	UpKeyPressed    bool           `json:"up_key_pressed"`
	RightKeyPressed bool           `json:"right_key_pressed"`
	DownKeyPressed  bool           `json:"down_key_pressed"`
	Shot            bool           `json:"shot"`
	Aim             geometry.Point `json:"aim"`
}
