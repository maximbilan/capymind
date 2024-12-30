## Setup

Run `chmod +x <path to script>` before running

## Set up Google Cloud service account

Run `./scripts/setup_gcloud_access.sh`

## Deploy cloud functions

Run `./scripts/deploy_functions.sh`

## Localize Telegram Bot description (short and full)

Run `./scripts/localize_telegram_bot.sh`

## Schedule (cron) jobs (like Morning and Evening messages)

Run `./scripts/schedule_jobs.sh`

## Bump version

Run `./scripts/bump_version.sh` (patching ex. `1.0.7 -> 1.0.8`)<br/>

Parameters:<br/>
    `--minor` ex. `1.0.7 -> 1.1.0`<br/>

To get current version, run `./scripts/get_version.sh`

## Create tag

Run `./scripts/create_tag.sh`

## Update dependencies

Run `./scripts/update_deps.sh`

## Generate test coverage badge

Run `./scripts/test_coverage_badge.sh`

