let NewsFrame = document.getElementById('CurrentFrame');
let NextNewsFrame = document.getElementById('nextlink');
let PreviousNewsFrame = document.getElementById('previouslink');

//detect swipe gesture
let pageWidth = window.innerWidth || document.body.clientWidth;
let threshold = Math.max(1, Math.floor(0.01 * (pageWidth)));
let touchstartX = 0;
let touchstartY = 0;
let touchendX = 0;
let touchendY = 0;

const limit = Math.tan(45 * 1.5 / 180 * Math.PI);
const gestureZone = document.getElementById('CurrentFrame');

gestureZone.addEventListener('touchstart', function(event) {
	touchstartX = event.changedTouches[0].screenX;
	touchstartY = event.changedTouches[0].screenY;
}, false);

gestureZone.addEventListener('touchend', function(event) {
	touchendX = event.changedTouches[0].screenX;
	touchendY = event.changedTouches[0].screenY;
	handlegesture(event);
}, false);

function handlegesture(e) {
	let x = touchendX - touchstartX;
	let y = touchendY - touchstartY;

	let xy = Math.abs(x / y);
	let yx = Math.abs(y / x);

	if (Math.abs(x) > threshold || Math.abs(y) > threshold) {
		if (yx <= limit) {
			if(x < 0) {
				console.log("Left");
			} else {
				console.log("Right");
			}
		}
		if (xy <= limit) {
			if (y < 0) {
				console.log("top");
				if (NextNewsFrame !== "") {
					window.location.replace(NextNewsFrame.href);
				}
			} else {
				console.log("bottom");
				if (PreviousNewsFrame !== "") {
					window.location.replace(PreviousNewsFrame.href);
				}
			}
		}
	} else {
		console.log("tap");
	}
}

//check if the corresponding gesture is valid
//if yes, then execute the corresponding function

//if the swipe gesture is up ==> Next Page


//if the swipe gesture is down ==> Previous Page
