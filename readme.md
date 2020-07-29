# Airtable Slack Notify

This bot check all new items arriving in the airtable base of DDFS and send a notification to a given Slack Channel. You can use a `.env` file, environment variable or command line argument to launch it.

This bot starts a web server listening on port 8080 and responds on the url defined in environment variable `URL_CHECK_ALL_NEW`.

# Airtable prerequisite

You need to have an airtable base with tables that you want to check. In your table, you must have fields:

- `name` (string) as the name of the item
- `contact_mail` (string) as the email of the person who proposed this item
- `status` (string / single item list) : Only items with status `proposal` will be handled

# Settings

Here are the environment variables that you have to set while using this project:

- `API_KEY`: API key that you will have to add to your request with `?api_key=API_KEY`. It is advised to use an UUID
- `AIRTABLE_API_URL`: Default Airtable API URL
- `AIRTABLE_API_KEY`: API Key to access Airtable API (https://support.airtable.com/hc/en-us/articles/219046777-How-do-I-get-my-API-key-)
- `AIRTABLE_BASE`: ID of the Airtable Base
- `AIRTABLE_TABLES`: Tables to check in airtable. You can provide multiple tables in the same string delimited by a comma (e.g: `Users,Cars,Appointments`)
- `SLACK_WEBHOOK_URL`: Your Slack webhook url to trigger to send a new notification message (https://api.slack.com/messaging/webhooks)

# Deploy to Google Cloud Run

This bot can be easily deployed on [Google Cloud Run](https://cloud.google.com/run?hl=fr). Just click here:

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)
