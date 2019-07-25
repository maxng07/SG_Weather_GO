# Obtaining Singapore Weather Forecast 

This is a complete rewrite of the application <a href="https://github.com/maxng07/sg-WhatsApp-Weather-DirectAPI">sg-WhatsApp-Weather-DirectAPI </a> using Golang. Application can be compiled into various binaries to run on Windows/EXE, Linux, Mac OSX and be executed locally to obtain the weather forecast of your choice town-name in Singapore. 
<br>
The Core Golang engine can also be deployed as a Serverless function in the Cloud (AWS/GCP/Twilio) replacing the Nodejs in <a href="https://github.com/maxng07/sg-WhatsApp-Weather-DirectAPI">sg-WhatsApp-Weather-DirectAPI </a> to GET and parse the real-time Weather forecast published by the Singapore National Environment Agency (NEA).
<br>
To run it locally on Windows/MAC/Linux, you will need the config.json and the application binary. Config.json contains the URL to obtain the weather data, it is not baked into the code to (1) allow modification of the URL if NEA changes it, and (2) to demonstrated the specific URL is needed for the application to run. NEA hosted their data in GovTech API in JSON format. The code appends the current system time as a query parameter to the specified URL in config.json, as requested by GovTech API Engine. <br>

The returned JSON data from Govtech API are nested in arrays and objects. The code parse these in golang struct and matches it with the input to provide the weather forecast value. <br>

The application will prompt the User for input, valid inputs are <br>
1. Town-name
2. "All" keyword to print out all weather forecast
3. "Area" keyword to print out all town-name.

The core Golang engine can be used as webhook on Cloud Communication Provider such as Twilio, Nexmo, Messagebird and others and send the weather forecast data as SMS, Voice, WhatsApp, LINE, Google Assistant.
