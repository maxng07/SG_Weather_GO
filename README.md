# Obtaining Singapore Weather Forecast 

This is a complete rewrite of the application <a href="https://github.com/maxng07/sg-WhatsApp-Weather-DirectAPI">sg-WhatsApp-Weather-DirectAPI </a> using Golang. Application can be compiled into various binaries to run on Windows/EXE, Linux, Mac OSX and be executed locally to obtain the weather forecast of your choice town-name in Singapore. <br>
<p>
The Core Golang engine can also be deployed as a Serverless function in the Cloud (AWS/GCP/Twilio) replacing the Nodejs in <a href="https://github.com/maxng07/sg-WhatsApp-Weather-DirectAPI">sg-WhatsApp-Weather-DirectAPI </a> to GET and parse the real-time Weather forecast published by the Singapore National Environment Agency (NEA). <br>
<p>
In Release 0.1, To run it locally on Windows/MAC/Linux, you will need the config.json and the application binary. Config.json contains the URL to obtain the weather data, it is not baked into the code to (1) allow modification of the URL if NEA changes it, and (2) to demonstrated the specific URL is needed for the application to run. NEA hosted their data in GovTech API in JSON format. The code appends the current system time as a query parameter to the specified URL in config.json, as requested by GovTech API Engine. <br>

In Release 0.2, the URL is baked into the code and no longer requires config.json after obtaining feedback from users. 
<p>
The returned JSON data from Govtech API are nested in arrays and objects. The code parse these in golang struct and matches it with the input to provide the weather forecast value. <br>
<p>
The application will prompt the User for input, inputs are not case-sensitive and valid inputs are <br>
1. Town-name <br>
2. "All" keyword to print out all weather forecast <br>
3. "Area" keyword to print out all town-name. <br>

<h2> Release 1.0 </h2> 
In Release 1.0, PSI readings are added. The application now accepts additional User Input, "PSI" or "psi" will extract PSI data from NEA and present PM2.5 Hourly and PSI Hourly Readings. The NEA PSI raw data has more datasets and more details can be included in the output. Reference can be looked up at <a href> https://www.haze.gov.sg </a> for meaning of PSI values and PM2.5 Air Quality values. Datasets of NEA PSI is on <a href> https://data.gov.sg/dataset/psi </a>
<p>
<h2> Release 2.0 </h2>
In Release 2.0, Real-time Air Temperature Readings across Singapore are added. The data are retrieve from <a href="https://data.gov.sg/dataset/realtime-weather-readings"> Realtime Weather Readings across Singapore </a> maintained by GovTech. The Air Temperature readings in degreeC are taken on 5 minute time interval across the Weather Stations in Singapore. A new User Input, "Temp" or "temp" will extract JSON data for Air Temperature Readings. Here is an overview of the <a href="https://data.gov.sg/dataset/realtime-weather-readings?view_id=81d91c06-158d-4f01-abd9-12108d954847&resource_id=17494bed-23e9-4b3b-ae89-232f87987163"> locations </a> on the Weather Stations </p>
The application logic is re-written for easier maintenance for future add-on. <br>
  
<h2>Demonstration </h2>
The animated gif will show a sample output of the <a href="https://github.com/maxng07/SG_Weather_GO/blob/master/weather_forecast.gif"> Weather App </a> and the Release1.0 <a href="https://github.com/maxng07/SG_Weather_GO/blob/master/psi.gif"> PSI readings enhancement </a>
</p>
The core Golang engine can be used as webhook on Cloud Communication Provider such as Twilio, Nexmo, Messagebird and others and send the weather forecast data as SMS, Voice, WhatsApp, LINE, Google Assistant.

<h2>Installation </h2>
Get the latest version for Windows and macOSX version <a href="https://github.com/maxng07/SG_Weather_GO/releases">here </a><br>
  
