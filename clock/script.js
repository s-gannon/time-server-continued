document.addEventListener("DOMContentLoaded" , function() {
	setTimeout(incrementTime, 1000);
}, false)

// window.location.reload();

function incrementTime() {
	const localElement = document.getElementById("main-time");
	const solarElement = document.getElementById("solar-time");
	const dateElement = document.getElementById("date");

	localElement.textContent = "A";
	localElement.setAttribute("datetime", "TIME A");
	solarElement.textContent = "B";
	dateElement.textContent = "C"
	dateElement.setAttribute("datetime", "DATE C");
}
