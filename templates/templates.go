package templates

var All = map[string]string{
	"/hltv":     "🎯 *{{.team1}}* vs *{{.team2}}*\nScore: {{.score}}\nEvent: {{.event}}",
	"/calendar": "📅 *{{.title}}*\n🕐 {{.time}}{{with .location}}\n📍 {{.}}{{end}}",
	"/email":    "✉️ *{{.subject}}*\nFrom: {{.from}}\n\n{{.preview}}",
	"/hn":       "🟠 *{{.title}}*\nby {{.by}} • {{.score}} pts • {{.comments}} comments\n{{.url}}",
	"/github":   "🐙 *{{.repo}}* — {{.event}}\n{{with .actor}}by {{.}}\n{{end}}{{.title}}{{with .url}}\n{{.}}{{end}}",
}
