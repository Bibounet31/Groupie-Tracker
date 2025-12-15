const input = document.getElementById("myInput");
const suggestions = document.getElementById("suggestions");
const searchType = document.getElementById("searchType");

if (input && suggestions && searchType) {
    input.addEventListener("input", async () => {
        const query = input.value.trim();
        const type = searchType.value; // "artist" or "member"

        if (!query) {
            suggestions.innerHTML = "";
            suggestions.style.display = "none";
            return;
        }

        const res = await fetch(`/search?query=${encodeURIComponent(query)}&type=${type}`);
        const results = await res.json();

        suggestions.innerHTML = "";
        if (results.length === 0) {
            suggestions.style.display = "none";
            return;
        }

        suggestions.style.display = "block";

        results.forEach(name => {
            const li = document.createElement("li");
            li.textContent = name;
            li.onclick = () => {
                input.value = name;
                suggestions.innerHTML = "";
                suggestions.style.display = "none";

                // Redirect based on type
                if (type === "artist") {
                    window.location.href = `/details/${encodeURIComponent(name)}`;
                } else {
                    // For member search, redirect to filtered artist list
                    window.location.href = `/artists?member=${encodeURIComponent(name)}`;
                }
            };
            suggestions.appendChild(li);
        });
    });

    input.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            const query = input.value.trim();
            const type = searchType.value;
            if (!query) return;

            if (type === "artist") {
                window.location.href = `/artists?query=${encodeURIComponent(query)}`;
            } else {
                window.location.href = `/artists?member=${encodeURIComponent(query)}`;
            }
            e.preventDefault();
        }
    });
}

// Clear filters function - made global so HTML onclick can access it
function clearFilters() {
    window.location.href = '/artistes';
}

// Make it available globally
window.clearFilters = clearFilters;