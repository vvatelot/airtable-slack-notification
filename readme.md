# DDFS checker

This bot check all new items arriving in the airtable base of DDFS and send a notification to a given Slack Channel. You can use a `.env` file, environment variable or command line argument to launch it.
This bot starts a web server listening on port 8080 and responds on the url defined in environment variable `URL_CHECK_ALL_NEW`.

# Airtable prerequisite

You need to have an airtable base with tables that you want to check. In your table, you must have fields :

- `name` (string) as the name of the item
- `contact_mail` (string) as the email of the person who proposed this item
- `status` (string / single item list) : Only items with status `proposal` will be handled

# Use the bot

## Settings

You first have to define the parameters in a `.env` file, environment variable or command line arguments :

- `URL_CHECK_ALL_NEW` : URL that you will call to trigger the bot action. It is advised to use an UUID
- `AIRTABLE_API_KEY` : API Key to access Airtable API (https://support.airtable.com/hc/en-us/articles/219046777-How-do-I-get-my-API-key-)
- `AIRTABLE_BASE` : ID of the Airtable Base
- `AIRTABLE_TABLES` : Tables to check in airtable. You can provide multiple tables in the same string delimited by a `,`
- `SLACK_WEBHOOK_URL` : Your Slack webhook url to trigger to send a new notification message (https://api.slack.com/messaging/webhooks)

## Start the server
