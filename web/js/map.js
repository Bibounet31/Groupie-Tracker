// Check if Leaflet is loaded and map element exists
if (typeof L !== 'undefined') {
    const mapElement = document.getElementById('map');

    if (mapElement) {
        // Initialize map
        const map = L.map('map').setView([20, 0], 2);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            maxZoom: 19,
            attribution: '© OpenStreetMap'
        }).addTo(map);

        // Get concert data from the script tag
        const concertDataElement = document.getElementById('concert-data');
        let concerts = {};

        if (concertDataElement) {
            try {
                concerts = JSON.parse(concertDataElement.textContent);
                console.log('Concert data loaded:', concerts);
            } catch (error) {
                console.error('Error parsing concert data:', error);
            }
        }

        // Geocode and add markers
        async function addMarkers() {
            for (const [location, count] of Object.entries(concerts)) {
                try {
                    // Clean location name (remove underscores, format properly)
                    const cleanLocation = location.replace(/_/g, ' ').replace(/-/g, ', ');

                    console.log(`Geocoding: ${cleanLocation}`);

                    // Use Nominatim API for geocoding
                    const response = await fetch(
                        `https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(cleanLocation)}&limit=1`
                    );
                    const data = await response.json();

                    if (data.length > 0) {
                        const { lat, lon, display_name } = data[0];

                        console.log(`Found: ${display_name} at ${lat}, ${lon}`);

                        // Add marker
                        L.marker([lat, lon])
                            .addTo(map)
                            .bindPopup(`<b>${display_name}</b><br>${count} concert${count > 1 ? 's' : ''}`);
                    } else {
                        console.log(`No results for: ${cleanLocation}`);
                    }

                    // Rate limiting - wait between requests
                    await new Promise(resolve => setTimeout(resolve, 1000));
                } catch (error) {
                    console.error(`Error geocoding ${location}:`, error);
                }
            }
        }

        if (Object.keys(concerts).length > 0) {
            addMarkers();
        } else {
            console.log('No concert data available');
        }
    }
}