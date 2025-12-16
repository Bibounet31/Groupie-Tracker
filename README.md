# 🎵 Groupie Tracker

A web application built in Go that displays, searches, and filters music artists and their concert information using the GroupieTrackers API.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

---

## 📋 Table of Contents

- [Features](#-features)
- [Demo](#-demo)
- [Project Structure](#-project-structure)
- [Installation](#-installation)
- [Usage](#-usage)
- [API Documentation](#-api-documentation)
- [How It Works](#-how-it-works)
- [Contributing](#-contributing)
- [Troubleshooting](#-troubleshooting)
- [License](#-license)

---

## ✨ Features

- 🎨 **Beautiful UI** - Glassmorphism design with smooth animations
- 🔍 **Smart Search** - Search by artist name or band member with autocomplete
- 🎯 **Advanced Filters** - Filter by creation year and member count
- 🗺️ **Interactive Map** - View concert locations on an interactive world map
- 📱 **Responsive Design** - Works on desktop and mobile devices
- 🎧 **Music Integration** - Direct links to Spotify and Deezer
- ⚡ **Fast Performance** - No database needed, all data in memory
- 🎭 **Smooth Transitions** - Page animations for a polished experience

---

## 🎬 Demo

```bash
# Start the server
go run cmd/main.go

# Open in browser
http://localhost:8080
```

### Screenshots

**Home Page**
```
┌─────────────────────────────────┐
│      Groupie Trackers 🎵        │
│                                 │
│         [Artist]                │
│                                 │
│   Search and discover music     │
│   artists and their concerts    │
└─────────────────────────────────┘
```

**Artist List with Search**
```
┌─────────────────────────────────┐
│  [Artist ▼] [Search...____]     │
│  Year: [1990] - [2020]          │
│                                 │
│  🎸 Queen        🎤 Beatles     │
│  🎹 Metallica    🎺 Jazz Band   │
└─────────────────────────────────┘
```

**Artist Details with Map**
```
┌─────────────────────────────────┐
│  Queen                    [img] │
│  Members: Freddie Mercury...    │
│  Created: 1970                  │
│                                 │
│  🗺️ Concert Map                 │
│  ┌─────────────────────────┐   │
│  │    🔴 London (5)        │   │
│  │       🔴 Paris (3)      │   │
│  └─────────────────────────┘   │
│                                 │
│  📍 London, UK                  │
│    • 10-01-2024                 │
│    • 11-01-2024                 │
└─────────────────────────────────┘
```

---

## 🛠️ Tech Stack

### Backend
- **[Go](https://golang.org/)** - Main programming language
- **net/http** - HTTP server and routing
- **html/template** - Template rendering
- **encoding/json** - JSON parsing

### Frontend
- **HTML5** - Page structure
- **CSS3** - Styling with glassmorphism effects
- **Vanilla JavaScript** - No frameworks needed
- **[Leaflet.js](https://leafletjs.com/)** - Interactive maps

### External APIs
- **[GroupieTrackers API](https://groupietrackers.herokuapp.com/api)** - Artist data
- **[Nominatim](https://nominatim.openstreetmap.org/)** - Geocoding service
- **[OpenStreetMap](https://www.openstreetmap.org/)** - Map tiles

---

## 📁 Project Structure

```
groupie/
│
├── cmd/
│   └── main.go                 # Entry point - starts server
│
├── server/
│   ├── server.go              # Server setup & routing
│   └── handlers.go            # Request handlers & logic
│
└── web/
    ├── html/
    │   ├── index.html         # Home page
    │   ├── artistes.html      # Artist list page
    │   ├── details.html       # Artist detail page
    │   └── 404.html           # Error page
    │
    ├── css/
    │   ├── style.css          # Home page styles
    │   ├── artistes.css       # Artist list styles
    │   └── details.css        # Details page styles
    │
    ├── js/
    │   ├── script.js          # Search & autocomplete
    │   ├── map.js             # Interactive map
    │   └── transitions.js     # Page animations
    │
    └── img/
        └── miku.jpg           # Background image
```

---

## 🚀 Installation

### Prerequisites

- Go 1.21 or higher
- Internet connection (for API access)

### Steps

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/groupie-tracker.git
cd groupie-tracker
```

2. **Verify Go installation**
```bash
go version
```

3. **Run the application**
```bash
go run cmd/main.go
```

4. **Open in browser**
```
http://localhost:8080
```

### Building for Production

```bash
# Build executable
go build -o groupie-tracker cmd/main.go

# Run executable
./groupie-tracker
```

---

## 💻 Usage

### Browsing Artists

1. Click **"Artist"** on the home page
2. Browse through the list of artists
3. Click any artist card to see details

### Searching

**By Artist Name:**
```
1. Select "Artist" from dropdown
2. Type artist name (e.g., "Queen")
3. Click suggestion or press Enter
```

**By Band Member:**
```
1. Select "Member" from dropdown
2. Type member name (e.g., "Freddie")
3. Click suggestion to see all their bands
```

### Filtering

**Year Range:**
```
Year min: 1970
Year max: 1990
→ Shows artists created between 1970-1990
```

**Member Count:**
```
Number of members: 4
→ Shows bands with exactly 4 members

Number of members: 5+
→ Shows bands with 5 or more members
```

**Combine Filters:**
```
Year min: 1990
Year max: 2000
Members: 4
→ Shows 4-member bands created in the 90s
```

### Viewing Concert Locations

1. Click any artist to view details
2. Scroll to **"Concert Locations"** section
3. Interactive map shows all concert venues
4. Click markers to see venue details
5. List below shows dates for each location

---

## 📡 API Documentation

### Internal Endpoints

#### `GET /`
**Home page**
- Returns: HTML home page

#### `GET /artistes`
**Artist list page**
- Returns: HTML with all artists

#### `GET /details/{name}`
**Artist details page**
- Parameters:
  - `name` (path) - Artist name (URL encoded)
- Returns: HTML with artist details and concert map

#### `GET /search`
**Autocomplete search**
- Parameters:
  - `query` (string) - Search term
  - `type` (string) - "artist" or "member"
- Returns: JSON array of matching names
```json
["Queen", "Queens of the Stone Age"]
```

#### `GET /artists`
**Filtered artist search**
- Parameters:
  - `query` (string, optional) - Artist name search
  - `member` (string, optional) - Member name search
  - `year_min` (int, optional) - Minimum creation year
  - `year_max` (int, optional) - Maximum creation year
  - `members_count` (string, optional) - Member count (e.g., "4" or "5+")
- Returns: HTML with filtered artists

### External APIs Used

#### GroupieTrackers API

**Get all artists:**
```http
GET https://groupietrackers.herokuapp.com/api/artists
```

Response:
```json
[
  {
    "id": 1,
    "name": "Queen",
    "image": "https://...",
    "members": ["Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon"],
    "creationDate": 1970,
    "locations": "https://.../api/locations/1",
    "concertDates": "https://.../api/dates/1",
    "relations": "https://.../api/relation/1"
  }
]
```

**Get concert locations:**
```http
GET https://groupietrackers.herokuapp.com/api/relation/{id}
```

Response:
```json
{
  "index": 1,
  "datesLocations": {
    "london-uk": ["10-01-2024", "11-01-2024"],
    "paris-france": ["15-01-2024"]
  }
}
```

#### Nominatim Geocoding API

**Geocode location:**
```http
GET https://nominatim.openstreetmap.org/search?format=json&q=Paris,France&limit=1
```

Response:
```json
[
  {
    "lat": "48.8566",
    "lon": "2.3522",
    "display_name": "Paris, Île-de-France, France"
  }
]
```

⚠️ **Important:** Nominatim requires a User-Agent header and has a 1 request/second rate limit.

---

## 🔧 How It Works

### Architecture Overview

```
┌─────────────┐
│   Browser   │
└──────┬──────┘
       │ HTTP Request
       ↓
┌─────────────────────────────────┐
│      Go Server (Port 8080)      │
│  ┌──────────────────────────┐  │
│  │  Router (server.go)      │  │
│  │  • Routes                │  │
│  │  • Static files          │  │
│  └────────┬─────────────────┘  │
│           ↓                     │
│  ┌──────────────────────────┐  │
│  │  Handlers (handlers.go)  │  │
│  │  • Process requests      │  │
│  │  • Fetch API data        │  │
│  │  • Render templates      │  │
│  └──────────────────────────┘  │
└─────────────────────────────────┘
       │ HTTP GET
       ↓
┌─────────────────────────────────┐
│     GroupieTrackers API         │
│  (External - Heroku)            │
└─────────────────────────────────┘
```

### Request Flow Example

**User visits artist details:**

```
1. User clicks "Queen" card
   ↓
2. Browser: GET /details/Queen
   ↓
3. Server (DetailsHandler):
   • Extract "Queen" from URL
   • Find artist in memory (AllArtists)
   • Fetch concert data from API
   • Combine data
   ↓
4. Render details.html template
   • Inject artist info
   • Embed concert JSON
   ↓
5. Browser receives HTML
   ↓
6. map.js executes:
   • Parse embedded JSON
   • Geocode each location (Nominatim)
   • Add markers to Leaflet map
   ↓
7. User sees interactive page with map
```

### Data Flow

```
Startup:
┌──────────────┐
│  main.go     │
│  • Fetch API │ ──GET──> GroupieTrackers API
│  • Store data│          (All artists)
│  • Start srv │
└──────────────┘
       ↓
   Server Ready
   
Runtime:
┌──────────────┐
│  User Request│
└──────┬───────┘
       ↓
┌──────────────┐
│  Handler     │
│  • Search    │ ──→ Filter AllArtists in memory
│  • Filter    │
│  • Details   │ ──GET──> GroupieTrackers API
└──────┬───────┘          (Concert data)
       ↓
┌──────────────┐
│  Template    │
│  • Render    │
└──────┬───────┘
       ↓
┌──────────────┐
│  Browser     │
│  • Display   │
│  • map.js    │ ──GET──> Nominatim API
└──────────────┘          (Coordinates)
```

### Key Features Explained

#### 1. Smart Search with Autocomplete

```javascript
// User types "qu"
input.addEventListener("input", async () => {
    // Fetch matching artists
    const res = await fetch(`/search?query=qu&type=artist`);
    const results = await res.json();
    // ["Queen", "Queens of the Stone Age"]
    
    // Display dropdown
    results.forEach(name => {
        // Click → navigate to artist page
    });
});
```

#### 2. Multi-Filter System

```go
// Server combines all filters
for _, artist := range AllArtists {
    match := true
    
    // Name search
    if query != "" && !strings.Contains(artist.Name, query) {
        match = false
    }
    
    // Year range
    if yearMin != 0 && artist.CreationDate < yearMin {
        match = false
    }
    
    // Member count
    if memberCount != len(artist.Members) {
        match = false
    }
    
    if match {
        results = append(results, artist)
    }
}
```

#### 3. Interactive Map

```javascript
// Initialize map
const map = L.map('map').setView([20, 0], 2);

// Add OpenStreetMap tiles
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png').addTo(map);

// Geocode and add markers
for (const [location, count] of Object.entries(concerts)) {
    // "Paris-France" → "Paris, France"
    const cleanLocation = location.replace(/_/g, ' ').replace(/-/g, ', ');
    
    // Get coordinates
    const response = await fetch(
        `https://nominatim.openstreetmap.org/search?format=json&q=${cleanLocation}`
    );
    const {lat, lon} = await response.json()[0];
    
    // Add marker
    L.marker([lat, lon])
        .addTo(map)
        .bindPopup(`${location} - ${count} concerts`);
}
```

---

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

### Reporting Bugs

1. Check if the bug has already been reported
2. Open a new issue with:
   - Clear description
   - Steps to reproduce
   - Expected vs actual behavior
   - Screenshots if applicable

### Suggesting Features

Open an issue with:
- Feature description
- Use case
- Proposed implementation (optional)

### Pull Requests

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices
- Comment complex logic
- Test your changes
- Keep commits atomic
- Update documentation

---

## 🐛 Troubleshooting

### Map Not Loading

**Symptom:** Gray box instead of map, no markers

**Solutions:**

1. **Add User-Agent header to Nominatim requests:**
```javascript
const response = await fetch(url, {
    headers: {
        'User-Agent': 'GroupieTracker/1.0'
    }
});
```

2. **Check browser console** (F12 → Console) for errors

3. **Verify Leaflet.js loaded:**
```javascript
console.log(typeof L); // Should be "object"
```

4. **Test Nominatim directly:**
```
https://nominatim.openstreetmap.org/search?format=json&q=Paris&limit=1
```

### Artists Not Displaying

**Check:**

1. **API accessibility:**
```bash
curl https://groupietrackers.herokuapp.com/api/artists
```

2. **Server logs:** Look for errors during startup

3. **Browser Network tab:** Check if `/artistes` request succeeds

### Search Not Working

**Check:**

1. **JavaScript console** for errors
2. **Network tab:** Is `/search` endpoint responding?
3. **Test endpoint directly:**
```
http://localhost:8080/search?query=queen&type=artist
```

### Port Already in Use

**Error:** `bind: address already in use`

**Solution:**
```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>

# Or change port in server.go:
http.ListenAndServe(":8081", handler)
```

### Slow Map Loading

**Cause:** Sequential geocoding (1 request/second)

**Temporary fix:** Reduce number of concert locations in data

**Better fix:** Implement geocoding cache:
```go
var geoCache = make(map[string]Coordinates)
```

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👥 Authors

- **Your Name** - *Initial work* - [YourGitHub](https://github.com/Bibounet31)

---

## 🙏 Acknowledgments

- [GroupieTrackers](https://groupietrackers.herokuapp.com) for the API
- [OpenStreetMap](https://www.openstreetmap.org/) for map tiles
- [Leaflet.js](https://leafletjs.com/) for the mapping library
- [Nominatim](https://nominatim.openstreetmap.org/) for geocoding

---

## 📞 Support

If you have any questions or need help:

- 🐛 Issues: [GitHub Issues](https://github.com/yourusername/groupie-tracker/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/yourusername/groupie-tracker/discussions)

---

## 🗺️ Roadmap

- [ ] Add user authentication
- [ ] Implement favorites system
- [ ] Add concert date filtering
- [ ] Improve mobile responsiveness
- [ ] Add dark/light theme toggle
- [ ] Implement geocoding cache
- [ ] Add pagination for artist list
- [ ] Export concert data to calendar
- [ ] Add artist comparisons
- [ ] Implement SQL database

---

## 📊 Statistics

![Lines of Code](https://img.shields.io/badge/lines%20of%20code-2500+-blue)
![Files](https://img.shields.io/badge/files-13-green)
![Languages](https://img.shields.io/badge/languages-3-orange)

---

**Made with ❤️ and Go**
