package modules

var Providers = map[string]string{

	// === Cloud ===
	"AWS":              `AKIA[0-9A-Z]{16}`,
	"GCP":              `"type":\s*"service_account"|AIza[0-9A-Za-z\-_]{35}`,
	"DigitalOcean":     `dop_v1_[a-f0-9]{64}`,
	"Azure":            `AZURE_[A-Z0-9]{32}`,

	// === CI/CD ===
	"GitHub":           `gh[pousr]_[0-9A-Za-z]{36}`,
	"GitLab":           `glpat-[0-9A-Za-z\-]{20}`,
	"Bitbucket":        `bbp_[0-9A-Za-z]{32}`,
	"CircleCI":         `circleci_[0-9a-f]{40}`,
	"Buildkite":        `bk_[0-9a-zA-Z]{32}`,

	// === API / SaaS ===
	"Algolia":          `algolia_[A-Za-z0-9]{32}`,
	"Stripe":           `sk_live_[0-9a-zA-Z]{24}`,
	"PayPal":           `access_token\$production\$[0-9a-z]{16}`,
	"SendGrid":         `SG\.[A-Za-z0-9_-]{22}\.[A-Za-z0-9_-]{43}`,
	"Mailgun":          `key-[0-9a-zA-Z]{32}`,
	"Mailchimp":        `[0-9a-f]{32}-us[0-9]{1,2}`,
	"Postmark":         `[0-9a-f]{32}`,
	"Postman":          `PMAK-[a-f0-9]{24}`,
	"Shodan":           `[a-fA-F0-9]{32}`,
	"Sentry":           `[0-9a-f]{64}`,

	// === Dev / Product ===
	"Vercel":           `vercel_[A-Za-z0-9]{32}`,
	"Clerk":            `clerk_[a-zA-Z0-9]{32}`,
	"Doppler":          `dp\.pt\.[a-zA-Z0-9]{40}`,
	"LaunchDarkly":     `ldapi_[A-Za-z0-9]{32}`,
	"HuggingFace":      `hf_[A-Za-z0-9]{32}`,
	"JFrog":            `AKCp[A-Za-z0-9]{10,}`,
	"MongoDB":          `mongodb\+srv:\/\/.*`,
	"Postgres":         `postgres:\/\/.*`,
	"RabbitMQ":         `amqp:\/\/.*`,

	// === Messaging ===
	"Slack":            `xox[baprs]-[0-9a-zA-Z]{10,48}`,
	"Telegram":         `[0-9]{8,10}:[a-zA-Z0-9_-]{35}`,
	"Twilio":           `SK[0-9a-fA-F]{32}`,

	// === AI ===
	"OpenAI":           `sk-[A-Za-z0-9]{48}`,
	"Nvidia":           `nvapi-[A-Za-z0-9]{32}`,

	// === Payments ===
	"Razorpay":         `rzp_live_[0-9a-zA-Z]{14}`,
	"Paystack":         `sk_live_[0-9a-z]{32}`,
	"Square":           `sq0atp-[0-9A-Za-z]{22}`,

	// === Misc ===
	"Notion":           `secret_[A-Za-z0-9]{43}`,
	"Trello":           `[0-9a-f]{32}`,
	"Zendesk":          `[0-9a-z]{40}`,
}
