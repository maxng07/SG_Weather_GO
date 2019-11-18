# Singapore Weather Forecast, Air Quality and Air Temperature Readings

This is a complete rewrite of the application <a href="https://github.com/maxng07/sg-WhatsApp-Weather-DirectAPI">sg-WhatsApp-Weather-DirectAPI </a> using Golang. Application can be compiled into binaries to run on Windows/EXE, Linux, Mac OSX and be executed locally to obtain the Weather Forecast, PSI and Temperature readings across Singapore. <br>
<p>
The Core Golang engine can also be deployed as a Serverless function in the Cloud (AWS/GCP/Twilio) replacing the Nodejs in <a href="https://github.com/maxng07/sg-WhatsApp-Weather-DirectAPI">sg-WhatsApp-Weather-DirectAPI </a> to GET and parse the real-time Weather forecast published by the Singapore National Environment Agency (NEA). <br>
<p>
The NEA Weather Forecast JSON data from Govtech API are nested in arrays and objects. The code parse these in golang struct and matches it with the user's input to provide the weather forecast value. <br>
<p>
The application will prompt the User for input, inputs are not case-sensitive and valid inputs are <br>
1. Town-name <br>
2. "All" keyword to print out all weather forecast <br>
3. "Area" keyword to print out all town-name. <br>
</p>
For the definition of Singapore Meteorological terminalogy used in the Weather Forecast, please refer to <a href="http://www.weather.gov.sg/forecasting-2/"> Weather Descriptor and Definition </a> used in forecast.


<h2> Release 1.0 </h2> 
In Release 1.0, PSI readings are added. The application now accepts additional User Input, "PSI" or "psi" will extract PSI data from NEA and present PM2.5 Hourly and PSI Hourly Readings. The NEA PSI raw data has more datasets and more details can be included in the output. Reference can be looked up at <a href> https://www.haze.gov.sg </a> for meaning of PSI values and PM2.5 Air Quality values. Datasets of NEA PSI is on <a href> https://data.gov.sg/dataset/psi </a>
<p>
<h2> Release 2.0 </h2>
In Release 2.0, Real-time Air Temperature Readings across Singapore are added. The data are retrieve from <a href="https://data.gov.sg/dataset/realtime-weather-readings"> NEA Real Time Weather Forecast across Singapore </a> maintained by GovTech. The Air Temperature readings in degreeC are taken on 5 minute time interval from the Weather Stations distributed across Singapore. To read this, a new User Input is supported, "Temp" or "temp" will extract JSON data for Air Temperature Readings. Here is the google map of the <a href="https://data.gov.sg/dataset/realtime-weather-readings?view_id=81d91c06-158d-4f01-abd9-12108d954847&resource_id=17494bed-23e9-4b3b-ae89-232f87987163"> locations </a> on the Weather Stations </p>
The application logic is also re-written for easier maintenance for any future feature add-on. <br>

# Usage - User Input
Run the application program either by executing it in Command Prompt/Terminal Window or clicking from Finder Window.
Application will prompt you for input. These are

Weather Forecast <br>
1. "Town-name" <br>
2. "All" keyword to print out all weather forecast <br>
3. "Area" keyword to print out all town-name. <br>

Air Quality - PSI <br>
4. "PSI" keyword to print PSI PM2.5 and PSI Aggregate readings across 5 Regions in Singapore

Real Time Air Temperature <br>
5. "Temp" keyword to print out 5min interval Real-time Air temperature across Weather Stations in Singapore

<h2>Demonstration </h2>
The animated gif will show the output of the <a href="https://github.com/maxng07/SG_Weather_GO/blob/master/weather-psi-temp.gif"> Weather App </a>. All features are showcased.
</p>
The core Golang engine can be used as webhook on Cloud Communication Provider such as Twilio, Nexmo, Messagebird and others and send the weather forecast data as SMS, Voice, WhatsApp, LINE, Google Assistant.

<h2>Installation </h2>
Get the latest version for Windows and macOSX version <a href="https://github.com/maxng07/SG_Weather_GO/releases">here </a>
</p>
All Weather Data are from <a href="https://www.data.gov.sg">Data.gov.sg </a>
  
