package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version set at compile-time
var (
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "telegram plugin"
	app.Usage = "telegram plugin"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "telegram token",
			EnvVar: "PLUGIN_TOKEN,TELEGRAM_TOKEN,INPUT_TOKEN",
		},
		cli.StringSliceFlag{
			Name:   "to",
			Usage:  "telegram user",
			EnvVar: "PLUGIN_TO,TELEGRAM_TO,INPUT_TO",
		},
		cli.StringSliceFlag{
			Name:   "message",
			Usage:  "send telegram message",
			EnvVar: "PLUGIN_MESSAGE,TELEGRAM_MESSAGE,INPUT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "message.file",
			Usage:  "send telegram message from file",
			EnvVar: "PLUGIN_MESSAGE_FILE,TELEGRAM_MESSAGE_FILE,INPUT_MESSAGE_FILE",
		},
		cli.StringFlag{
			Name:   "template.vars",
			Usage:  "additional template vars to be used in message, as JSON string",
			EnvVar: "PLUGIN_TEMPLATE_VARS,TELEGRAM_TEMPLATE_VARS,INPUT_TEMPLATE_VARS",
		},
		cli.StringSliceFlag{
			Name:   "photo",
			Usage:  "send photo message",
			EnvVar: "PLUGIN_PHOTO,PHOTO,INPUT_PHOTO",
		},
		cli.StringSliceFlag{
			Name:   "document",
			Usage:  "send document message",
			EnvVar: "PLUGIN_DOCUMENT,DOCUMENT,INPUT_DOCUMENT",
		},
		cli.StringSliceFlag{
			Name:   "sticker",
			Usage:  "send sticker message",
			EnvVar: "PLUGIN_STICKER,STICKER,INPUT_STICKER",
		},
		cli.StringSliceFlag{
			Name:   "audio",
			Usage:  "send audio message",
			EnvVar: "PLUGIN_AUDIO,AUDIO,INPUT_AUDIO",
		},
		cli.StringSliceFlag{
			Name:   "voice",
			Usage:  "send voice message",
			EnvVar: "PLUGIN_VOICE,VOICE,INPUT_VOICE",
		},
		cli.StringSliceFlag{
			Name:   "location",
			Usage:  "send location message",
			EnvVar: "PLUGIN_LOCATION,LOCATION,INPUT_LOCATION",
		},
		cli.StringSliceFlag{
			Name:   "venue",
			Usage:  "send venue message",
			EnvVar: "PLUGIN_VENUE,VENUE,INPUT_VENUE",
		},
		cli.StringSliceFlag{
			Name:   "video",
			Usage:  "send video message",
			EnvVar: "PLUGIN_VIDEO,VIDEO,INPUT_VIDEO",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug message",
			EnvVar: "PLUGIN_DEBUG,DEBUG,INPUT_DEBUG",
		},
		cli.BoolFlag{
			Name:   "match.email",
			Usage:  "send message when only match email",
			EnvVar: "PLUGIN_ONLY_MATCH_EMAIL,INPUT_ONLY_MATCH_EMAIL",
		},
		cli.BoolTFlag{
			Name:   "webpage.preview",
			Usage:  "toggle web-page preview",
			EnvVar: "PLUGIN_WEBPAGE_PREVIEW,INPUT_WEBPAGE_PREVIEW",
		},
		cli.StringFlag{
			Name:   "format",
			Value:  formatMarkdown,
			Usage:  "telegram message format",
			EnvVar: "PLUGIN_FORMAT,FORMAT,INPUT_FORMAT",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "repository owner and repository name",
			EnvVar: "DRONE_REPO,GITHUB_REPOSITORY",
		},
		cli.StringFlag{
			Name:   "repo.namespace",
			Usage:  "repository namespace",
			EnvVar: "DRONE_REPO_OWNER,DRONE_REPO_NAMESPACE,GITHUB_ACTOR",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA,GITHUB_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF,GITHUB_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "git commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "pull.request",
			Usage:  "pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.Float64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Float64Flag{
			Name:   "job.finished",
			Usage:  "job finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "env-file",
			Usage:  "source env file",
			EnvVar: "ENV_FILE",
		},
		cli.BoolFlag{
			Name:   "github",
			Usage:  "Boolean value, indicates the runtime environment is GitHub Action.",
			EnvVar: "PLUGIN_GITHUB,GITHUB",
		},
		cli.StringFlag{
			Name:   "github.workflow",
			Usage:  "The name of the workflow.",
			EnvVar: "GITHUB_WORKFLOW",
		},
		cli.StringFlag{
			Name:   "github.action",
			Usage:  "The name of the action.",
			EnvVar: "GITHUB_ACTION",
		},
		cli.StringFlag{
			Name:   "github.event.name",
			Usage:  "The webhook name of the event that triggered the workflow.",
			EnvVar: "GITHUB_EVENT_NAME",
		},
		cli.StringFlag{
			Name:   "github.event.path",
			Usage:  "The path to a file that contains the payload of the event that triggered the workflow. Value: /github/workflow/event.json.",
			EnvVar: "GITHUB_EVENT_PATH",
		},
		cli.StringFlag{
			Name:   "github.workspace",
			Usage:  "The GitHub workspace path. Value: /github/workspace.",
			EnvVar: "GITHUB_WORKSPACE",
		},
		cli.StringFlag{
			Name:   "deploy.to",
			Usage:  "Provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.StringFlag{
			Name:   "socks5",
			Usage:  "Socks5 proxy URL",
			EnvVar: "PLUGIN_SOCKS5,SOCKS5",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		GitHub: GitHub{
			Workflow:  c.String("github.workflow"),
			Workspace: c.String("github.workspace"),
			Action:    c.String("github.action"),
			EventName: c.String("github.event.name"),
			EventPath: c.String("github.event.path"),
		},
		Repo: Repo{
			FullName:  c.String("repo"),
			Namespace: c.String("repo.namespace"),
			Name:      c.String("repo.name"),
		},
		Commit: Commit{
			Sha:     c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Branch:  c.String("commit.branch"),
			Link:    c.String("commit.link"),
			Author:  c.String("commit.author"),
			Email:   c.String("commit.author.email"),
			Avatar:  c.String("commit.author.avatar"),
			Message: c.String("commit.message"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Link:     c.String("build.link"),
			Started:  c.Float64("job.started"),
			Finished: c.Float64("job.finished"),
			PR:       c.String("pull.request"),
			DeployTo: c.String("deploy.to"),
		},
		Config: Config{
			Token:        c.String("token"),
			Debug:        c.Bool("debug"),
			MatchEmail:   c.Bool("match.email"),
			WebPreview:   c.Bool("webpage.preview"),
			To:           c.StringSlice("to"),
			Message:      c.StringSlice("message"),
			MessageFile:  c.String("message.file"),
			TemplateVars: c.String("template.vars"),
			Photo:        c.StringSlice("photo"),
			Document:     c.StringSlice("document"),
			Sticker:      c.StringSlice("sticker"),
			Audio:        c.StringSlice("audio"),
			Voice:        c.StringSlice("voice"),
			Location:     c.StringSlice("location"),
			Video:        c.StringSlice("video"),
			Venue:        c.StringSlice("venue"),
			Format:       c.String("format"),
			GitHub:       c.Bool("github"),
			Socks5:       c.String("socks5"),
		},
	}

	return plugin.Exec()
}
