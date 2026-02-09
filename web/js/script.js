//SUGGESTION SCRIPT

// get elements for search functionality
const input = document.getElementById("myInput");           // Search input field
const suggestions = document.getElementById("suggestions"); // Dropdown list for suggestions
const searchType = document.getElementById("searchType");   // Dropdown to select "artist" or "member"

// Only initialize if all required elements are filled
if (input && suggestions && searchType) {
    // Triggers on every character typed in the search box
    input.addEventListener("input", async () => {
        const query = input.value.trim();  // Get search query, remove whitespace
        const type = searchType.value;     // Get selected search type ("artist" or "member")

        // If input is empty, hide suggestions and clear the list
        if (!query) {
            suggestions.innerHTML = "";
            suggestions.style.display = "none";
            return;
        }

        try {
            // Fetch matching results from the server
            const res = await fetch(`/search?query=${encodeURIComponent(query)}&type=${type}`);
            const results = await res.json();  // Parse JSON response (array of names)

            // Clear previous suggestions
            suggestions.innerHTML = "";

            // Show the suggestions dropdown (even for "no results" message)
            suggestions.style.display = "block";

            // If no results found, show "No results" message in dropdown
            if (!results || results.length === 0) {
                const li = document.createElement("li");
                li.textContent = "No results found";
                li.classList.add("no-results"); // Add class for styling
                li.style.cursor = "default";    // Change cursor to indicate non-clickable
                suggestions.appendChild(li);
                return;
            }

            // Create a list item for each result
            results.forEach(name => {
                const li = document.createElement("li");
                li.textContent = name;  // Display the artist or member name

                // Handle click on a suggestion
                li.onclick = () => {
                    input.value = name;              // Fill input with selected name
                    suggestions.innerHTML = "";       // Clear suggestions list
                    suggestions.style.display = "none"; // Hide dropdown
                    window.location.href = `/artists?query=${encodeURIComponent(query)}`; //search in artist list
                };

                // Add the suggestion item to the dropdown
                suggestions.appendChild(li);
            });
        } catch (error) {
            console.error("Search error:", error);
            suggestions.style.display = "none";
        }
    });

    // SEARCH ON ENTER KEY
    // Handle Enter key press to perform search
    input.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            const query = input.value.trim();  // Get trimmed query
            const type = searchType.value;     // Get search type

            // Don't search if query is empty
            if (!query) return;

            // Hide suggestions dropdown before redirecting
            suggestions.innerHTML = "";
            suggestions.style.display = "none";

            // Redirect to search results page
            window.location.href = `/artists?query=${encodeURIComponent(query)}`;

            // Prevent default form submission behavior
            e.preventDefault();
        }
    });
}

// Check if there are no results on page load
document.addEventListener('DOMContentLoaded', () => {
    const artistsList = document.getElementById('myUL');
    const noResultsMsg = document.getElementById('no-results-message');

    // Only proceed if both elements exist
    if (artistsList && noResultsMsg) {
        // Check if the URL has any search or filter parameters
        const urlParams = new URLSearchParams(window.location.search);
        const hasQuery = urlParams.get('query');
        const hasMember = urlParams.get('member');
        const hasYearMin = urlParams.get('year_min');
        const hasYearMax = urlParams.get('year_max');
        const hasMembersCount = urlParams.get('members_count');

        // Check if any filter or search parameter exists
        const hasAnyFilter = hasQuery || hasMember || hasYearMin || hasYearMax || hasMembersCount;

        // If there's any filter/search but no artists shown, display message
        if (hasAnyFilter && artistsList.children.length === 0) {
            noResultsMsg.style.display = 'block';
            console.log("No results found with current filters/search");
        }
    }
});

// Function to reset all filters and return to full artist list
function clearFilters() {
    window.location.href = '/album';
}

// Make clearFilters available globally for HTML onclick attributes
window.clearFilters = clearFilters;