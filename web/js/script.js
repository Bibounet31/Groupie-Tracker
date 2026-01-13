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

        // Fetch matching results from the server
        const res = await fetch(`/search?query=${encodeURIComponent(query)}&type=${type}`);
        const results = await res.json();  // Parse JSON response (array of names)

        // Clear previous suggestions
        suggestions.innerHTML = "";

        // If no results found, hide the suggestions dropdown
        if (results.length === 0) {
            suggestions.style.display = "none";
            return;
        }

        // Show the suggestions dropdown
        suggestions.style.display = "block";

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
    });



    // SEARCH ON ENTER KEY
    // Handle Enter key press to perform search
    input.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            const query = input.value.trim();  // Get trimmed query
            const type = searchType.value;     // Get search type

            // Don't search if query is empty
            if (!query) return;

            window.location.href = `/artists?query=${encodeURIComponent(query)}`; //search in artist list

            // Prevent default form submission behavior
            e.preventDefault();
        }
    });
}



// Function to reset all filters and return to full artist list
function clearFilters() {
    window.location.href = '/albums';
}

// Make clearFilters available globally for HTML onclick attributes
window.clearFilters = clearFilters;