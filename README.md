Listens for texts on a Twilio number, and does clever things with them. (Adds them to an evernote notebook, at the moment).

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
