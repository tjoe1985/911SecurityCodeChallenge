# Assumptions 
Assuming the code is designed to be a long-running service accessible via an API, there are a few key assumptions:

The service will be implemented as a simple API that can handle incoming requests. This will allow users to interact with the service as needed.

The number of sensors and storms is not fixed and may change dynamically. Therefore, the service will read the sensor and storm data as required, allowing for flexibility in the system.

Regarding including the time of overlaps, although the storm information doesn't provide a specific time reference, the assumption is that the "registered" field in the sensor data can be used as an event trigger. By utilizing this field, the code can be easily modified to include the time when the overlaps occurred. The implementation of this feature will be added to the code in the near future.
# Test coverage: 87.9% 
# Questions and answers
## The output for only active sensors
Was coded already on the code but would be nice to add the options for engineers to decide if the want to use it or not. 
## Is there any overlap between the sensor’s area and a given storm’s area?
Not on the provided samples but for testing purporses we added some on the test to check the program is running fine.
## How much does a storm overlap with a sensor?
Not coded on the challenge but we already had access to the storm's radius and sensor's range and are calculating the 
distance so another function could be implemente to make better use of the information. 
## Can the output contain the start/end time of a storm’s overlap, and the duration of that time?
The output can take anything that we put in to it. However we would have to add more information to the incoming request
or start tracking the storm movement. Bases on current information and understanding we are not able to add it. 
## How could the output indicate if the storm is moving towards the sensor or if it’s passing by on the edge of the sensor’s radius?
With the provided information and no more context not sure. But by adding more fields to the information coming like 
capturing lat and lon on different time intervals would allow us to tell direction.
## Grouping the next three questions 
## If the storm data is sent once per second or once per hour, how does that change the solution theoretically?
## If we have 1000 sensors, how does that change the solution theoretically?
## If the list of storms is between 0 and 1000, how does that change the solution theoretically?
The rate and amount of data will affect how we scale the application. Too fast or too much
we would want to implement some sort of concurrency as needed to be able to handle the rate and amount 
of data being processed. It would all depend on the hardware, and resources we have available and able to spend 
## If a storm’s radius is never larger than 30 kilometers, how does that change the solution?
Not sure how to answer this without more context at the moment with just math context the calculation would run the same. But i have a lot to learn.  
## What other data could the engineers provide that would help in determining if there is a correlation between the sensors and the storms?
With current context, time of storm occurrence, direction or second read of storm location, was there thunder or wind? 
anything that can show influence on the sensors other than location. A small info class as to how sensor works and what 
can mess them up? not sure reaching at this point.   
## How could this be adapted to send an alert in real-time to users/operators of these sensors?
We can integrate to any messaging app and relate permanent information as needed. would talk to the operators to see 
what would useful information is and then generate it as needed. 
# There is a Storm Coming
At GlobalTechSensors, a few engineers have noticed that the sensors seem to go haywire
whenever there is a large storm, but they don’t have any proof. They want you to make a new
service that they can send data for a few weeks, and then see if there is any correlation
between storms and sensors afterwards.
These engineers want to send a static list of sensor locations, their ranges, and whether they
are online. They also want to send a list of storms between 0 and 20 that will contain the center
of a storm and the approximate radius of the storm.
The engineers want to know if a storm overlaps with any of the sensors. For simplicity, if the
distance between the gps location of a sensor is within the radius of a storm, then it counts as
overlapping.
Create a Go service that accepts these inputs and then logs the output. The output should
contain the id of sensors that have overlapping storms, the names of the overlapping storms,
and the time that the overlap(s) occurred.
Questions the engineers might ask:
- The output for only active sensors
- Is there any overlap between the sensor’s area and a given storm’s area?
- How much does a storm overlap with a sensor?
- Can the output contain the start/end time of a storm’s overlap, and the duration of that
  time?
- How could the output indicate if the storm is moving towards the sensor or if it’s passing
  by on the edge of the sensor’s radius?
- If the storm data is sent once per second or once per hour, how does that change the
  solution theoretically?
- If we have 1000 sensors, how does that change the solution theoretically?
- If the list of storms is between 0 and 1000, how does that change the solution
  theoretically?
- If a storm’s radius is never larger than 30 kilometers, how does that change the solution?
- What other data could the engineers provide that would help in determining if there is a
  correlation between the sensors and the storms?
- How could this be adapted to send an alert in real-time to users/operators of these
  sensors?
  Please do not spend more than a couple hours on this solution. The questions the engineers
  might ask are not expected to be coded if there is not enough time. Assume range and radius
  are in kilometers. For any clarifying questions, do not hesitate to email.
  Example data
  This data is generated randomly. It is not guaranteed to have any overlaps between the storms
  and the sensors (active or inactive).
  Sensor list:
  [
  {
  "id": "6479f6f5148d1c1febf67a09",
  "guid": "747250c4-0b70-4804-8fb8-41b27d1748ed",
  "isActive": false,
  "registered": "2022-09-23T11:38:46 +06:00",
  "latitude": 19.432564,
  "longitude": 52.556912,
  "range": 10
  },
  {
  "id": "6479f6f56ef6a7ef5e934c58",
  "guid": "4fe5d522-22d2-4d1d-840a-3f1354688bf9",
  "isActive": false,
  "registered": "2016-04-02T05:59:13 +06:00",
  "latitude": 76.549056,
  "longitude": 132.095063,
  "range": 10
  },
  {
  "id": "6479f6f57d083ca1c0e7a304",
  "guid": "0019d4cd-1818-4451-a76f-e449a0842cff",
  "isActive": true,
  "registered": "2019-10-04T10:31:24 +06:00",
  "latitude": 74.888799,
  "longitude": 7.693088,
  "range": 10
  },
  {
  "id": "6479f6f5843109f23e3bb833",
  "guid": "9c5d6482-616d-4ed3-a43f-2169d8431994",
  "isActive": false,
  "registered": "2020-12-08T07:23:15 +07:00",
  "latitude": -61.846087,
  "longitude": -116.957072,
  "range": 10
  },
  {
  "id": "6479f6f534dcc4cb802db4d4",
  "guid": "522faa6c-0014-49b2-b71b-c28b4b60a6aa",
  "isActive": true,
  "registered": "2019-07-15T11:44:28 +06:00",
  "latitude": -45.048556,
  "longitude": -80.469664,
  "range": 10
  },
  {
  "id": "6479f6f59200306f3b125389",
  "guid": "e0346175-28c4-4dd9-a442-c27ed30c4a79",
  "isActive": false,
  "registered": "2016-04-11T03:01:47 +06:00",
  "latitude": 82.139298,
  "longitude": -136.910771,
  "range": 10
  }
  ]
  Storm Data
  Zero Storms:
  []
  Two Storms:
  [
  {
  "name": "Melissa",
  "id": "6479f982aa4b36aa1b97c538",
  "latitude": 16.756599,
  "longitude": -67.755892,
  "radius": 43.8447
  },
  {
  "name": "Tom",
  "id": "6479f9829cf990195bf82976",
  "latitude": -29.555985,
  "longitude": 61.693921,
  "radius": 19.2287
  }
  ]
  Five Storms:
  [
  {
  "name": "Dave",
  "id": "6479f8f89e052557dc4a1f76",
  "latitude": 39.983811,
  "longitude": 149.055087,
  "radius": 28.6745
  },
  {
  "name": "Melissa",
  "id": "6479f8f88eaaa3408675d03f",
  "latitude": 53.780909,
  "longitude": -152.684978,
  "radius": 28.7489
  },
  {
  "name": "Bill",
  "id": "6479f8f88f6e636370f537d6",
  "latitude": -37.048947,
  "longitude": -47.625284,
  "radius": 12.4276
  },
  {
  "name": "Olga",
  "id": "6479f8f8decabd4378a814c0",
  "latitude": -5.262222,
  "longitude": 27.932393,
  "radius": 18.597
  },
  {
  "name": "Harry",
  "id": "6479f8f87d117ec441470286",
  "latitude": -80.36302,
  "longitude": 122.946741,
  "radius": 49.2088
  }
