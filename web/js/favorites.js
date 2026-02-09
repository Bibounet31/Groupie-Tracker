
// FAVORITES MANAGEMENT SYSTEM

// Get favorites from localStorage
function getFavorites() {
    const favs = localStorage.getItem('favoriteArtists');
    return favs ? JSON.parse(favs) : [];
}

// Save favorites to localStorage
function saveFavorites(favorites) {
    localStorage.setItem('favoriteArtists', JSON.stringify(favorites));
}

// Check if an artist is in favorites
function isFavorite(artistName) {
    const favorites = getFavorites();
    return favorites.includes(artistName);
}

// Toggle favorite status
function toggleFavorite(artistName) {
    let favorites = getFavorites();

    if (favorites.includes(artistName)) {
        // Remove from favorites
        favorites = favorites.filter(name => name !== artistName);
        console.log(`❌ Removed "${artistName}" from favorites`);
    } else {
        // Add to favorites
        favorites.push(artistName);
        console.log(`⭐ Added "${artistName}" to favorites`);
    }

    saveFavorites(favorites);
    return favorites.includes(artistName);
}

// Update star icon appearance
function updateStarIcon(star, isFav) {
    if (isFav) {
        star.innerHTML = '⭐'; // Filled star
        star.classList.add('favorited');
        star.title = 'Remove from favorites';
    } else {
        star.innerHTML = '☆'; // Empty star
        star.classList.remove('favorited');
        star.title = 'Add to favorites';
    }
}


// INITIALIZE FAVORITES ON ALBUMS PAGE

document.addEventListener('DOMContentLoaded', () => {
    const artistsList = document.getElementById('myUL');

    // Add stars to each artist on the albums page
    if (artistsList) {
        const artistItems = artistsList.querySelectorAll('li');

        artistItems.forEach(item => {
            const link = item.querySelector('a');
            if (!link) return;

            const artistName = link.textContent.trim();

            // Create star icon
            const star = document.createElement('span');
            star.className = 'favorite-star';
            star.style.cursor = 'pointer';
            star.style.marginLeft = '8px';
            star.style.fontSize = '20px';
            star.style.userSelect = 'none';

            // Set initial state
            updateStarIcon(star, isFavorite(artistName));

            // Add click handler
            star.addEventListener('click', (e) => {
                e.preventDefault();
                e.stopPropagation();

                const nowFavorited = toggleFavorite(artistName);
                updateStarIcon(star, nowFavorited);
            });

            // Add star after the link
            link.parentNode.insertBefore(star, link.nextSibling);
        });

        console.log('✅ Favorites system initialized on albums page');
    }


    // INITIALIZE FAVORITE ON DETAILS PAGE
    const detailsTitle = document.querySelector('body > h1');

    if (detailsTitle && !artistsList) { // Only on details page
        const artistName = detailsTitle.textContent.trim();

        // Create star icon for details page
        const star = document.createElement('span');
        star.className = 'favorite-star detail-star';
        star.style.cursor = 'pointer';
        star.style.marginLeft = '15px';
        star.style.fontSize = '28px';
        star.style.userSelect = 'none';
        star.style.verticalAlign = 'middle';

        // Set initial state
        updateStarIcon(star, isFavorite(artistName));

        // Add click handler
        star.addEventListener('click', (e) => {
            e.preventDefault();

            const nowFavorited = toggleFavorite(artistName);
            updateStarIcon(star, nowFavorited);
        });

        // Add star to title
        detailsTitle.appendChild(star);

        console.log('✅ Favorites system initialized on details page');
    }
});


// FILTER FAVORITES FUNCTION
function showOnlyFavorites() {
    const artistsList = document.getElementById('myUL');
    if (!artistsList) return;

    const favorites = getFavorites();
    const artistItems = artistsList.querySelectorAll('li');

    if (favorites.length === 0) {
        alert('You have no favorite artists yet! Click the ⭐ to add some.');
        return;
    }

    let visibleCount = 0;

    artistItems.forEach(item => {
        const link = item.querySelector('a');
        if (!link) return;

        const artistName = link.textContent.trim();

        if (favorites.includes(artistName)) {
            item.style.display = 'flex';
            visibleCount++;
        } else {
            item.style.display = 'none';
        }
    });

    console.log(`📌 Showing ${visibleCount} favorite artists`);

    // Show message if no favorites visible
    const noResultsMsg = document.getElementById('no-results-message');
    if (noResultsMsg && visibleCount === 0) {
        noResultsMsg.querySelector('p').textContent = 'Your favorites are currently filtered out. Clear filters to see them.';
        noResultsMsg.style.display = 'block';
    }
}


// SHOW ALL ARTISTS FUNCTION

function showAllArtists() {
    const artistsList = document.getElementById('myUL');
    if (!artistsList) return;

    const artistItems = artistsList.querySelectorAll('li');

    artistItems.forEach(item => {
        item.style.display = 'flex';
    });

    // Hide no results message
    const noResultsMsg = document.getElementById('no-results-message');
    if (noResultsMsg) {
        noResultsMsg.style.display = 'none';
    }

    console.log('📋 Showing all artists');
}

// Make functions available globally
window.showOnlyFavorites = showOnlyFavorites;
window.showAllArtists = showAllArtists;