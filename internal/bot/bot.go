package bot

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/https-whoyan/MafiaCore/roles"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"
)

type Bot struct {
	sync.RWMutex
	token   string
	userID  snowflake.ID
	sess    bot.Client
	members map[snowflake.ID]*discord.Member
	roles   map[snowflake.ID]*roles.Role
	logger  *log.Logger
}

func NewBot(config *Config, logger *log.Logger) *Bot {
	var b *Bot
	botPrefix := "Bot " + strconv.Itoa(config.Num) + " " + logger.Prefix()
	logger = log.New(
		log.Writer(),
		botPrefix,
		log.Flags(),
	)
	b = &Bot{
		token:   config.Token,
		members: make(map[snowflake.ID]*discord.Member),
		roles:   make(map[snowflake.ID]*roles.Role),
		logger:  logger,
	}
	return b
}

func (b *Bot) Init() error {
	sess, err := disgo.New(
		b.token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsNonPrivileged,
			),
		),
	)
	if err != nil {
		return err
	}
	b.sess = sess
	return nil
}

func (b *Bot) Run(ctx context.Context) error {
	sc := make(chan os.Signal, 1)
	if err := b.sess.OpenGateway(ctx); err != nil {
		return err
	}
	b.userID = b.sess.ID()
	defer func(ctx context.Context) {
		b.sess.Close(ctx)
	}(ctx)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	return nil
}

func (b *Bot) Close(ctx context.Context) error {
	b.sess.Close(ctx)
	return nil
}
