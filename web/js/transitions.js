// Page transition handler
function smoothTransition(url) {
    // Add transitioning class to body
    document.body.classList.add('page-transitioning');

    // Wait for animation to complete, then navigate
    setTimeout(() => {
        window.location.href = url;
    }, 300); // Match the warpOut animation duration
}

// Intercept all artist card clicks for smooth transitions
document.addEventListener('DOMContentLoaded', () => {
    // Get all artist links
    const artistLinks = document.querySelectorAll('#myUL li a');

    artistLinks.forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const url = link.getAttribute('href');
            smoothTransition(url);
        });
    });

    // Also intercept back button
    const backButton = document.querySelector('a[href="/"]');
    if (backButton) {
        backButton.addEventListener('click', (e) => {
            e.preventDefault();
            smoothTransition('/');
        });
    }

    const backToArtists = document.querySelector('a[href="/artistes"]');
    if (backToArtists) {
        backToArtists.addEventListener('click', (e) => {
            e.preventDefault();
            smoothTransition('/artistes');
        });
    }

    // Intercept music links (Spotify/Deezer) - but let them open in new tab
    const musicLinks = document.querySelectorAll('.music-link');
    musicLinks.forEach(link => {
        // Don't intercept these - they open in new tabs naturally
        // But we can add a cool click effect
        link.addEventListener('click', () => {
            link.style.transform = 'scale(0.95)';
            setTimeout(() => {
                link.style.transform = '';
            }, 150);
        });
    });
});

// Make smoothTransition available globally
window.smoothTransition = smoothTransition;