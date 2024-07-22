package internal

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"MafiaBotHelper/log"
	"github.com/bwmarrin/discordgo"
	"github.com/https-whoyan/MafiaCore/roles"
)

type Bot struct {
	sync.RWMutex
	token   string
	userID  string
	sess    *discordgo.Session
	members map[string]*discordgo.Member
	roles   map[string]*roles.Role
	logger  log.Logger
}

var (
	botOnce sync.Once
)

func NewBot(token string, logger log.Logger) *Bot {
	var bot *Bot
	botOnce.Do(func() {
		bot = &Bot{
			token:   token,
			members: make(map[string]*discordgo.Member),
			roles:   make(map[string]*roles.Role),
			logger:  logger,
		}
	})
	return bot
}

func (b *Bot) Init() error {
	sess, err := discordgo.New(b.token)
	if err != nil {
		return err
	}
	if err := sess.Open(); err != nil {
		return err
	}
	b.sess = sess
	b.userID = sess.State.User.ID
	b.registerHandler()
	return nil
}

func (b *Bot) Run() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func (b *Bot) registerHandler() {
	b.sess.AddHandler(b.handle)
}
