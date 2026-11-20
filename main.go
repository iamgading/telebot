package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN belum di-set. Contoh: $env:BOT_TOKEN=\"token-dari-botfather\"")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatalf("Gagal membuat bot: %v", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/halo", bot.MatchTypeExact, haloHandler)

	log.Println("Bot berjalan... tekan Ctrl+C untuk berhenti")
	b.Start(ctx)
}

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Halo! Saya bot Telegram pertama kamu. \n\nPerintah yang tersedia:\n/start - mulai\n/halo - sapa bot\n\nKirim pesan apa pun, saya akan balas!",
	})
	if err != nil {
		log.Printf("Gagal mengirim pesan: %v", err)
	}
}

func haloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Halo juga! 👋",
	})
	if err != nil {
		log.Printf("Gagal mengirim pesan: %v", err)
	}
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.Text == "" {
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Kamu bilang: " + strings.TrimSpace(update.Message.Text),
	})
	if err != nil {
		log.Printf("Gagal mengirim pesan: %v", err)
	}
}
