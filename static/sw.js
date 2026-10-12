const CACHE_NAME = "arb-calculator-v1"
const APP_SHELL = ["/", "/manifest.webmanifest", "/icon.svg"]

self.addEventListener("install", event => {
	event.waitUntil(
		caches.open(CACHE_NAME)
			.then(cache => cache.addAll(APP_SHELL))
			.then(() => self.skipWaiting())
	)
})

self.addEventListener("activate", event => {
	event.waitUntil(
		caches.keys()
			.then(keys => Promise.all(
				keys.filter(key => key.startsWith("arb-calculator-") && key !== CACHE_NAME)
					.map(key => caches.delete(key))
			))
			.then(() => self.clients.claim())
	)
})

self.addEventListener("fetch", event => {
	const request = event.request
	if (request.method !== "GET" || new URL(request.url).origin !== self.location.origin) {
		return
	}

	if (request.mode === "navigate") {
		event.respondWith(
			fetch(request)
				.then(response => {
					if (!response.ok) {
						return response
					}
					const copy = response.clone()
					return caches.open(CACHE_NAME)
						.then(cache => cache.put("/", copy))
						.then(() => response)
				})
				.catch(() => caches.match("/"))
		)
		return
	}

	event.respondWith(
		caches.match(request).then(cached => cached || fetch(request))
	)
})
