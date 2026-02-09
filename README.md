# Groupie Trackers

## Project Description

Groupie Trackers is a web application built with Go that displays musical artists and their concert locations using the Groupie Trackers API. Users can search, filter artists, and view concert locations on an interactive map.

## Features

* Display all artists
* Search by artist name or band member
* Filter by creation year and number of members
* Artist detail page (members, creation year, concerts)
* Interactive concert map (Leaflet.js)
* Responsive design
* Custom 404 error page

## How to Run the Project

### Requirements

* Go 1.16 or higher
* Internet connection

### Steps

1. Clone the repository:
```bash
git clone <your-repository>
cd groupie-tracker
```

2. Run the application:
```bash
go run cmd/main.go
```

3. Open your browser at:
```
http://localhost:8080
```

## Main Routes

| Route | Method | Description |
|-------|--------|-------------|
| `/` | GET | Home page |
| `/albums` | GET | List of artists |
| `/artists` | GET | Filtered search results |
| `/details/{name}` | GET | Artist details |
| `/search` | GET | Search autocomplete |
| `/css/*` | GET | Static CSS |
| `/js/*` | GET | Static JS |
| `/img/*` | GET | Static images |

## Technologies Used

### Backend
* Go
* net/http
* html/template

### Frontend
* HTML
* CSS
* JavaScript
* Leaflet.js

### APIs
* Groupie Trackers API
* Nominatim (OpenStreetMap)

## Notes

* Internet access is required for API data and maps
* Geocoding requests are limited to 1 request per second

---

© 2025 – Educational Project
