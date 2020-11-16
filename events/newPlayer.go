package events

import (
	"CatLegends/game"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func NewPlayer(cb *tgbotapi.CallbackConfig, chatId int64, msgId int, bot *tgbotapi.BotAPI) {
	p := game.InitPlayer(chatId)

	p.Inventory.Items = []game.Item{
		{
			Name:        "A",
			Emoji:       game.SwordEmoji,
			Quantity:    1,
			Description: "Sword",
			Price:       10,
			Rarity:      game.CommonRarity,
		},
		{
			Name:        "B",
			Emoji:       game.BowEmoji,
			Quantity:    1,
			Description: "Bow",
			Price:       15,
			Rarity:      game.UncommonRarity,
		},
		{
			Name:        "C",
			Emoji:       game.ClothingEmoji,
			Quantity:    1,
			Description: "Chest",
			Price:       10,
			Rarity:      game.RareRarity,
		},
		{
			Name:        "D",
			Emoji:       game.JewelryEmoji,
			Quantity:    2,
			Description: "Ring",
			Price:       15,
			Rarity:      game.CommonRarity,
		},
		{
			Name:        "E",
			Emoji:       game.AccessoriesEmoji,
			Quantity:    13,
			Description: "Amulet",
			Price:       10,
			Rarity:      game.CommonRarity,
		},
		{
			Name:        "F",
			Emoji:       game.SwordEmoji,
			Quantity:    1,
			Description: "Sword",
			Price:       15,
			Rarity:      game.CommonRarity,
		},
		{
			Name:        "G",
			Emoji:       game.ClothingEmoji,
			Quantity:    1,
			Description: "Pants",
			Price:       10,
			Rarity:      game.CommonRarity,
		},
	}

	db := utils.GetDB()

	_, ok := game.GetPlayerById(chatId)

	if !ok {
		_, err := db.Players.InsertOne(db.Ctx, p)
		if err != nil {
			log.Error(err)
			cb.Text = ErrorText
			cb.ShowAlert = true
			return
		}
	}

	cb.Text = "Персонаж створений"
	msgEdit := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, existingPlayerKeyboard)
	if _, err := bot.Send(msgEdit); err != nil {
		log.Error(err)
	}
}
