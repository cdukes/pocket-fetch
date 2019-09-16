A Golang tool for backing up your Pocket history to a Postgres database.

## Setup
1. Set these environment variables:
	- `POCKET_CONSUMER_KEY` + `POCKET_ACCESS_TOKEN` - [Instructions](https://getpocket.com/developer/docs/authentication) on how to obtain these
	- `POCKET_DATABASE_URL` - Postgres URL for the database you want to send content to
2. Download the repo and `go install`
3. Run `pocket-fetch`
4. Monitor the CLI output. Once you start seeing duplicate URLs, cancel the script. 