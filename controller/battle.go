package controller

import (
	"github.com/oarielg/BattleGo/battle"
	"github.com/oarielg/BattleGo/data"

	"github.com/labstack/echo/v4"
)

func BattleHandler(c echo.Context) error {
	player := data.LoadCharacter(true)
	monster := data.LoadCharacter(false)

	currentBattle := battle.NewBattle(player, monster)

	params := []string{
		c.Request().PostFormValue("action"),
		c.Request().PostFormValue("spellid"),
	}

	currentBattle.BattleTurn(params)
	battleData := currentBattle.GetBattleData()

	player = currentBattle.Player
	monster = currentBattle.Monster

	if currentBattle.State == battle.InProgress {
		data.SaveCharacter(player, true)
		data.SaveCharacter(monster, false)
	} else {
		data.ResetCharacterData()
	}

	return c.Render(200, "index", battleData)
}
