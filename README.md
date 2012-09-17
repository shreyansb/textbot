Listens for texts on a Twilio number, and does clever things with them. (Sends me an email with the contents, at the moment).

You'll need a config file called `textbot.conf`, in the directory the program is run from. It's a JSON file with values for the following keys:

```
{
	"TwilioAccountSID" : "",
	"TwilioAuthToken"  : "",
	"TwilioNumber"     : "",
	"EmailRecipient"   : "",
	"EmailAddress"     : "",
	"EmailPassword"    : "",
	"EmailHost"        : "",
	"EmailPort"        : ""
}
```
