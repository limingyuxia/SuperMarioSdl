package collision

import (
	"log"
	"superMario/model"
)

func CheckCollison(role *model.ImgTextureInfo, barriers []*model.ImgTextureInfo) {
	for _, barrier := range barriers {
		var xCollision bool = role.PosX+role.Width > barrier.PosX && role.PosX < barrier.PosX+barrier.Width
		var yCollision bool = role.PosY+role.Height > barrier.PosY && role.PosY < barrier.PosY+barrier.Height

		if xCollision && yCollision {
			log.Print("be collistion")
		}

	}

}
