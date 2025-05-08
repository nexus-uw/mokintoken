async function main() {
  if (location.pathname === '/') {
    const { Welcome } = await import('./welcome')
    new Welcome()
  }

  if (location.pathname.startsWith('/decrypt')) {
    const { Decrypt } = await import('./decrypt')
    new Decrypt()
  }

  if (location.pathname.startsWith('/noteSaved')) {
    const { NoteSaved } = await import('./noteSaved')
    new NoteSaved()
  }
}
main()

const registerServiceWorker = async () => {
  if ("serviceWorker" in navigator) {
    try {
      const registration = await navigator.serviceWorker.register("/service-worker.js", {
        scope: "/",
      })
      if (registration.installing) {
        console.log("Service worker installing")
      } else if (registration.waiting) {
        console.log("Service worker installed")
      } else if (registration.active) {
        console.log("Service worker active")
      }
    } catch (error) {
      console.error(`Registration failed with ${error}`)
    }
  }
}
if (window.matchMedia('(display-mode: standalone)').matches) {
  registerServiceWorker()
}

