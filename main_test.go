package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{

		{
			name: "0 k",
			args: args{
				lat1: 30.014829754521358,
				lon1: -82.74469823223264,
				lat2: 30.014829754521358,
				lon2: -82.74469823223264,
			},
			want: 0,
		},
		{
			name: "9 k",
			args: args{
				lat1: 30.007717,
				lon1: -82.644516,
				lat2: 30.014829754521358,
				lon2: -82.74469823223264,
			},
			want: 9.678584332152075,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistance(tt.args.lat1, tt.args.lon1, tt.args.lat2, tt.args.lon2); got != tt.want {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWithinRadius(t *testing.T) {

	type args struct {
		SensorLatitude  float64
		SensorLongitude float64
		StormLatitude   float64
		StormLongitude  float64
		radius          float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "In Radius",
			args: args{
				SensorLatitude:  30.007717,
				SensorLongitude: -82.644516,
				StormLatitude:   30.014829754521358,
				StormLongitude:  -82.74469823223264,
				radius:          20.0,
			},
			want: true,
		},
		{
			name: "Out Radius",
			args: args{
				SensorLatitude:  61.693921,
				SensorLongitude: -29.555985,
				StormLatitude:   30.014829754521358,
				StormLongitude:  -82.74469823223264,
				radius:          20.0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWithinRadius(tt.args.SensorLatitude, tt.args.SensorLongitude, tt.args.StormLatitude, tt.args.StormLongitude, tt.args.radius); got != tt.want {
				t.Errorf("IsWithinRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_degreesToRadians(t *testing.T) {
	type args struct {
		degrees float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 45",
			args: args{degrees: 45.0},
			want: 0.7853981633974483,
		},
		{
			name: "Test 90",
			args: args{degrees: 90.0},
			want: 1.5707963267948966,
		},
		{
			name: "Test 30",
			args: args{degrees: 30.0},
			want: 0.5235987755982988,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := degreesToRadians(tt.args.degrees); got != tt.want {
				t.Errorf("degreesToRadians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOverlabsHandler(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name: "ValidInput",
			requestBody: `{
    "sensor": [
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
        },
        {
            "id": "joels",
            "guid": "e034d175-23c4-4dd9-aa42-c57ed30c4a79",
            "isActive": true,
            "registered": "2016-04-11T03:01:47 +06:00",
            "latitude": 30.014829754521358,
            "longitude": -82.74469823223264,
            "range": 10
        },
        {
            "id": "qkis",
            "guid": "qkis-23c4-4dd9-aa42-c57ed30c4a79",
            "isActive": true,
            "registered": "2016-04-11T03:01:47 +06:00",
            "latitude": 30.007717,
            "longitude": -82.644516,
            "range": 10
        }
    ],
    "storms": [
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
        },
        
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
        },
      
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
        },
        {
            "name": "JoelsStorm",
            "id": "111111111111111",
            "latitude": 30.014829754521358,
            "longitude": -82.74469823223264,
            "radius": 19.2287
        },
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
}`,
			expectedCode: http.StatusOK,
			expectedBody: `["Sensor ID: joels Overlapping Storm:JoelsStorm","Sensor ID: qkis Overlapping Storm:JoelsStorm"]`,
		},
		{
			name: "EmptyInput",
			requestBody: `{
				"Sensor": [],
				"Storms": []
			}`,
			expectedCode: http.StatusOK,
			expectedBody: "null",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP POST request with the test request body
			req, err := http.NewRequest("POST", "/find-overlaps", strings.NewReader(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder to capture the response from the handler
			rr := httptest.NewRecorder()

			// Call the findOverlabsHandler function with the test request and response recorder
			findOverlabsHandler(rr, req)

			// Check the HTTP status code
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}

			// Check the response body
			if body := strings.TrimSpace(rr.Body.String()); body != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", body, tt.expectedBody)
			}
		})
	}
}

func Test_checkSensorStormOverlap(t *testing.T) {
	type args struct {
		sensors []Sensor
		storms  []Storm
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			name: "Overlap",
			args: args{
				sensors: []Sensor{
					{
						ID:         "6479f6f5148d1c1febf67a09",
						GUID:       "Joels-0b70-4804-8fb8-41b27d1748ed",
						IsActive:   true,
						Registered: "2022-09-23T11:38:46 +06:00",
						Latitude:   30.014829754521358,
						Longitude:  -82.74469823223264,
						Range:      10,
					},
					{
						ID:         "6479f6f56ef6a7ef5e934c58",
						GUID:       "4fe5d522-22d2-4d1d-840a-3f1354688bf9",
						IsActive:   true,
						Registered: "2016-04-02T05:59:13 +06:00",
						Latitude:   76.549056,
						Longitude:  132.095063,
						Range:      10,
					},
				},
				storms: []Storm{
					{
						Name:      "Bad Storm",
						ID:        "6479f8f8decabd4378a814c0",
						Latitude:  30.014829754521358,
						Longitude: -82.74469823223264,
						Radius:    18.597,
					},
					{
						Name:      "Harry",
						ID:        "6479f8f87d117ec441470286",
						Latitude:  -80.36302,
						Longitude: 122.946741,
						Radius:    49.2088,
					},
				},
			},
			want: []string{"Sensor ID: 6479f6f5148d1c1febf67a09 Overlapping Storm:Bad Storm"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSensorStormOverlap(tt.args.sensors, tt.args.storms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkSensorStormOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
