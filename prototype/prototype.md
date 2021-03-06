# Move from Prototype to Full Release
Extension of existing Golang Weather Forecast program to perform correlation between Meterological Stations (MSS) that provides Weather Forecast across Singapore and Stations that has TB1 Tipping Bucket Rain Gauge to measure Rain falls. Weather Forecast data tells you the forecasted weather conditions across Singapore, but the accuracy depends on various factors and metrics. For example, at times, Weather forecast may indicate "Shower" in "Clementi" town area, it would be nice to know in near real-time if rain has arrived in "Clementi" or the intensity. Rainfall readings from the rain gauge deployed across Singapore, measured rain falling rate based on a 0.2mm gauge over 5 mins interval. It is able to provide near real-time insight on intensity of the rain in that area within a 5mins interval. TB1 (Tipping Bucket) Rain Gauge limitation is however not able to measure light rain/drizzle (<0.2mm per 5 min interval or commonly known as light liquid precipitation).

By correlating the 2, it provides a greater insight for the user to know the weather conditions so that they can plan ahead. This type of data might be useful for avid sports users or any users that likes to beat the weather forecast. The main objective is to provide more accurate informations for user <b>to make informed decision</b>.

The 2 datasets measured in different station locations across Singapore, will need mapping in order to provide meaningful correlation. For example, Weather forecast station in "Clementi" town is in location with latitude 1.315, longitude:103.76 and according to Google-map is near AYE highway exit into Clementi Avenue 6. The rain gauge is located in latitude:1.3337, longitude:103.7768, from Google-map is in Ngee Ann Polytechnic along Clementi Road. Observing within Clementi town area over a period of 7 days, with 3 raining condition, correlating both datasets provided an accurate situation in the area with a high degree of confidence, 99.9% level in this case. Rainfall readings not only provided information if rain has arrived but the intensity of the rain and if rain has passed (minus light rain/drizzle). This data along with Weather Forecast provided additional information for any users who can be out in the open area to seek shelter for a period of time, delay or alter their journey plans. It is to be noted that the Weather Forecast does change and reflect the current weather situation, albeit not as accurate and with some delays.

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
```
clementi
The weather forecast for Clementi is Partly Cloudy (Day)
Valid Start Time: 2019-11-22T12:30:00+08:00
Valid End Time: 2019-11-22T14:30:00+08:00
Rain Gauge Station -  Clementi Road - Current Measurement - mm/5min: 0
2019-11-22 12:55:00 +0800 +08
```
From the sample output, Weather Forecast for Clementi is shown as "Showers" at 12:30H, but the Rain Gauge reading does not measure any rain. The Weather Forecast is subsequently updated to reflect the rain has passed and is no longer raining in Clementi Town.

In addition, the ingested Rain Fall readings of all stations can be printed out by using the keyword "Rainfall" or "rainfall". This can be used by the User to have an idea if a nearby town or station is experiencing rain, there could be a higher chance of rain, depending on wind conditions.

```
rainfall
Admiralty Road West - mm: 0
Alexandra Road - mm: 0
Ang Mo Kio Avenue 5 - mm: 0
Banyan Road - mm: 0
Bukit Panjang Road - mm: 0
Bukit Timah Road - mm: 0
Chai Chee Street - mm: 0
Choa Chu Kang Avenue 4 - mm: 0.6
Choa Chu Kang Road - mm: 0.4
Clementi Road - mm: 0
East Coast Parkway - mm: 0
GEYLANG EAST CENTRAL - mm: 0
Handy Road - mm: 0
Holland Road - mm: 0
Jurong Pier Road - mm: 0.2
Kent Ridge Road - mm: 0
Kim Chuan Road - mm: 0
Kranji Way - mm: 8
Lornie Road - mm: 0
Mandai Lake Road - mm: 0.4
Marine Parade Road - mm: 0
Nanyang Avenue - mm: 0.4
Nicoll Highway - mm: 0
Old Choa Chu Kang Road - mm: 1.2
Old Toh Tuck Road - mm: 0
Pasir Ris Drive 12 - mm: 0.6
Pasir Ris Street 51 - mm: 2.8
Poole Road - mm: 0
Pulau Ubin - mm: 0
Punggol Central - mm: 0.2
Scotts Road - mm: 0
Seletar Aerospace View - mm: 0
Sembawang Road - mm: 0.2
Sentosa - mm: 0
Sime Road - mm: 0
Simei Avenue - mm: 0
Somerset Road - mm: 0
South Buona Vista Road - mm: 0
Toa Payoh North - mm: 0
Towner Road - mm: 0
Tuas Road - mm: 0
Tuas South Avenue 3 - mm: 0
Tuas West Road - mm: 0
Upper Changi Road North - mm: 5
Upper Peirce Reservoir Park - mm: 1.2
Upper Serangoon Road - mm: 0
Upper Thomson Road - mm: 0
West Coast Highway - mm: 0
Woodlands Avenue 9 - mm: 1.8
Woodlands Road - mm: 5.6
Yishun Avenue 5 - mm: 0
2019-11-24 14:45:00 +0800 +08
```
In the above example, Woodlands Road experience heavy rain (intensity of 5.6mm measured in 5mins interval) and Woodlands Ave 9 is having a rain fall of 1.8mm at 14:45hour. Output of "All" Weather Forecast to correlate

```
The weather for All of Singapore
Valid Start Time: 2019-11-24T14:30:00+08:00
Valid End Time: 2019-11-24T16:30:00+08:00
Ang Mo Kio Thundery Showers
Bedok Thundery Showers
Bishan Thundery Showers
Boon Lay Showers
Bukit Batok Light Showers
Bukit Merah Light Showers
Bukit Panjang Thundery Showers
Bukit Timah Light Showers
Central Water Catchment Thundery Showers
Changi Heavy Thundery Showers with Gusty Winds
Choa Chu Kang Thundery Showers
Clementi Light Showers
City Showers
Geylang Thundery Showers
Hougang Thundery Showers
Jalan Bahar Showers
Jurong East Showers
Jurong Island Showers
Jurong West Showers
Kallang Showers
Lim Chu Kang Heavy Thundery Showers with Gusty Winds
Mandai Heavy Thundery Showers with Gusty Winds
Marine Parade Thundery Showers
Novena Light Showers
Pasir Ris Heavy Thundery Showers with Gusty Winds
Paya Lebar Thundery Showers
Pioneer Showers
Pulau Tekong Heavy Thundery Showers with Gusty Winds
Pulau Ubin Heavy Thundery Showers with Gusty Winds
Punggol Heavy Thundery Showers with Gusty Winds
Queenstown Light Showers
Seletar Heavy Thundery Showers with Gusty Winds
Sembawang Heavy Thundery Showers with Gusty Winds
Sengkang Thundery Showers
Sentosa Light Showers
Serangoon Thundery Showers
Southern Islands Light Showers
Sungei Kadut Heavy Thundery Showers with Gusty Winds
Tampines Thundery Showers
Tanglin Showers
Tengah Heavy Thundery Showers with Gusty Winds
Toa Payoh Thundery Showers
Tuas Showers
Western Islands Light Showers
Western Water Catchment Heavy Thundery Showers with Gusty Winds
Woodlands Heavy Thundery Showers with Gusty Winds
Yishun Heavy Thundery Showers with Gusty Winds
```
At times, you might want to know the Rain situation a distance from where you are. It is possible to correlate Rain fall measurement seen in all Rain Gauge Stations within the radius distance of the query town area. In the example below, all Rain Gauge measurement within a 7.5km radius from "Clementi" Town Area are shown. At this time, you can manually edit and add in the additional available station's readings you will like on the stations.json file. Using Latitude and Longitude to calculate stations distance within 7.5km could be a considered option or you can verify this with Google-Map. Having an external json file option provides greater flexibility. The auto calculation could be baked into the code in the future if this proves useful after careful monitoring, using Haversine Distance calculations.

```
The weather forecast for Clementi is Heavy Thundery Showers with Gusty Winds
Forecast Time: 2019-12-02T14:39:00+08:00

Rain Gauge Station -  Clementi Road - Current Measurement - mm/5min: 0.2
Rain Gauge Station -  Kent Ridge Road - Current Measurement - mm/5min: 1.2
Rain Gauge Station -  Holland Road - Current Measurement - mm/5min: 3
Rain Gauge Station -  Bukit Timah Road - Current Measurement - mm/5min: 4.2
Rain Gauge Station -  Old Toh Tuck Road - Current Measurement - mm/5min: 0
Rain Gauge Station -  South Buona Vista Road - Current Measurement - mm/5min: 2
2019-12-02 14:55:00 +0800 +08
```
Release 2.0, make use of Haversine Formulae to calculate distance based on Coordinates (Latitutde and Longitude). This workaround town areas to station mapping where the author does not have enough local knowledge and may be hard to pinpoint with Google-Map. With a Mathematical method using Haversine, it can compare and display all Rain Gauge Stations within the user defined radius distance. All Rain Gauge Stations within distance can be correlated and displayed.

Dataset used for correlations:
1. Mapping for correlation - stations.json
2. Weather Forecast station locations 
3. TB1 Rain Gauge station locations

The software is robust and tested, however the accuracy of the correlation still needs to be work on for some locations.
You are welcome to download a copy of the software <a href="https://github.com/maxng07/SG_Weather_GO/releases/tag/2.0.1"> here </a> and provide any feedback to the author. If you like to help out or provide inputs to the correlation of the Weather Station and Rain Gauge Station, please contact the author.

References:-
1. <a href="https://www.nea.gov.sg/corporate-functions/resources/facts-figures/mq">NEA/MSS Near real time data </a>
2. <a href="http://www.weather.gov.sg/faq/"> Meterological Service Singapore </a>
3. <a href="http://www.emesystems.com/davis/documents/D7852%20Installation.pdf">Tipping Bucket Rain Gauge</a> and <a href="https://en.wikipedia.org/wiki/Rain_gauge"> Rain Gauge Wiki </a>
4. <a href="https://kids.frontiersin.org/article/10.3389/frym.2018.00038"> How to measure Rain Fall </a>
5. Google Map
6. <a href="https://data.gov.sg/dataset/realtime-weather-readings?resource_id=8bd37e06-cdd7-4ca4-9ad8-5754eb70a33d"> Near Real Time Rainfall Across Singapore </a> dataset
7. <a href="https://data.gov.sg/dataset/weather-forecast"> 2 Hour Weather Forecast </a> dataset
