const input = document.getElementById("myInput");
const suggestions = document.getElementById("suggestions");

input.addEventListener("input", async () => {
    const query = input.value.trim();

    if (!query) {
        suggestions.innerHTML = "";
        suggestions.style.display = "none";
        return;
    }

    const res = await fetch(`/search?query=${encodeURIComponent(query)}`);
    const names = await res.json();

    suggestions.innerHTML = "";
    if (names.length === 0) {
        suggestions.style.display = "none";
        return;
    }

    suggestions.style.display = "block";

    names.forEach(name => {
        const li = document.createElement("li");
        li.textContent = name;
        li.onclick = () => {
            input.value = name;
            suggestions.innerHTML = "";
            suggestions.style.display = "none";
            window.location.href = `/details/${encodeURIComponent(name)}`;
        };
        suggestions.appendChild(li);
    });
});

// Handle Enter key
input.addEventListener("keydown", (e) => {
    if (e.key === "Enter") {
        const query = input.value.trim();
        if (query) {
            // Redirect to artistes page with query
            window.location.href = `/artists?query=${encodeURIComponent(query)}`;
        }
        e.preventDefault();
    }
});
