# This is still a prototype.
Extension of existing Golang Weather Forecast program to perform correlation between Meterological Stations (MSS) that provided Weather Forecast, PSI, Air temperature across Singapore and Stations that has TB1 Tipping Point Rain Gauge to measure Rain falls. Weather Forecast data tells you the forecasted weather conditions across Singapore, but the accuracy depends on various factors and metrics. For example, at times, Weather forecast may indicate "Shower" in "Clementi" area, it would be nice to know in near real-time if rain has arrived in "Clementi" or the intensity. Rainfall readings from the rain gauge deployed across Singapore, measured rain falling rate based on a 0.2mm gauge over 5 mins interval. It is able to provide near real-time insight on how heavy is the rain in that area within a 5mins interval. TB1 (Tipping Bucket) Rain Gauge limitation is however not able to measure light rain/drizzle (<0.2mm per 5 min interval or commonly known as light liquid precipitation).

By correlating the 2, it provides a greater insight for the user to know the weather conditions so that they can plan ahead. This type of data might be useful for avvid sports users or any users that likes to beat the weather forecast. The main objective is to provide more accurate informations for user <b>to make informed decision</b>.

The 2 datasets measured in different station locations across Singapore, will need mapping in order to provide meaningful correlation. For example, Weather forecast station in Clenenti is in location with lattitude 1.315,"longitude":103.76 and according to Google-map is near AYE highway exit into Clementi Avenue 6. The rain gauge is located in latitude":1.3337,"longitude":103.7768, from Google-map is in Ngee Ann Polytechnic along Clementi Road. Observing within Clementi over a period of 7 days, with 3 raining condition, correlating both datasets provided an accurate situation in the area. Rainfall readings not only provided information if rain has arrived but the intensity of the rain and if rain has passed (minus light rain/drizzle). This data along with Weather Forecast provided additional information for any users who can be out in the open area to seek shelter for a period of time, delay or alter their journey plans. It is to be noted that the Weather Forecast does change and reflect the current weather situation, albeit not as accurate and with some delays.

Stations.json is a dataset correlating the 2 different station location datasets from MSS. The mapping between these 2 datasets are based on proximity and some degree of local knowledge between Weather station and Rain Gauge station. However, the familiarity and local knowledge of residents will help improve the mapping. Some degree of unknown or "guessimation" are used in the mapping of selected locations where the author has no complete local information. By Open Sourcing the mapping, the author hopes, with the help of others to improves the mapping, allowing more users to benefit on the accuracy of the information.

A sample output from the correlation <br>
```
clementi 
The weather forecast for Clementi is Showers
Valid Start Time: 2019-11-22T12:00:00+08:00
Valid End Time: 2019-11-22T14:00:00+08:00 
Rain Gauge Station -  Clementi Road - Current Measurement - mm/5min: 0 
2019-11-22 12:30:00 +0800 +08
```

Dataset for correlations:
1. Mapping for correlation - stations.json
2. Weather Forecast station locations 
3. TB1 Rain Gauge station locations

References:-
1. <a href="https://www.nea.gov.sg/corporate-functions/resources/facts-figures/mq">NEA/MSS Near real time data </a>
2. <a href="http://www.weather.gov.sg/faq/"> Meterological Service Singapore </a>
3. <a href="http://www.emesystems.com/davis/documents/D7852%20Installation.pdf">Tipping Bucket Rain Gauge</a> and <a href="https://en.wikipedia.org/wiki/Rain_gauge"> Rain Gauge Wiki </a>
4. <a href="https://kids.frontiersin.org/article/10.3389/frym.2018.00038"> How to measure Rain Fall </a>
5. Google Map
6. <a href="https://data.gov.sg/dataset/realtime-weather-readings?resource_id=8bd37e06-cdd7-4ca4-9ad8-5754eb70a33d"> Near Real Time Rainfall Across Singapore </a> dataset
7. <a href="https://data.gov.sg/dataset/weather-forecast"> 2 Hour Weather Forecast </a> dataset
